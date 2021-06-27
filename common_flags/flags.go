package common_flags

import "flag"

var (
	Client_dial = flag.String(
		"client_dial", "ws://192.168.1.2:9551", "could be websocket or IPC",
	)
)
