package main

import (
	"encoding/json"
	"fmt"
)

func AddUpper() func (int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}
//func main(){
//	f:=AddUpper()
//	fmt.Println(f(1))
//	fmt.Println(f(2))
//	fmt.Println(f(3))
//}
type RemarkUnit struct {
	Type string `json:"type,omitempty"`
	Remark string `json:"remark,omitempty"`
}
type StoreRemarkIn struct {
	StoreUrl     *RemarkUnit `json:"storeUrl,omitempty"`
	StoreName    *RemarkUnit `json:"storeName,omitempty"`
	StoreType    *RemarkUnit `json:"storeType,omitempty"`
	StoreTel     *RemarkUnit `json:"storeTel,omitempty"`
	StoreAddress *RemarkUnit `json:"storeAddress,omitempty"`
	Longitude    *RemarkUnit `json:"longitude,omitempty"`
	StorePhoto   *RemarkUnit `json:"storePhoto,omitempty"`
	StoreBsHours *RemarkUnit `json:"storeBsHours,omitempty"`
	PerConsume   *RemarkUnit `json:"perConsume,omitempty"` //人均消费
}
type VoucherRemarkIn struct {
	OriginPrice          *RemarkUnit `json:"originPrice,omitempty"`    //面值
	ActivePrice          *RemarkUnit `json:"activePrice,omitempty"`    //销售价格
	Inventory            *RemarkUnit `json:"inventory,omitempty"`      //库存
	UseDate              *RemarkUnit `json:"date,omitempty"`           //使用开始日期、使用结束日期
	UseDays              *RemarkUnit `json:"useDays,omitempty"`        //使用时段
	UseHours             *RemarkUnit `json:"useHours,omitempty"`       //使用时间
	SingleClientBuyLimit *RemarkUnit `json:"singleClientBuyLimit,omitempty"` //单人购买数量
	IsOrder              *RemarkUnit `json:"isOrder,omitempty"`        //是否预约
	Notes                *RemarkUnit `json:"notes,omitempty"`          //备注
}
type VoucherRemarkOut struct {
	OriginPrice          *RemarkUnit `json:"origin_price"`            //面值
	ActivePrice          *RemarkUnit `json:"active_price"`            //销售价格
	Inventory            *RemarkUnit `json:"inventory"`               //库存
	UseDate              *RemarkUnit `json:"use_date"`                //使用开始日期、使用结束日期
	UseDays              *RemarkUnit `json:"use_days"`                //使用时段
	UseHours             *RemarkUnit `json:"use_hours"`               //使用时间
	SingleClientBuyLimit *RemarkUnit `json:"single_client_buy_limit"` //单人购买数量
	IsOrder              *RemarkUnit `json:"is_order"`                //是否预约
	Notes                *RemarkUnit `json:"notes"`                   //备注
}
func main() {
	s := "{'active_price':{'remark':'1'},'inventory':{'remark':'2'}}" // 分配存储"A1"的内存空间，s结构体里的str指针指向这快内存
	remarkIn := VoucherRemarkIn{}
	err := json.Unmarshal([]byte(s), &remarkIn)
	if err != nil {
		//fmt.Println(err)
	}
	fmt.Println(&VoucherRemarkOut{
		OriginPrice:          remarkIn.OriginPrice,
		ActivePrice:          remarkIn.ActivePrice,
		Inventory:            remarkIn.Inventory,
		UseDate:              remarkIn.UseDate,
		UseDays:              remarkIn.UseDays,
		UseHours:             remarkIn.UseHours,
		SingleClientBuyLimit: remarkIn.SingleClientBuyLimit,
		IsOrder:              remarkIn.IsOrder,
		Notes:                remarkIn.Notes,
	})

	//s := []byte{'a','b'} // 分配存储1数组的内存空间，s结构体的array指针指向这个数组。
	//fmt.Println(s)
	//fmt.Println(&s)
	//s = []byte{['c'],['d']}  // 将array的内容改为2
	//fmt.Println(s)
	//fmt.Println(&s)

	//m := make(map[string]interface{})
	//m["int"] = 123
	//m["string"] = "hello"
	//m["bool"] = true
	//
	//a :=1
	//fmt.Println(reflect.TypeOf(a))
	//
	//for _, v := range m {
	//	switch v.(type) {
	//	case string:
	//		fmt.Println(v, "is string")
	//	case int:
	//		fmt.Println(v, "is int")
	//	default:
	//		fmt.Println(v, "is other")
	//	}
	//}
	//fmt.Println(m)

	//var str1 string = "123我爱学习"
	//fmt.Println(str1)
	//a := [4]float64{67.7, 89.8, 21, 78}
	//b := [...]int{2, 3, 5}
	//// 遍历数组⽅方式 1
	//for i := 0; i < len(a); i++ {
	//	fmt.Print(a[i], "\t")
	//}
	//fmt.Println()
	//// 遍历数组⽅方式 2
	//for _, value := range b {
	//	fmt.Print(value, "\t")
	//}

	//var a = "1"
	//b := "2"
	//c := "3"
	////number := [3]int{5, 6, 7}
	//ptrs := [3]*string{&a,&b,&c} //指针数组
	////将number数组的值的地址赋给ptrs
	////for i, x := range &number {
	////	ptrs[i] = &x
	////}
	//for i, x := range ptrs {
	//	fmt.Printf("%d %v %d\n", i, *x, &x)
	//}
}