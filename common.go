package gcrawl

import (
	"github.com/axgle/mahonia"
)

// 判断参数是否符合
func InList(list []map[string]string, value string) (string,bool) {
	Inlist := false
	str := ""
	for _,item := range list {
		if item["displayName"] == value {
			Inlist = true
			str = item["value"]
		}
	}
	return str,Inlist
}

func Gbk2Utf8(gbkStr string) string  {
	enc := mahonia.NewDecoder("gbk")
	utf8Str := enc.ConvertString(gbkStr)
	return utf8Str
}