package pushers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Send(url string, body interface{}) {

	var res *bytes.Buffer

	switch body := body.(type) {
	case string:
		res = bytes.NewBufferString(body)
	default:
		res1, _ := json.Marshal(body)
		res = bytes.NewBuffer(res1)
	}

	client := http.DefaultClient
	resp, err := client.Post(url, "application/json", res)
	fmt.Println(resp, err)
}
