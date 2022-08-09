package traceroute

import (
	"errors"
	"net"
	"syscall"
	"time"
)

const DefaultPort = 33434
const DefaultMaxHops = 64
const DefaultFirstHop = 1
const DefaultTimeoutMs = 500
const DefaultRetries = 3
const DefaultPacketSize = 52

// Return the first non-loopback address as a 4 byte IP address. This address is used for sending packets out.
func socketAddr() (addr [4]byte, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if len(ipnet.IP.To4()) == net.IPv4len {
				copy(addr[:], ipnet.IP.To4())
				return
			}
		}
	}
	err = errors.New("you do not appear to be connected to the Internet")
	return
}

// Given a host name convert it to a 4 byte IP address.
func destAddr(dest string) (destAddr [4]byte, err error) {
	addrs, err := net.LookupHost(dest)
	if err != nil {
		return
	}
	addr := addrs[0]

	ipAddr, err := net.ResolveIPAddr("ip", addr)
	if err != nil {
		return
	}
	copy(destAddr[:], ipAddr.IP.To4())
	return
}

func notify(hop Hop, channels []chan Hop) {
	for _, c := range channels {
		c <- hop
	}
}

func closeNotify(channels []chan Hop) {
	for _, c := range channels {
		close(c)
	}
}

// Traceroute uses the given dest (hostname) and options to execute a traceroute from your machine to the remote host.
// Outbound packets are UDP packets and inbound packets are ICMP.
// Returns a Result which contains an array of hops. Each hop includesthe elapsed time and its IP address.
func traceroute(dest string, options *Options, c ...chan Hop) (result Result, err error) {
	result.Hops = []Hop{}
	destAddr, err := destAddr(dest)
	result.DestinationAddress = destAddr
	socketAddr, err := socketAddr()
	if err != nil {
		return
	}

	timeoutMs := (int64)(options.TimeoutMs())
	tv := syscall.NsecToTimeval(1000 * 1000 * timeoutMs)

	ttl := options.FirstHop()
	retry := 0
	for {
		start := time.Now()

		// Set up the socket to receive inbound packets
		recvSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
		if err != nil {
			return result, err
		}

		// Set up the socket to send packets out.
		sendSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
		if err != nil {
			return result, err
		}
		// This sets the current hop TTL
		syscall.SetsockoptInt(sendSocket, 0x0, syscall.IP_TTL, ttl)
		// This sets the timeout to wait for a response from the remote host
		syscall.SetsockoptTimeval(recvSocket, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &tv)

		defer syscall.Close(recvSocket)
		defer syscall.Close(sendSocket)

		// Bind to the local socket to listen for ICMP packets
		syscall.Bind(recvSocket, &syscall.SockaddrInet4{Port: options.Port(), Addr: socketAddr})

		// Send a single null byte UDP packet
		syscall.Sendto(sendSocket, []byte{0x0}, 0, &syscall.SockaddrInet4{Port: options.Port(), Addr: destAddr})

		var p = make([]byte, options.PacketSize())
		n, from, err := syscall.Recvfrom(recvSocket, p, 0)
		elapsed := time.Since(start)
		if err == nil {
			currAddr := from.(*syscall.SockaddrInet4).Addr

			hop := Hop{Success: true, Address: currAddr, N: n, ElapsedTime: elapsed, TTL: ttl}

			// TODO: this reverse lookup appears to have some standard timeout that is relatively high. Consider switching to something where there is greater control.
			currHost, err := net.LookupAddr(hop.AddressString())
			if err == nil {
				hop.Host = currHost[0]
			}

			notify(hop, c)

			result.Hops = append(result.Hops, hop)

			ttl += 1
			retry = 0

			if ttl > options.MaxHops() || currAddr == destAddr {
				closeNotify(c)
				return result, nil
			}
		} else {
			retry += 1
			if retry > options.Retries() {
				notify(Hop{Success: false, TTL: ttl}, c)
				ttl += 1
				retry = 0
			}

			if ttl > options.MaxHops() {
				closeNotify(c)
				return result, nil
			}
		}
	}
}
