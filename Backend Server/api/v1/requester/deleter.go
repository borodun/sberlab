package requester

import (
	"backend/api/v1/auth"
	"backend/api/v1/requester/core"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func RequestDelete(reqUrl string, method string) string {
	signer := core.Signer{
		Key:    os.Getenv("ACCESSKEY"),
		Secret: os.Getenv("SECRETKEY"),
	}

	req, err := http.NewRequest(method, reqUrl, ioutil.NopCloser(bytes.NewBuffer([]byte(""))))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-Project-Id", auth.InfoAuth.ProjectID)
	req.Header.Add("X-Auth-Token", Tok.Token)
	err = signer.Sign(req)
	if err != nil {
		return ""
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}
