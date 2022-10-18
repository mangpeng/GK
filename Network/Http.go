package Network

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HttpHeader struct {
	Name  string
	Value string
}

type HttpConnector struct {
}

func (c *HttpConnector) GET(urlString string, headers ...*HttpHeader) (string, error) {
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return "", err
	}

	if headers != nil {
		for _, v := range headers {
			req.Header.Add(v.Name, v.Value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return string(b), nil
}

func (c *HttpConnector) PostPlainText(urlString, body string, headers ...*HttpHeader) (string, error) {
	req, err := procPostRequest(urlString, body, headers...)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "plain/text")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func (c *HttpConnector) PostJSON(urlString string, jsonData interface{}, headers ...*HttpHeader) ([]byte, error) {
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	jsonString := string(jsonBytes)

	req, err := procPostRequest(urlString, jsonString, headers...)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func procPostRequest(urlString, body string, headers ...*HttpHeader) (*http.Request, error) {
	reqBody := bytes.NewBufferString(body)
	req, err := http.NewRequest("POST", urlString, reqBody)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for _, v := range headers {
			req.Header.Add(v.Name, v.Value)
		}
	}

	return req, nil
}
