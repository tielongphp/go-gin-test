package tool

import "net"

func GetHostAddress() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, value := range addresses {
		if ipNet, ok := value.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
