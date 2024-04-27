package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/spf13/cast"
	"strings"
)

type List struct {
	ID   string
	Name string
}

func test() {
	client, _ := elasticsearch.NewDefaultClient()

	list := make([]List, 0)

	// 构建批量插入请求的正文
	builder := strings.Builder{}
	for _, transaction := range list {
		builder.WriteString(`{ "index" : { "_id" : "`)
		builder.WriteString(cast.ToString(transaction.ID))
		builder.WriteString(`" } }`)
		builder.WriteString("\n")

		data, err := json.Marshal(transaction)
		if err != nil {
			fmt.Println("Error marshaling transaction:", err)
			continue
		}
		builder.Write(data)
		builder.WriteString("\n")
	}

	fmt.Println(builder.String())

	// 执行批量插入操作
	req := esapi.BulkRequest{
		Index: "test",
		Body:  strings.NewReader(builder.String()),
	}
	res, err := req.Do(context.Background(), client)
	if err != nil {
		fmt.Println("Error executing bulk request:", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Println("Bulk insert failed:", res.Status())
	} else {
		fmt.Println("Bulk insert successful")
	}
	fmt.Println(res.String())
}
