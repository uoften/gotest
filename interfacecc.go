package main
func GetValue() int {
	return 1
}

func main() {
	i := GetValue()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}
//interfacecc.go:8:2: cannot type switch on non-interface value i (type int)
//i (type int)是非接口类型不能使用类型switch选择
//只有接口才可以做类型switch选择
