package ardugorpc

import "fmt"

const _protocol = "simpleRPC"

var _version = [3]int{3, 0, 0}

func _assert_protocol(protocol string) {
	if protocol != _protocol {
		panic("invalid protocol header")
	}
}

func _assert_version(version []int) {
	if version[0] != _version[0] || version[1] > _version[1] {
		panic(fmt.Sprintf("version mismatch (device: %v, client: %v)", version, _version))
	}
}
