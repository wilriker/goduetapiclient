package machine

const (
	DefaultName     = "My Duet"
	DefaultHostName = "duet"
	DefaultPassword = "reprap"
)

type InterfaceType string

const (
	WiFi InterfaceType = "wifi"
	LAN                = "lan"
)

type NetworkProtocol string

const (
	HTTP   NetworkProtocol = "http"
	FTP                    = "ftp"
	Telnet                 = "telnet"
)

type Network struct {
	Name       string
	Hostname   string
	Password   string
	Interfaces []NetworkInterface
}

type NetworkInterface struct {
	Type            InterfaceType
	FirmwareVersion string
	Speed           uint64
	Signal          int64
	MacAddress      string
	ConfiguredIP    string
	ActualIP        string
	Subnet          string
	Gateway         string
	NumReconnnects  uint64
	ActiveProtocols []NetworkProtocol
}
