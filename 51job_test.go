package gcrawl

import (
	"testing"

	"github.com/tidwall/gjson"
	"github.com/stretchr/testify/assert"
)

func Test_51job_Where(t *testing.T) {
	c := New51job()
	conditions := map[string]interface{}{
		"page":   2,
		"salary": "0-10000",
		"city":   "成都",
	}
	c.Where(conditions)
	assert.Len(t, c.whereConditions, 3)
}

func Test_51job_Get(t *testing.T) {
	c := New51job()
	conditions := map[string]interface{}{
		"page":   2,
		"salary": "0-10000",
		"city":   "成都",
		"position" : "后端开发",
	}
	json := c.Where(conditions).WithMoreInfo().Get()
	result := gjson.Get(json, "engine_search_result")
	assert.NotNil(t,result)
}