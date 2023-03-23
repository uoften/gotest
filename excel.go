package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

type hosts struct {
	Ip string
	Brand string
	Type string
	TypeNum string
	Version string
	VersionRun string
	Snmp string
	Username string
	Password string
}

type cmdList struct {
	Brand string
	Type string
	TypeNum string
	VersionRun string
	CmdList string
}

type errList struct {
	Brand string
	WrongStr string
	ErrorStr string
}

func main() {
	f, err := excelize.OpenFile("hosts.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	hostsMap := make(map[string]hosts)
	for key, row := range rows {
		if key==0{
			continue
		}
		var h hosts
		for k, colCell := range row {
			if k==0{
				h.Ip = colCell
			}else if k==1{
				h.Brand=colCell
			}else if k==2{
				h.Type=colCell
			}else if k==3{
				h.TypeNum=colCell
			}else if k==4{
				h.Version=colCell
			}else if k==5{
				h.VersionRun=colCell
			}else if k==6{
				h.Snmp=colCell
			}else if k==7{
				h.Username=colCell
			}else if k==8{
				h.Password=colCell
			}
		}
		hostsMap[h.Ip] = h
	}

	cmdListRows, err := f.GetRows("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	cmdListMap := make(map[string]cmdList)
	for key, hostCmd := range cmdListRows {
		if key==0{
			continue
		}
		var c cmdList
		for k, item := range hostCmd {
			if k==0{
				c.Brand = item
			}else if k==1{
				c.Type=item
			}else if k==2{
				c.TypeNum=item
			}else if k==3{
				c.VersionRun=item
			}else if k==4{
				c.CmdList = item
			}
		}
		cmdListMap[c.Brand+c.Type+c.VersionRun] = c
		//fmt.Printf("%v\n",c)
	}

	errListRows, err := f.GetRows("Sheet3")
	if err != nil {
		fmt.Println(err)
		return
	}
	errListMap := make(map[string]errList)
	for key, hostCmd := range errListRows {
		if key==0{
			continue
		}
		var e errList
		for k, item := range hostCmd {
			if k==0{
				e.Brand = item
			}else if k==3{
				e.WrongStr=item
			}else if k==5{
				e.ErrorStr=item
			}
		}

		//var cmdStr map[string]string
		//if err := json.Unmarshal([]byte(e.ErrorStr), &cmdStr); err != nil {
		//	//命令解析不到不再执行
		//	fmt.Printf("%v解析json命令组失败\n",e.Brand)
		//	continue
		//}
		errListMap[e.Brand] = e
		fmt.Printf("%#v\n",e)
	}
}