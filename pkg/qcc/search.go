package qcc

import (
	"fmt"
	"go-qcc/pkg/util"
	"net/url"
	"regexp"
	"strings"
)

const (
	DOMAIN   = "https://www.qcc.com"
	API_PATH = "/api/search/searchMind?mindKeyWords=true&mindType=9&pageSize=5&person=true&searchKey=%s&suggest=true"
)

type out struct {
	Name  string
	KeyNo string
}

func (c *Client) Search(companyName string) (ret []*out, err error) {

	result := new(searchMindResult)

	path := fmt.Sprintf(API_PATH, url.QueryEscape(companyName))

	key, value, err := Calculate(strings.ToLower(path), ``, c.tid)

	if err != nil {
		return
	}

	headers := map[string]string{
		key: value,
	}

	response, err := c.Client.R().SetResult(result).SetHeaders(headers).Get(DOMAIN + path)
	if err != nil {
		return
	}

	if len(result.List) == 0 {
		resp := response.String()
		if strings.Contains(resp, "verify.qcc.com") {
			r, err := regexp.Compile(`<iframe src="(.*?)"`)
			if err != nil {
				return nil, err
			}
			fmt.Println(r.FindString(resp))
		} else if strings.Contains(resp, "<title>会员登录 - 企查查</title>") {
			err = util.NewError(-1, "Search: "+"login required")
			return
		} else if strings.Contains(resp, "<h3>服务器有点累，需要喘口气…</h3>") {
			err = util.NewError(-1, "Search: "+"the parameter is incorrect")
			return
		} else {
			fmt.Println(response.String())
		}
	}

	for _, i := range result.List {
		if i.Reason != "名称匹配" || i.Name0 != companyName {
			continue
		}
		ret = append(ret, &out{
			Name:  i.Name0,
			KeyNo: i.KeyNo,
		})
	}
	return
}

type searchMindResult struct {
	List []struct {
		KeyNo     string `json:"KeyNo"`
		QccCode   string `json:"QccCode"`
		Name      string `json:"Name"`
		OperName  string `json:"OperName"`
		ImageURL  string `json:"ImageUrl"`
		IsAuth    bool   `json:"IsAuth"`
		AuthLevel int    `json:"AuthLevel"`
		IsHide    bool   `json:"isHide"`
		OperInfo  string `json:"OperInfo"`
		Name0     string `json:"name"`
		Reason    string `json:"Reason"`
	} `json:"list"`
	OtherList []struct {
		ImageURL  string `json:"ImageUrl"`
		KeyNo     string `json:"KeyNo"`
		Name      string `json:"Name"`
		Type      string `json:"Type"`
		Name0     string `json:"name"`
		SearchKey string `json:"searchKey"`
		Count     int    `json:"count"`
		TagText   string `json:"tagText"`
	} `json:"otherList"`
}
