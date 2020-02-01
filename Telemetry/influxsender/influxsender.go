package influxsender

// Client documentation https://godoc.org/github.com/influxdata/influxdb1-client/v2
//Go specific doumentationL: http://167.114.231.105/client_libraries/libraries/go/

/*A channel is safe for concurrent access.
select and case is used to wait for different channels
can send data over channels

Read files into a queue in one 'thread'
Publish on udp port on another 'thread', by removing from front of queue
To be able to use queue, use "container/list" which is a doubly linked list
need to call
var list  *List
dataToPub := list.Front() // of type Element
list.Remove(dataToPub)
*/

import (
	"F1-2012-telemetry/f1packet"
	"fmt"
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)

//InfluxSender contains connection data with InfluxDB
type InfluxSender struct {
	connection client.Client
}

// NewInfluxSender makes a new struct
func NewInfluxSender() *InfluxSender {
	sender := InfluxSender{}
	ConnectClient(&sender)
	return &sender
}

//ConnectClient creates a connection to local instance of InfluxDB
func ConnectClient(sender *InfluxSender) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	sender.connection = c
}

//SendData sends data to InfluxDB
func SendData(sender *InfluxSender, packet *f1packet.F1Packet) {

	// Create a new batch of points
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "F1_Telemetry",
		Precision: "ms",
	})

	// Create a new point
	//TODO: convert packet struct to a hashmap, in code, instead of manually
	fields := f1packet.StructToMap(packet)
	// no tags
	pt, err := client.NewPoint("telemPacket", nil, fields, time.Now()) //TODO: use packet.Time or store value of time.Now when we actually receive the new da
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	// Add current point to batch of points
	bp.AddPoint(pt)

	// Write the batch to db
	err = sender.connection.Write(bp)
	if err != nil {
		log.Println("Error: ", err.Error())
	}
}

//CloseConnection closes connection with InfluxDB
func CloseConnection(sender InfluxSender) {
	// Close client resources
	if err := sender.connection.Close(); err != nil {
		fmt.Println(err)
	}
}
