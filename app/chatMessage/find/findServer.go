package find

import (
	"fmt"
	"net"
	"os"
	"time"
)

// OpenBeFound 开启局域网被发现
func (f *Find) OpenBeFound() {
	// Get the hostname and IP address of the server
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error getting IP address:", err)
		return
	}

	var serverIP string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				serverIP = ipnet.IP.String()
				break
			}
		}
	}

	if serverIP == "" {
		fmt.Println("Error finding a suitable IP address.")
		return
	}

	// Continuously send broadcast messages with name and IP address
	for {
		message := fmt.Sprintf("%s|%s", hostname, serverIP)
		_, err := f.ServerConn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending broadcast message:", err)
		}

		time.Sleep(5 * time.Second) // Adjust the interval as needed
	}
}

// CloseBeFound 关闭局域网被发现
func (f *Find) CloseBeFound() {
	f.ServerConn.Close()
}
