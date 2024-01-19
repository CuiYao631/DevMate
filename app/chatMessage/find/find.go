package find

import (
	"fmt"
	"net"
)

const discoveryPort = 55555

type Find struct {
	//用户名
	Name string
	//用户头像
	Avatar string
	//用户id
	UserId string

	clientConn *net.UDPConn
	ServerConn *net.UDPConn
}

func NewFind(name, avatar, userId string) *Find {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", discoveryPort))
	if err != nil {
		fmt.Println("Error resolving address:", err)
	}
	clientConn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
	}
	defer clientConn.Close()
	// Create a UDP connection to broadcast messages
	broadcastAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("255.255.255.255:%d", discoveryPort))
	if err != nil {
		fmt.Println("Error resolving broadcast address:", err)
	}

	ServerConn, err := net.DialUDP("udp", nil, broadcastAddr)
	if err != nil {
		fmt.Println("Error dialing UDP connection:", err)
	}

	return &Find{
		Name:       name,
		Avatar:     avatar,
		UserId:     userId,
		clientConn: clientConn,
		ServerConn: ServerConn,
	}
}
