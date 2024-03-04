package find

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"net"
	"strings"
)

type DiscoveredUser struct {
	Hostname string
	IP       string
}
type users struct {
	Name   string
	IP     string
	Avatar string
}

// OpenFind 开启局域网发现
func (f *Find) OpenFind() {
	// Create a UDP address to listen on
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", discoveryPort))
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create a UDP listener
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()
	// Get the IP addresses associated with the local machine
	localIPs, err := getLocalIPs()
	if err != nil {
		fmt.Println("Error getting local IP addresses:", err)
		return
	}

	fmt.Println("User discovery client is listening on", conn.LocalAddr())

	buffer := make([]byte, 1024)

	userList := make(map[string]users)
	for {
		// Read data from the connection
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		if n == 0 {
			runtime.EventsEmit(f.ctx, "dataUpdated", nil)
		}

		// Extract hostname and IP address from received data
		data := string(buffer[:n])
		parts := strings.Split(data, "|")
		if len(parts) != 2 {
			fmt.Println("Invalid data format:", data)
			continue
		}

		hostname := parts[0]
		ip := parts[1]

		if !contains(localIPs, ip) {

			userList[ip] = users{
				Name:   hostname,
				IP:     ip,
				Avatar: "",
			}
		}
		runtime.EventsEmit(f.ctx, "dataUpdated", userList)
		// 删除不在当前UDP消息中的用户
		for k := range userList {
			if !strings.Contains(data, k) {
				delete(userList, k)
			}
		}

	}
}

// CloseFind 关闭局域网发现
func (f *Find) CloseFind() {
	//f.clientConn.Close()
}

// getLocalIPs returns a list of local machine IP addresses
func getLocalIPs() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var localIPs []string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIPs = append(localIPs, ipnet.IP.String())
			}
		}
	}

	return localIPs, nil
}

// contains checks if a string is present in a slice of strings
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
