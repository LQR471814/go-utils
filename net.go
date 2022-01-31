package utils

import "net"

//ListenerPort returns the port a listener is listening on
func ListenerPort(l net.Listener) int {
	return l.Addr().(*net.TCPAddr).Port
}

//CheckLocal checks if the given IP is a local address
func CheckLocal(target net.IP) (bool, error) {
	intfs, err := net.Interfaces()
	if err != nil {
		return false, err
	}

	for _, i := range intfs {
		addrs, err := i.Addrs()
		if err != nil {
			return false, err
		}

		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				return false, err
			}

			if target.Equal(ip) {
				return true, nil
			}
		}
	}

	return false, nil
}

//AddrInInterface checks if the given UDPAddr belongs to a certain interface
func UDPAddrInInterface(intf *net.Interface, addr *net.UDPAddr) (bool, error) {
	addrs, err := intf.Addrs()
	if err != nil {
		return false, err
	}

	for _, a := range addrs {
		if a.String() == addr.String() {
			return true, nil
		}
	}

	return false, nil
}

//StringToUDPAddr converts a string to a UDPAddr
func StringToUDPAddr(s string) (*net.UDPAddr, error) {
	ip := net.ParseIP(s)
	addr, err := net.ResolveUDPAddr("udp", ip.String())
	if err != nil {
		return nil, err
	}

	return addr, nil
}
