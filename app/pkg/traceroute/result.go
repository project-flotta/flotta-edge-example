package traceroute

// Result TracerouteResult type
type Result struct {
	DestinationAddress [4]byte
	Hops               []Hop
}
