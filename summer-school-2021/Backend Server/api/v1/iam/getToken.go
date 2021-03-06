package iam

import (
	"backend/api/v1/auth"
	errorcheck "backend/api/v1/error-check"
	"backend/api/v1/requester"
	"bytes"
	"fmt"
	"github.com/borodun/SberLab/core"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func MakeRequestBody(reqUrl string, bodyStr string) (string, string) {
	signer := core.Signer{
		Key:    os.Getenv("ACCESSKEY"),
		Secret: os.Getenv("SECRETKEY"),
	}

	req, err := http.NewRequest("POST", reqUrl, ioutil.NopCloser(bytes.NewBuffer([]byte(bodyStr))))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-Project-Id", auth.InfoAuth.ProjectID)

	err = signer.Sign(req)
	if err != nil {
		return "", ""
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
	return string(resp.Header.Get("X-Subject-Token")), string(body)
}

func GetToken(username string, password string, domainName string) string {
	var reqUrl = fmt.Sprintf("https://iam.ru-moscow-1.hc.sbercloud.ru/v3/auth/tokens")
	tokenRequest := fmt.Sprintf(`{
  "auth": {
    "identity": {
      "methods": [
        "password"
      ],
      "password": {
        "user": {
          "name": "%s",
          "password": "%s",
          "domain": {
            "name": "%s"
          }
        }
      }
    },
    "scope": {
      "domain": {
        "name": "%s"
      }
    }
  }
}`, username, password, domainName, domainName)

	token, errStr := MakeRequestBody(reqUrl, tokenRequest)
	str := errorcheck.CheckError(errStr)

	requester.Tok.Token = token
	requester.Tok.CreationTime = time.Now()
	return str
}
