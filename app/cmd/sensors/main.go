package sensors

import (
	"fmt"
	"time"
)

func main() {
	//generateDummyLogs()
}

// GenerateDummyLogs dummy logs generator, generate log every one second
func GenerateDummyLogs() {
	for {
		fmt.Println("New message at: ", time.Now().Format(time.UnixDate))
		time.Sleep(time.Second)
	}
}
