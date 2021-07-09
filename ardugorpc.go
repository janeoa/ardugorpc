package ardugorpc

const _protocol = "simpleRPC"

var _version = [3]int{3, 0, 0}

func _assert_protocol(protocol string) {
	if protocol != _protocol {
		panic("invalid protocol header")
	}
}
