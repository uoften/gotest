package library

import (
	"fmt"
	"github.com/sony/sonyflake"
)

func GenSonyflake() string{
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		fmt.Println("flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	return string(id)
	//fmt.Printf("github.com/sony/sonyflake:   %x\n", id)
}
