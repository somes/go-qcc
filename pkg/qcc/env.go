package qcc

import (
	"regexp"

	"go-qcc/pkg/util"
)

/*
get winodw.pid window.tid
*/
func (c *Client) GetPidAndTid() (pid, tid string, err error) {
	response, err := c.Client.R().Get("https://www.qcc.com/")
	if err != nil {
		return
	}

	// <script>window.pid=''; window.tid=''</script>
	regex := `window.pid='(.*?)'; window.tid='(.*?)'`
	r, err := regexp.Compile(regex)
	if err != nil {
		err = util.NewError(-1, "parsing regular expression failed")
		return
	}
	result := r.FindStringSubmatch(response.String())
	if len(result) != 3 {
		err = util.NewError(-1, "failed to get pid tid")
		return
	}

	pid = result[1]
	tid = result[2]

	c.tid = tid

	return
}
