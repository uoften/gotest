package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"sort"
)
type Person struct {
	Name string
	Age  int
}
type Age []int

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func main()  {
	//fmt.Println("%h%m%d.log")
	//fmt.Printf("%v\n",helper.GetIpByString("asd127.0.0.333a"))
	//
	//reServiceArr1 := strings.Split("8004", " ")
	//reServiceArr2 := strings.Split("8001 8006", " ")
	//reServiceArr3 := strings.Split("8001 8006 src-port 1 65535", " src-port ")
	//fmt.Printf("%#v\n",reServiceArr1)
	//fmt.Printf("%#v\n",reServiceArr2)
	//fmt.Printf("%#v\n",reServiceArr3)
	//
 	p1 :=new(Person)
 	p2 :=new(Person)
 	fmt.Println(p1==p2)

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	//用余数分库分表的分布结果
	var db,tb,i int
	var m = make(map[int][]int)
	for i=1;i<=90;i++ {
		db=i%10+1
		tb=i%9+1
		m[db] = append(m[db],tb)
	}
	for k,v:=range m{
		fmt.Println(k, " ----> ", v)
	}

	//re := regexp.MustCompile("(\\S+) dst-port")
	//fmt.Printf("%#v",re.FindString("  udp dst-port 3000 3388\n  tcp dst-port 3390 65535"))
	//
	//config_bytes, err := ioutil.ReadFile("./hillstone_ip_route.txt")
	//if err != nil {
	//	fmt.Printf("读取对应的配置文件失败")
	//	return
	//}
	//configStr := string(config_bytes)
	//configStr = strings.ReplaceAll(configStr, "\r\n", "\n")
	//configStr = strings.ReplaceAll(configStr, "\r", "\n")
	//
	//// 获取解析规则
	//rule := AnalysisRule{}
	//bytes, err := ioutil.ReadFile("./hillstone.json")
	//if err != nil {
	//	fmt.Printf("读取对应的解析规则文件失败")
	//	return
	//}
	//err = json.Unmarshal([]byte(bytes), &rule)
	//if err != nil {
	//	fmt.Printf("获取解析规则json格式失败")
	//	return
	//}
	//var line int
	//unAnalysisLine := make(map[int]string)
	//f := func(rule []string, lineConf string) (reData []string, reIndex int) {
	//	// 正则匹配获取对应数据
	//	for _, ruleString := range rule {
	//		reData = regexp.MustCompile(ruleString).FindStringSubmatch(lineConf)
	//		reIndex += 1
	//		if reData == nil {
	//			continue
	//		} else {
	//			return reData[1:], reIndex
	//		}
	//	}
	//	return
	//}
	//handRouteRe1 := func(reRoute []string) (ret RoutePolicy) {
	//	//['0.0.0.0', '0', 'O_ASE2', '150', '1', '10.54.255.70', 'Reth255']   "0.0.0.0/0          O_ASE2  150 1           10.54.255.70    Reth255"
	//	//ret.IpSegment = fmt.Sprintf("%s/%s", reRoute[0], reRoute[1])
	//	//ret.Mask, _ = strconv.Atoi(reRoute[1])
	//	//ret.Protocol = reRoute[2]
	//	//ret.Priority, _ = strconv.Atoi(reRoute[3])
	//	//ret.Cost, _ = strconv.Atoi(reRoute[4])
	//	//ret.NextHop = reRoute[5]
	//	//ret.Interface = reRoute[6]
	//	return
	//}
	//configList := strings.Split(configStr, "\n")
	//var mRoutePolicy []RoutePolicy
	//for _, lineConf := range configList {
	//	// 处理每一行的数据
	//	line += 1
	//	if strings.TrimSpace(lineConf) == "" {
	//		continue
	//	}
	//	lineConf = lineConf + "\n"
	//	reRoute, reIndex := f(rule.RoutePolicyRule, lineConf)
	//	if reRoute != nil {
	//		fmt.Printf("\n%#v\n",lineConf)
	//		fmt.Printf("%#v\n",rule.RoutePolicyRule)
	//		fmt.Printf("%#v\n\n\n",reRoute)
	//		var reHandRoutePolicy RoutePolicy
	//		switch reIndex {
	//		case 1:
	//			reHandRoutePolicy = handRouteRe1(reRoute)
	//		case 2:
	//			reHandRoutePolicy = handRouteRe1(reRoute)
	//		default:
	//			fmt.Printf("【配置解析-路由策略】 路由策略正则数据处理方法缺失")
	//			unAnalysisLine[line] = lineConf
	//		}
	//		if reHandRoutePolicy.IpSegment != "" {
	//			mRoutePolicy = append(mRoutePolicy, reHandRoutePolicy)
	//		}
	//		continue
	//	}
	//	unAnalysisLine[line] = lineConf
	//}
}











// route_policy model
type RoutePolicy struct {
	Id          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	DeviceIp    string        `json:"device_ip" bson:"device_ip"`
	IpSegment   string        `json:"ip_segment" bson:"ip_segment"`
	StartIp     int           `json:"start_ip" bson:"start_ip"`
	EndIp       int           `json:"end_ip" bson:"end_ip"`
	Mask        int           `json:"mask" bson:"mask"`
	Protocol    string        `json:"protocol" bson:"protocol"`
	Priority    int           `json:"priority" bson:"priority"`
	Cost        int           `json:"cost" bson:"cost"`
	NextHop     string        `json:"next_hop" bson:"next_hop"`
	Interface   string        `json:"interface" bson:"interface"`
	VpnInstance string        `json:"vpn_instance" bson:"vpn_instance"`
	UpdateTime  string        `json:"update_time"  bson:"update_time"`
}


type AnalysisRule struct {
	Address         []string        `json:"address"`
	Service         []string        `json:"service"`
	Time            []string        `json:"time"`
	Zone            []string        `json:"zone"`
	AclTime         []string        `json:"acl_time"`
	Acl             []string        `json:"acl"`
	PolicyRoute     []string        `json:"policy_route"`
	Policy          []string        `json:"policy"`
	InterfacePort   []string        `json:"interface_port"`
	RoutePolicyRule []string        `json:"route_policy_rule"`
	AddressRule     AddressRule     `json:"address_rule"`
	ServiceRule     ServiceRule     `json:"service_rule"`
	TimeRule        TimeRule        `json:"time_rule"`
	ZoneRule        ZoneRule        `json:"zone_rule"`
	AclTimeRule     AclTimeRule     `json:"acl_time_rule"`
	AclRule         AclRule         `json:"acl_rule"`
	PolicyRouteRule RolicyRouteRule `json:"policy_route_rule"`
	PolicyRule      PolicyRule      `json:"policy_rule"`
	InterfaceRule   InterfaceRule   `json:"interface_rule"`
}

type AddressRule struct {
	Name    []string `json:"name"`
	Address []string `json:"address"`
}

type ServiceRule struct {
	Name    []string `json:"name"`
	Service []string `json:"service"`
}

type TimeRule struct {
	Name []string `json:"name"`
	Time []string `json:"time"`
}

type ZoneRule struct {
	Name      []string `json:"name"`
	Priority  []string `json:"priority"`
	Interface []string `json:"interface"`
}

type AclTimeRule struct {
	Time []string `json:"time"`
}

type AclRule struct {
	Number     []string `json:"number"`
	Desc       []string `json:"desc"`
	RuleObject []string `json:"rule_object"`
}

type RolicyRouteRule struct {
	Name      []string `json:"name"`
	AclNumber []string `json:"acl_number"`
	NextHop   []string `json:"next_hop"`
}

type PolicyRule struct {
	Name             []string `json:"name"`
	Desc             []string `json:"desc"`
	SrcZone          []string `json:"src_zone"`
	DestZone         []string `json:"dest_zone"`
	SrcAddress       []string `json:"src_address"`
	SrcGroupAddress  []string `json:"src_group_address"`
	DestAddress      []string `json:"dest_address"`
	DestGroupAddress []string `json:"dest_group_address"`
	GroupService     []string `json:"group_service"`
	ServicePort      []string `json:"service_port"`
	Protocol         []string `json:"protocol"`
	GroupTime        []string `json:"group_time"`
	Action           []string `json:"action"`
	Status           []string `json:"status"`
}

type InterfaceRule struct {
	Name        []string `json:"name"`
	VpnInstance []string `json:"vpn_instance"`
	InterfaceIp []string `json:"interface_ip"`
}


type AnalysisParams struct {
	DeviceBrand string `json:"device_brand"`
	DeviceType  string `json:"device_type"`
	Version     string `json:"version"`
	DeviceIp    string `json:"device_ip"`
	FilePath    string `json:"file_path"`
	FileType    string `json:"file_type"`
	VpnInstance string `json:"vpn_instance"`
}

type GetHomePageDataReq struct {
	UserName string `json:"username"`
}

type GetHomePageDataResp struct {
	TotalDeviceCount    int        `json:"total_device_count"`
	FireWallDeviceCount int        `json:"firewall_device_count"`
	RouteDeviceCount    int        `json:"route_device_count"`
	SwitchDeviceCount   int        `json:"switch_device_count"`
	TotalPolicyCount    int        `json:"total_policy_count"`
	RiskPolicyCount     int        `json:"risk_policy_count"`
	ProblemPolicyCount  int        `json:"problem_policy_count"`
	TodayFireWallCount  int        `json:"today_firewall_count"`
	TotalTaskCount      int        `json:"total_task_count"`
	WaitTaskCount       int        `json:"wait_task_count"`
	FailTaskCount       int        `json:"fail_task_count"`
	WaitTaskList        []TaskList `json:"wait_task_list"`
	FailTaskList        []TaskList `json:"fail_task_list"`
}

type TaskList struct {
	TaskName   string `json:"task_name"`
	ReqUser    string `json:"req_user"`
	UpdateTime string `json:"update_time"`
}