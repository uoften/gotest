package main

import (
	"fmt"
	"strconv"
)
func IndexArray(haystack []string, needle string) int {
	var index = -1
	for k, e := range haystack {
		if e == needle {
			index = k
		}
	}
	return index
}
func main()  {
	//for i:=1;i<=334;i++ {
	//	ii := strconv.Itoa(3*i)
	//	fmt.Println("PARTITION s" + strconv.Itoa(i) +" VALUES LESS THAN("+ii+"0000),")
	//}
	//fmt.Println("PARTITION s501 VALUES LESS THAN(20000000),")
//
//fmt.Println("")
//fmt.Println("")
//fmt.Println("")

	for i:=1;i<=100;i++ {
		ii := strconv.Itoa(20*i)
		fmt.Println("PARTITION t" + strconv.Itoa(i) +" VALUES LESS THAN("+ii+"0000),")
	}


	//fmt.Printf("%v\n\n\n","")
	//var arr0 = [3]string{"a", "s", "d"}
	//fmt.Printf("%#v\n",arr0)
	//s1 := arr0[0:2]
	//fmt.Printf("%#v\n",reflect.TypeOf(arr0).String()) //[3]string
	//fmt.Printf("%#v\n",reflect.TypeOf(s1).String()) //[]string
	////
	//sshErrHost := []string{"192.200.0.1","192.200.0.2","192.200.0.85"}
	//fmt.Println(sshErrHost)
	////删除第i个元素
	//index := IndexArray(sshErrHost,"192.200.0.2")
	//if index >= -1{
	//	sshErrHost = append(sshErrHost[:index], sshErrHost[index+1:]...)
	//}
	//fmt.Println(sshErrHost)
}
