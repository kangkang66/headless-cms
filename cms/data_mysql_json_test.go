package cms

import (
	"context"
	"fmt"
	"testing"
)

func TestDataMysqlJson_Insert(t *testing.T) {
	c := newDataMysqlJson()
	ctx := context.Background()
	source := map[string]interface{}{
		"name":"taobao",
	}
	fmt.Println(c.Insert(ctx, "source",source))

	likes := []map[string]interface{}{
		{
			"name":"apple",
			"buy_source":1,
		},
		{
			"name":"banana",
			"buy_source":1,
		},
	}
	for _,v := range likes {
		fmt.Println(c.Insert(ctx, "like",v))
	}

	user := map[string]interface{}{
		"name":"kang",
		"likes":[]int{2,3},
		"user_source":1,
	}
	c.Insert(ctx, "user",user)
}
