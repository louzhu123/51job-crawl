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
	"github.com/louzhu123/gcrawl"
)

func main() {
	var json string
	gcrawl := gcrawl.New51job()
	json = gcrawl.Where("page", 1).Where("salary", "0-10000").Where("city", "广州").Get()
	result := gjson.Get(json, "engine_search_result")
	fmt.Println(result)
}

```

