package gcrawl

type Crawl interface{
	NewCrawl(interface{}) interface{}
	Links() []string
}



