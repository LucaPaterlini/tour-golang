package old

import (
	"fmt"
	"strings"
	"generics"
)

type IPAddr [4]byte


func (ip IPAddr) String() string {
	r := generics.GenericMap(ip, func(x byte) string {
		return fmt.Sprintf("%v",x)
	})
	return strings.Join(r.([]string),".")

	// solution easy
	//return fmt.Sprintf("%v.%v.%v.%v",ip[0],ip[1],ip[2],ip[3])
}

func main(){
	hosts:= map[string]IPAddr{
		"loopback": {127,0,0,1},
		"googleDNS": {8,8,8,8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n",name,ip)
	}
}
