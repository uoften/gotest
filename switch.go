package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%#v\n",getProtocolByCode("asd"))
	reIndex := 3
	switch reIndex {
	case 1:
		fmt.Println("111")
	case 2:
		fmt.Println("222")
	case 3,4:
		fmt.Println("333")
	default:
		fmt.Println("1234")
	}
}

func getProtocolByCode(code string) (protocol string) {
	ProtocolMap := map[string]string{
		"L": "local",
		"C": "connected",
		"S": "static",
		"R": "RIP",
		"M": "mobile",
		"B": "BGP",
		"E": "EIGRP",
		"EX": "EIGRP external",
		"O": "OSPF",
		"IA": "OSPF inter area",
		"N1": "OSPF NSSA external type 1",
		"N2": "OSPF NSSA external type 2",
		"E1": "OSPF external type 1",
		"E2": "OSPF external type 2",
		"i": "IS-IS",
		"su": "IS-IS summary",
		"L1": "IS-IS level-1",
		"L2": "IS-IS level-2",
		"ia": "IS-IS inter area",
		"*": "candidate default",
		"U": "per-user static route",
		"o": "ODR",
		"P": "periodic downloaded static route",
		"H": "NHRP",
		"I": "LISP",
		"a": "application route",
		"+": "replicated route",
		"%": "next hop override",
	}
	protocol = ProtocolMap[code]
	return
}