package service

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"githun.com/captainstdin/airenche-temp/service/consts"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type envSiteStruct struct {
	Sid       string `json:"sid"`       //场地sid
	SecretKey string `json:"secretKey"` //密钥
}

var EnvSite envSiteStruct

func LoadEnv() error {

	configFile, err := os.Open("./env.json")

	if err == nil {
		defer configFile.Close()
		var FileJson envSiteStruct
		decoder := json.NewDecoder(configFile)
		err = decoder.Decode(&FileJson)
		EnvSite = FileJson
	}

	if EnvSite.Sid == "" {
		EnvSite.Sid = os.Getenv("sid")
	}
	if EnvSite.SecretKey == "" {
		EnvSite.SecretKey = os.Getenv("secretKey")
	}
	return nil
}

func CallRemoteCurl(path string, data map[string]string) ([]byte, error) {
	const Tk = "5381"
	// Generate the timestamp and secret key
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	//secretKey := "8212c97494d7cbd0cf8809709e25ce6c"

	// Generate the sign using the timestamp and secret key
	sign := fmt.Sprintf("%x", md5.Sum([]byte(timestamp+EnvSite.SecretKey)))

	// Set the cookies for the request
	cookies := fmt.Sprintf("stamp=%s; sign=%s", timestamp, sign)
	// Send the request

	var req *http.Request
	var err error
	if len(data) == 0 {
		req, err = http.NewRequest("GET", fmt.Sprintf("%s%s?tk=%s&sid=%s", consts.RencheEndpoint, path, Tk, EnvSite.Sid), nil)
		if err != nil {
			return nil, err
		}
	} else {
		// Send a POST request with the data as the request body
		values := url.Values{}
		for k, v := range data {
			values.Add(k, v)
		}
		req, err = http.NewRequest("POST", fmt.Sprintf("%s%s?tk=%s&sid=%s", consts.RencheEndpoint, path, Tk, EnvSite.Sid), strings.NewReader(values.Encode()))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	req.Header.Set("Cookie", cookies)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		// Handle error
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Handle error
		return nil, err
	}
	return body, nil
}
