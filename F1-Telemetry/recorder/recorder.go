// Package recorder will consume data published by F1 2012, or the playback module
package recorder

import (
	"fmt"
	"net"
	"sync"
)

//PACKETSIZE is UDP packet size for F1 2012
const PACKETSIZE int = 152

//PacketRecorder holds udp playback settings
type PacketRecorder struct {
	connection     *net.UDPConn
	mutex          sync.Mutex
	dataStoreMutex sync.Mutex
}

// NewPacketRecorder instantiates and returns a new PacketRecorder with supplied config
func NewPacketRecorder(IP string, PORT string) *PacketRecorder {
	player := PacketRecorder{}
	udpAddr, err := net.ResolveUDPAddr("udp", IP+":"+PORT)
	if err != nil {
		fmt.Println(err)
	}
	// Make a udp server
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
	}
	player.connection = conn
	return &player
}

// RecordPackets reads bytes from a folder and publishes to udp port
func RecordPackets(player *PacketRecorder) {
	defer player.connection.Close()
	for {
		udpData := make([]byte, PACKETSIZE)
		_, _, err := player.connection.ReadFromUDP(udpData)
		player.mutex.Lock()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(udpData)
		//TODO: convert UDP packet to struct
		//TODO: Add to data store
		player.mutex.Unlock()
	}
}
