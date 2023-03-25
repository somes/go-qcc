package main

import (
	"go-qcc/pkg/qcc"
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {

	var companyName = ""

	headers := map[string]string{
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
	}
	client := &qcc.Client{
		Client: resty.New().SetHeaders(headers),
	}
	_, _, err := client.GetPidAndTid()
	if err != nil {
		log.Println(err)
		return
	}

	result, err := client.Search(companyName)
	if err != nil {
		log.Println(err)
		return
	}
	if len(result) == 0 {
		log.Println("query result is empty")
		return
	}
	for _, v := range result {
		log.Println(v.Name)
		log.Println(v.KeyNo + "\n")
	}

}
