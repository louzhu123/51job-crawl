# gcrawl-51job爬虫



## Installation

```
$ go get github.com/louzhu123/gcrawl
```

## Examples

```go
package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/louzhu123/gcrawl"
)

func main() {
	gcrawl := gcrawl.New51job()
	conditions := map[string]interface{}{
		"page":   2,
		"salary": "0-10000",
		"city":   "成都",
	}
	json := gcrawl.Where(conditions).Get()
	result := gjson.Get(json, "engine_search_result")
	fmt.Println(result)
}

```



## Where Paramter

| key         | value                                                        |
| ----------- | ------------------------------------------------------------ |
| page        | int                                                          |
| city        | 热门城市 [ 北京，上海，广州，深圳，武汉，西安，杭州，南京，成都，重庆，东莞，大连，沈阳，苏州，昆明，长沙，合肥，宁波，郑州，天津，青岛，济南，哈尔滨，长春，福州 ] |
| cotype      | [ 国企，外资（欧美），外资（非欧美），上市公司，合资，民营公司，外企代表处，政府机关，事业单位，非盈利组织，创业公司 ] |
| companysize | [ 少于50人，50-150人，150-500人，500-1000人，1000-5000人，5000-10000人，10000人以上 ] |
| degreefrom  | [ 初中及以下，高中/中技/中专，大专，本科，硕士，博士，无学历要求 ] |
| salary      | n-m，例如0-10000                                             |
| search      | 搜索词                                                       |

