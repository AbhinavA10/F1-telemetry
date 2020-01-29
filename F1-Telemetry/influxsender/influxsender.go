package influxsender

// To add a new point to influxdb, I will need to convert the F1Packet struct to a hashmap
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
	"fmt"
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
	_ "log"
	"time"
)

//ConnectClient creates a connection to local instance of InfluxDB
func ConnectClient() {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	//defer c.Close()

	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "F1_Telemetry",
		Precision: "ms",
	})

	// Create a point and add to batch
	tags := map[string]string{"type": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("precise_measurement", tags, fields, time.Now())
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	fmt.Println(pt)
	bp.AddPoint(pt)

	// Write the batch
	err = c.Write(bp)
	if err != nil {
		fmt.Println("failed")
	} else {
		fmt.Println("wrote")
	}

	// Close client resources
	if err := c.Close(); err != nil {
		fmt.Println(err)
	}

}
