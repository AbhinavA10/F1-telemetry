## Install InfluxDB
Documentation about InfluxDB https://docs.influxdata.com/influxdb/v1.7/

Ubuntu:
`wget https://dl.influxdata.com/influxdb/releases/influxdb_1.7.9_amd64.deb`

`sudo dpkg -i influxdb_1.7.9_amd64.deb`

then

```bash
sudo systemctl unmask influxdb.service
sudo systemctl start influxdb
```

```bash
sudo systemctl disable influxdb.service
sudo systemctl start influxdb
```

## GO

At the time of writing this, Only 1.x client was available for Go

Importing client for Golang: https://github.com/influxdata/influxdb1-client

`go get github.com/influxdata/influxdb1-client/v2`

```go
import(
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)
```



