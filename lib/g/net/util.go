package net

import (
	"fmt"
)

func IPv4AddrString(ipv4 string, port uint16) string {
	return fmt.Sprintf("%s:%d", ipv4, port)
}
