package gcrawl

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)



type _51job struct {
	url string
	whereConditions []map[string]interface{}
	Crawl
}

func NewCrawl() *_51job {
	return &_51job{}
}

func (c *_51job) Where(query string,value interface{}) *_51job {
	c.whereConditions = append(c.whereConditions,map[string]interface{}{"query":query,"value":value})
	return c
}

func (c *_51job) processUrl() {
	city := "000000"
	search := " "
	page := 1
	salary,workyear,companyType,degree,jobType,companySize := "99","99","99","99","99","99"

	//TODO 参数处理
	for _,item := range c.whereConditions {
		switch item["query"] {
		case "page":
			page, _ = item["value"].(int)
		case "city":
			fmt.Println(item["value"])
		case "companyType":
			fmt.Println(item["value"])
		case "companySize":
			fmt.Println(item["value"])
		case "degree":
			fmt.Println(item["value"])
		case "jobType":
			fmt.Println(item["value"])
		case "salary":
			fmt.Println(item["value"])
		case "search":
			fmt.Println(item["value"])
		}
	}
	url := "https://search.51job.com/list/"+city+",000000,0000,00,9,"+salary+"," + search + ",2," + strconv.Itoa(page) + ".html?lang=c&postchannel=0000&workyear=" + workyear + "&cotype=" + companyType + "&degreefrom=" + degree + "&jobterm="+jobType+"&companysize="+companySize+"&ord_field=0&dibiaoid=0&line=&welfare="
	c.url = url
}

func (c *_51job) Links() []string {
	c.processUrl()
	res,err:= http.Get(c.url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	links := []string{}
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.Contains(text,"window.__SEARCH_RESULT__ = ") {
			json := strings.Split(text,"window.__SEARCH_RESULT__ = ")[1]
			value := gjson.Get(json, "engine_search_result.#.job_href")
			for _,value := range value.Array()  {
				fmt.Println(value.Value())
				s,_ := value.Value().(string)
				links = append(links,s)
			}
		}
	})
	return links
}




