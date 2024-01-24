package find

import (
	"fmt"
	"log"
	"net"
	"os/user"
	"time"
)

// OpenBeFound 开启局域网被发现
func (f *Find) OpenBeFound() {

	log.Println("OpenBeFound")
	// Get the hostname and IP address of the server
	//hostname, err := os.Hostname()
	//if err != nil {
	//	fmt.Println("Error getting hostname:", err)
	//	return
	//}
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("无法获取当前用户信息:", err)
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

	// Create a UDP connection to broadcast messages
	broadcastAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("255.255.255.255:%d", discoveryPort))
	if err != nil {
		fmt.Println("Error resolving broadcast address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, broadcastAddr)
	if err != nil {
		fmt.Println("Error dialing UDP connection:", err)
		return
	}
	defer conn.Close()

	// Continuously send broadcast messages with name and IP address
	for {
		message := fmt.Sprintf("%s|%s", currentUser.Username, serverIP)
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending broadcast message:", err)
		}

		time.Sleep(5 * time.Second) // Adjust the interval as needed
	}
}

// CloseBeFound 关闭局域网被发现
func (f *Find) CloseBeFound() {
	//f.ServerConn.Close()
}
