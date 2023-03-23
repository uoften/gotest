package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//for i:=0;i<10;i++ {
	//	unixRandFlag := rand.Intn(99999)
	//
	//	fmt.Printf("%05d\n", unixRandFlag)
	//	fmt.Printf("%0*d\n", 5, unixRandFlag)
	//}
	No,err := createStoreNo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(No)
}

func createStoreNo() (string, error) {
	format := "20060102"
	var (
		merchantTag string
		storeNo     string
	)
	merchantTag = GenerateMerchantUniqueTag()
	for i := 0; i < 3; i++ {
		dateFlag := time.Now().Format(format)[2:]
		fmt.Println(time.Now().UnixMilli())
		unixFlag := strconv.FormatInt(time.Now().Unix(), 10)[5:]
		fmt.Println(unixFlag)
		storeNo = merchantTag + dateFlag + unixFlag
		//_, err := db_handler.GetStore(bson.M{"store_no": storeNo})
		//if err != nil {
		//	if err == mgo.ErrNotFound {
		//		return storeNo, nil
		//	} else {
		//		utils.Logger.Info("createStoreNo duplicate check error=" + err.Error())
		//		return "", err
		//	}
		//} else {
		//	time.Sleep(time.Second)
		//}
		return storeNo, nil
	}
	return "", errors.New("createStoreNo duplicate error")
}

const letters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateMerchantUniqueTag() string {
	rand.Seed(time.Now().Unix())
	index := []int32{rand.Int31() % 26, rand.Int31() % 26, rand.Int31() % 26}

	tag := ""
	for _, v := range index {
		tag += string(letters[v])
	}

	return tag
}
