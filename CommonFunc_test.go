package utils

import (
	"fmt"
	"testing"
)

func Test_GetPublicIpFromHostsByAlias(t *testing.T){
	ip,err :=GetPublicIpFromHostsByAlias("/etc/hosts","local_service_ip")
	fmt.Println(ip)
}
