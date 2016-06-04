package management

/*
URL: https://github.com/mccoyst/myip/blob/master/myip.go
URL: http://changsijay.com/2013/07/28/golang-get-ip-address/
*/

import (
	"net"
)

func GetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}