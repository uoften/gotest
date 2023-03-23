package main

func main() {
	for ii := 0; ii < 5; ii++ {
		defer func() {
			println(ii)
		}()
	}

	//for i := 0; i < 5; i++ {
	//	defer func(i int) {
	//		println(i)
	//	}(i)
	//}
}
