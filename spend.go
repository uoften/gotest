package main

/*
#include <stdio.h>

void myhello(int i) {
  printf("Hello C: %d\n", i+1);
}
*/
import "C"

import "fmt"

func main() {
	C.myhello(12)
	fmt.Println("Hello Go")
}