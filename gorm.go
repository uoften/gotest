package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"strconv"
	"time"
)

type Test struct {
	Id int
	Num int
}

func createTables(db *gorm.DB) {
	db.CreateTable(&Test{})
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		fmt.Println("open db sucess")
	} else {
		fmt.Println("open db error ", err)
		return
	}
	//syncMaxIdle := db.maxOpen > 0 && db.maxIdleConnsLocked() > db.maxOpen成立的时候
	//会把maxidleconn设置成和maxopen一样.
	//所以空闲MaxIdleConns应该要小于最大MaxOpenConns
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)

	if !db.HasTable("tests") {
		createTables(db)
	}

	test := Test{Num:Rand(1,99)}
	db.Create(&test)
	fmt.Println("test.id is ", test.Id)

	var tests []Test
	db.Find(&tests)
	//fmt.Println(tests)

	for index, line := range tests {
		fmt.Println("index", index, " line ", line)
	}

	defer db.Close()
}


func Rand(min, max int) int {
	if min > max {
		panic("min: min cannot be greater than max")
	}
	// PHP: getrandmax()
	if int31 := 1<<31 - 1; max > int31 {
		panic("max: max can not be greater than " + strconv.Itoa(int31))
	}
	if min == max {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max+1-min) + min
}