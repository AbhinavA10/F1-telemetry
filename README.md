# F1-telemetry
Telemetry client for F1 2012, in GoLang

UDP Packets were recorded using [this C# function](https://github.com/AbhinavA10/F1-telemetry/blob/master/ByteWriter/F1_Capture/Main_Program.cs#L67) .

Playback of packets, and parsing of packets was done in Golang


| Folder             | Purpose                            |
| ------------------ | ---------------------------------- |
| F1-Packet-Playback | Playback UDP data from `.bin` file |
| F1-Telemetry       | Consumes UDP packets               |

I decided to use InfluxDB as my time series Database, as it is quite popular for realtime analytics.
I installed this locally on a ubuntu machine as per [this](./F1-Telemetry/influxsender/InfluxConnectionInstructions.md).

Structure:
udp packet player in GO publishes on `20777`
recorder reads this datagram from `20777` and converts to a F1Packet struct
...
influxdb object sends this packet, converted, to influxdb
chronograf or graphana queries influxdb and visualizes it.


Using Grafana and InfluxDB
https://grafana.com/docs/grafana/latest/features/datasources/influxdb/


TODO:
- [ ] Write data from go to InfluxDB using go client
- [ ] Install Grafana or Chronograf
- [ ] Interface the above with InfluxDB