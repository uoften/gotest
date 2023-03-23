package main

import "fmt"

func main()  {
	exit := 0
	i:=100
	for {
		if i == 110 {
			return
		}
		for {
			if exit == 20 {
				break
			}
			if exit == 10 {
				goto END
			}
			fmt.Println(exit)
			if true{
				exit++
			}
		}
		fmt.Println(i)
		i++
	}
	fmt.Println("ok")
	END:
		fmt.Println("ok")
}
