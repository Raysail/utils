package utils

import (
	"errors"
	"io/ioutil"
	"net"
	"strings"
)

func GetPublicIpFromHostsByAlias(hostsPath, alias string) (ip string, err error) {
	content, err := ioutil.ReadFile(hostsPath)
	if err != nil {
		return "", err
	}
	hosts := strings.Split(string(content), "\n")
	for _, hostLine := range hosts {
		ipHostAliasMap := strings.Split(hostLine, "\t")
		if len(ipHostAliasMap) < 2 {
			continue
		}
		for _, host := range ipHostAliasMap[1:] {
			if strings.EqualFold(host, alias) {
				return ipHostAliasMap[0], nil
			}
		}
	}
	return "", errors.New("cann't found ip addr with given hostname: " + alias)
}
func GetAFreePort(ip string) (port int, err error) {
	addr, err := net.ResolveTCPAddr("tcp", ip+":0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
