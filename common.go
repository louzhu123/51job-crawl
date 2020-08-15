package gcrawl

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