// Package receiver will consume data published by F1, or the playback module
package receiver

import (
	"F1-telemetry/f1packet"
	"fmt"
	"net"
	"sync"
)

//PACKETSIZE is UDP packet size for F1
const PACKETSIZE int = 152

//PacketReceiver holds udp playback settings
type PacketReceiver struct {
	connection     *net.UDPConn
	mutex          sync.Mutex
	dataStoreMutex sync.Mutex
}

// NewPacketReceiver instantiates and returns a new PacketReceiver with supplied config
func NewPacketReceiver(IP string, PORT string) *PacketReceiver {
	receiver := PacketReceiver{}
	udpAddr, err := net.ResolveUDPAddr("udp", IP+":"+PORT)
	if err != nil {
		fmt.Println(err)
	}
	// Make a udp server
	conn, err := net.ListenUDP("udp", udpAddr)
	// maybe need to use setReuseAddress(true)
	if err != nil {
		fmt.Println(err)
	}
	receiver.connection = conn
	return &receiver
}

// ReceivePacket reads bytes from UDP connection, and converts to struct
func ReceivePacket(receiver *PacketReceiver) *f1packet.F1Packet {
	udpData := make([]byte, PACKETSIZE)
	receiver.mutex.Lock()
	_, _, err := receiver.connection.ReadFromUDP(udpData)
	if err != nil {
		fmt.Println(err)
	}
	telemPacket := f1packet.DatagramToStruct(udpData)
	receiver.mutex.Unlock()
	return telemPacket
}
