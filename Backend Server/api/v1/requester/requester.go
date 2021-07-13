package requester

import (
	"backend/api/configuration"
	"backend/api/v1/auth"
	"bytes"
	"fmt"
	"github.com/borodun/SberLab/core"
	"io"
	"io/ioutil"
	"net/http"
)

var Token string

func MakeRequest(reqUrl string) string {
	/*signer := core.Signer{
		Key:    auth.InfoAuth.Signer.AKey,
		Secret: auth.InfoAuth.Signer.SKey,
	}*/
	config, errBool := configuration.LoadConfig("../../")
	if errBool {
		return "{error: 'Wrong config'}"
	}

	signer := core.Signer{
		Key:    config.AccessKey,
		Secret: config.SecretKey,
	}

	req, err := http.NewRequest("GET", reqUrl, ioutil.NopCloser(bytes.NewBuffer([]byte(""))))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-Project-Id", auth.InfoAuth.ProjectID)
	req.Header.Add("X-Auth-Token", Token)
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

	//fmt.Println(string(body))
	return string(body)
}
