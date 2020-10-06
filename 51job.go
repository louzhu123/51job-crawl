package gcrawl

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/axgle/mahonia"
)

type _51job struct {
	url             string
	whereConditions []map[string]interface{}
}

func New51job() *_51job {
	return &_51job{}
}

func (c *_51job) Where(params ...interface{}) *_51job {
	switch params[0].(type) {
	case map[string]interface{}:
		for query,value := range params[0].(map[string]interface{}) {
			c.whereConditions = append(c.whereConditions, map[string]interface{}{"query": query, "value": value})
		}
	case string:
		c.whereConditions = append(c.whereConditions, map[string]interface{}{"query": params[0], "value": params[1]})
	}
	return c
}

func (c *_51job) processUrl() {
	parameters := map[string]string{
		"city" : "000000",
		"search" : "+",
		"page" : "1",
		"salary" : "99",
		"workyear" : "99",
		"cotype" : "99",
		"degreefrom" : "99",
		"companysize" : "99",
		"position" : "0000",
	}
	for _, condition := range c.whereConditions {
		switch condition["query"] {
		case "page":
			page,_ := condition["value"].(int)
			parameters["page"] = strconv.Itoa(page)
		case "city":
			cityList := []map[string]string{
				{"displayName": "北京", "value": "010000"},
				{"displayName": "上海", "value": "020000"},
				{"displayName": "广州", "value": "030200"},
				{"displayName": "深圳", "value": "040000"},
				{"displayName": "武汉", "value": "180200"},
				{"displayName": "西安", "value": "200200"},
				{"displayName": "杭州", "value": "080200"},
				{"displayName": "南京", "value": "070200"},
				{"displayName": "成都", "value": "090200"},
				{"displayName": "重庆", "value": "060000"},
				{"displayName": "东莞", "value": "030800"},
				{"displayName": "大连", "value": "230300"},
				{"displayName": "沈阳", "value": "230200"},
				{"displayName": "苏州", "value": "070300"},
				{"displayName": "昆明", "value": "250200"},
				{"displayName": "长沙", "value": "190200"},
				{"displayName": "合肥", "value": "150200"},
				{"displayName": "宁波", "value": "080300"},
				{"displayName": "郑州", "value": "170200"},
				{"displayName": "天津", "value": "050000"},
				{"displayName": "青岛", "value": "120300"},
				{"displayName": "济南", "value": "120200"},
				{"displayName": "哈尔滨", "value": "220200"},
				{"displayName": "长春", "value": "240200"},
				{"displayName": "福州", "value": "110200"},
			}
			if value,flag := InList(cityList,condition["value"].(string)); flag == true {
				parameters["city"] =  value
			}else {
				println("city仅包含以下值")
				for _,item := range cityList {
					print(item["displayName"] + " ")
				}
				log.Fatal()
			}
		case "cotype":
			cotypeList := []map[string]string{
				{"displayName": "国企", "value": "04"},
				{"displayName": "外资（欧美）", "value": "01"},
				{"displayName": "外资（非欧美）", "value": "02"},
				{"displayName": "上市公司", "value": "10"},
				{"displayName": "合资", "value": "03"},
				{"displayName": "民营公司", "value": "05"},
				{"displayName": "外企代表处", "value": "06"},
				{"displayName": "政府机关", "value": "07"},
				{"displayName": "事业单位", "value": "08"},
				{"displayName": "非盈利组织", "value": "09"},
				{"displayName": "创业公司", "value": "11"},
			}
			if value,flag := InList(cotypeList,condition["value"].(string)); flag == true {
				parameters["cotype"] =  value
			}else {
				println("cotype仅包含以下值")
				for _,item := range cotypeList {
					print(item["displayName"] + " ")
				}
				log.Fatal()
			}
		case "companysize":
			companysizeList := []map[string]string{
				{"displayName": "少于50人", "value": "01"},
				{"displayName": "50-150人", "value": "02"},
				{"displayName": "150-500人", "value": "03"},
				{"displayName": "500-1000人", "value": "04"},
				{"displayName": "1000-5000人", "value": "05"},
				{"displayName": "5000-10000人", "value": "06"},
				{"displayName": "10000人以上", "value": "07"},
			}
			if value,flag := InList(companysizeList,condition["value"].(string)); flag == true {
				parameters["cotycompanysizepe"] =  value
			}else {
				println("companysize仅包含以下值")
				for _,item := range companysizeList {
					print(item["displayName"] + " ")
				}
				log.Fatal()
			}
		case "degreefrom":
			degreefromList := []map[string]string{
				{"displayName": "初中及以下", "value": "01"},
				{"displayName": "高中/中技/中专", "value": "02"},
				{"displayName": "大专", "value": "03"},
				{"displayName": "本科", "value": "04"},
				{"displayName": "硕士", "value": "05"},
				{"displayName": "博士", "value": "06"},
				{"displayName": "无学历要求", "value": "07"},
			}
			if value,flag := InList(degreefromList,condition["value"].(string)); flag == true {
				parameters["cotype"] = value
			}else {
				println("degreefrom仅包含以下值")
				for _,item := range degreefromList {
					print(item["displayName"] + " ")
				}
				log.Fatal()
			}
		case "position":
			degreefromList := []map[string]string{
				{"displayName": "后端开发", "value": "0100"},
				{"displayName": "移动开发", "value": "7700"},
				{"displayName": "前端开发", "value": "7200"},
				{"displayName": "人工智能", "value": "7300"},
				{"displayName": "游戏", "value": "7800"},
				{"displayName": "设计", "value": "7400"},
				{"displayName": "测试", "value": "2700"},
				{"displayName": "运维/技术支持", "value": "7900"},
				{"displayName": "数据", "value": "7500"},
				{"displayName": "产品", "value": "6600"},
				{"displayName": "运营", "value": "8000"},
				{"displayName": "技术管理", "value": "2600"},
				{"displayName": "电子商务", "value": "6100"},
				{"displayName": "半导体/芯片", "value": "6700"},
				{"displayName": "电子/电器/仪表仪器", "value": "2900"},
				{"displayName": "通信技术开发及应用", "value": "2800"},
			}
			if value,flag := InList(degreefromList,condition["value"].(string)); flag == true {
				parameters["position"] = value
			}else {
				println("position仅包含以下值")
				for _,item := range degreefromList {
					print(item["displayName"] + " ")
				}
				log.Fatal()
			}
		case "salary":
			parameters["salary"] = condition["value"].(string)
		case "search":
			parameters["search"] = condition["value"].(string)
		}
	}
	url := "https://search.51job.com/list/" + parameters["city"] + ",000000," + parameters["position"] + ",00,9," + parameters["salary"] + "," + parameters["search"] + ",2," + parameters["page"] + ".html?lang=c&postchannel=0000&workyear=" + parameters["workyear"] + "&cotype=" + parameters["cotype"] + "&degreefrom=" + parameters["degreefrom"]+ "&jobterm=99&companysize=" + parameters["companysize"] + "&ord_field=0&dibiaoid=0&line=&welfare="
	c.url = url
}

func (c *_51job) Get() string {
	c.processUrl()
	client := &http.Client{}
	req,_ :=  http.NewRequest("GET",c.url,nil)
	req.Header.Add("Accept","application/json, text/javascript, */*; q=0.01")
	resp,_ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	enc := mahonia.NewDecoder("gbk")
	return enc.ConvertString(string(body))
}
