package dynamicip

import "net"

// IPer defines an interface for obtaining the public IP address.
type IPer interface {
	IP() (net.IP, error)
}
