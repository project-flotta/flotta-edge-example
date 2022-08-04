package traceroute

import (
	"fmt"
	"testing"
)

func TestTraceroute(t *testing.T) {
	fmt.Println("Testing synchronous traceroute")
	out, err := traceroute("google.com", new(Options))
	if err == nil {
		if len(out.Hops) == 0 {
			t.Errorf("TestTraceroute failed. Expected at least one hop")
		}
	} else {
		t.Errorf("TestTraceroute failed due to an error: %v", err)
	}

	for _, hop := range out.Hops {
		printHop(hop)
	}
	fmt.Println()
}

func TestTracerouteChannel(t *testing.T) {
	fmt.Println("Testing asynchronous traceroute")
	c := make(chan Hop, 0)
	go func() {
		for {
			hop, ok := <-c
			if !ok {
				fmt.Println()
				return
			}
			printHop(hop)
		}
	}()

	out, err := traceroute("google.com", new(Options), c)
	if err == nil {
		if len(out.Hops) == 0 {
			t.Errorf("TestTracerouteChannel failed. Expected at least one hop")
		}
	} else {
		t.Errorf("TestTraceroute failed due to an error: %v", err)
	}
}
