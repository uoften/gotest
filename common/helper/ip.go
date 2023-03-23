package helper

import "regexp"

func GetIpByString(str string) (ip string) {
	regStr := `\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}`
	reg, _ := regexp.Compile(regStr)
	ip = string(reg.Find([]byte(str)))
	return
}
