package ardugorpc

import (
	"fmt"
	"io"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

const _protocol = "simpleRPC"

var _version = [3]byte{3, 0, 0}

func _assert_protocol(protocol string) {
	if protocol != _protocol {
		panic("invalid protocol header")
	}
}

func _assert_version(version []byte) {
	if version[0] != _version[0] || version[1] > _version[1] {
		panic(fmt.Sprintf("version mismatch (device: %v, client: %v)", version, _version))
	}
}

type Interface struct {
	self        io.ReadWriteCloser
	device      string
	baudrate    uint
	wait        int
	autoconnect bool
	// load TextIO=None
}

func NewInterface(device string, baudrate uint) *Interface {
	options := serial.OpenOptions{
		PortName:        device,
		BaudRate:        baudrate,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	defer port.Close()

	return &Interface{self: port, device: device, baudrate: baudrate, wait: 0, autoconnect: true}
}

func (in Interface) MethodDiscovery() {
	var tosend []byte
	tosend = append(tosend, 0x00)
	tosend = append(tosend, _version[:]...)
	tosend = append(tosend, []byte(">")...)
	tosend = append(tosend, []byte("H")...)

	in.send(tosend)
}

func (in Interface) send(data []byte) {
	_, err := in.self.Write(data)
	if err != nil {
		panic(err)
	}
}
