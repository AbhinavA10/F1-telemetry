// Package recorder will consume data published by F1 2012, or the playback module
package recorder

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

//PACKETSIZE is UDP packet size for F1 2012
const PACKETSIZE int = 152
const baseFilePath string = "/media/abhi_ubuntu18/600C7EC53DCD4841/f1/NewPackets"

//PacketRecorder holds udp playback settings
type PacketRecorder struct {
	connection *net.UDPConn
	mutex      sync.Mutex
}

// NewPacketRecorder instantiates and returns a new PacketRecorder with supplied config
func NewPacketRecorder(IP string, PORT string) *PacketRecorder {
	recorder := PacketRecorder{}
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
	recorder.connection = conn
	return &recorder
}

// RecordPackets reads bytes from UDP connection, and saves to a file
func RecordPackets(recorder *PacketRecorder) {
	var n int
	for {
		udpData := make([]byte, PACKETSIZE)
		recorder.mutex.Lock()
		_, _, err := recorder.connection.ReadFromUDP(udpData)
		if err != nil {
			fmt.Println(err)
		}
		WriteData(&udpData, n)
		n++
		recorder.mutex.Unlock()
	}
}

//WriteData writes given bytes to a file.
func WriteData(data *[]byte, n int) {
	filepath := baseFilePath + "/Packet" + strconv.Itoa(n) + ".bin"
	f, err := os.Create(filepath)
	check(err)
	defer f.Close()
	n2, err := f.Write(*data)
	check(err)
	// fmt.Printf("wrote %d bytes\n", n2)
	f.Sync()
	_ = n2
}

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}
