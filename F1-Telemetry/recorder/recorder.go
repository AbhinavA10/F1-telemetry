// Package recorder will consume data published by F1 2012, or the playback module
package recorder

import (
	"F1-2012-telemetry/f1packet"
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
	// maybe need to use setReuseAddress(true)
	if err != nil {
		fmt.Println(err)
	}
	player.connection = conn
	return &player
}

// RecordPackets reads bytes from UDP connection, and converts to struct
func RecordPackets(player *PacketRecorder) {
	defer player.connection.Close()
	for {
		udpData := make([]byte, PACKETSIZE)
		_, _, err := player.connection.ReadFromUDP(udpData)
		player.mutex.Lock()
		if err != nil {
			fmt.Println(err)
		}
		telemPacket := f1packet.DatagramToStruct(udpData)
		fmt.Println(telemPacket)
		//TODO: Somewhere, send to influxDB with a time of telemPacket.time
		player.mutex.Unlock()
	}
}
