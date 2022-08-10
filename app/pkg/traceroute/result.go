package traceroute

// Result is the traceroute result
type Result struct {
	DestinationAddress [4]byte
	Hops               []Hop
}
