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

## Configuring InfluxDB
- Data and WAL directories were left as is
- `influx`, then `CREATE DATABASE F1_Telemetry`
- edited `[[udp]]` portion of `influxdb.conf` to look like this as per https://docs.influxdata.com/influxdb/v1.7/supported_protocols/udp/
```
# influxd.conf
[[udp]]
  enabled = true
  bind-address = ":8089"
  database = "F1_Telem"
  precision="ms"
  batch-size = 5000
  batch-timeout = "1s"
  batch-pending = 10
  read-buffer = 0
```

TSM (Time Structured Merge tree) - The purpose-built data storage format for InfluxDB. 
WAL (Write Ahead Log) - The temporary cache for recently written points. 

Data Format: https://docs.influxdata.com/influxdb/v1.7/introduction/getting-started/
Data Concepts: (Good visuals) https://docs.influxdata.com/influxdb/v1.7/concepts/key_concepts/

- data is organized by 'time series', which can have many `points`
- a `point` consists of a `measurement`, 0-inf `tags`, 1-inf `field` and a optional `time`
- measurment: the actual label for what we are measuring. Ex. temp
- tag: usually metadata. these are indexed.  ex. computer=15inch,cpu_type=ARM
  - Value for this key-val par is always a string
- field: the measured value iteself: ex. internal=21,external=10
  - value for this key-val pair can be whatever we want
  - unindexed

Essentially, measreument is an sql table, where primary index is `time`. `tag` and `fields` are columns in the table. But there's no defined schema in influxdb
```
<measurement>[,<tag-key>=<tag-value>...] <field-key>=<field-value>[,<field2-key>=<field2-value>...] [unix-nano-timestamp]
```

Example:

```
> INSERT temperature,machine=unit42,type=assembly external=25,internal=37   ## the space signifies the next param type
> SELECT * FROM "temperature"
name: temperature
time                external internal machine type
----                -------- -------- ------- ----
1580094311037678743 25       37       unit42  assembly # no time stamp was supplied, so influx assigned the local time.
```


- `retention policy` is how long InfluxDB will keep the data, and how many times to replicate the data
  - default is `autogen` which has inf duration

- A point represents a single data record that has four components: a measurement, tag set, field set, and a timestamp.