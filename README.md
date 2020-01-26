# F1-telemetry
Telemetry client for F1 2012, in GoLang

UDP Packets were recorded using [this C# function](https://github.com/AbhinavA10/F1-telemetry/blob/master/ByteWriter/F1_Capture/Main_Program.cs#L67) .

Playback of packets, and parsing of packets was done in Golang


| Folder             | Purpose                            |
| ------------------ | ---------------------------------- |
| F1-Packet-Playback | Playback UDP data from `.bin` file |
| F1-Telemetry       | Consumes UDP packets               |