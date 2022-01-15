package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type LaMetric struct {
	Api_Key string
	IP      string
}

const port = "8080"

func (l *LaMetric) MakeRequest(method, endpoint string, json []byte) (string, error) {
	url := "http://" + l.IP + ":" + port + "/api/v2/" + endpoint
	req, err := http.NewRequest(method, url, bytes.NewBuffer(json))
	if err != nil {
		return "", err
	}
	// fmt.Println(url)
	req.Header.Set("Authorization", "Basic "+l.Api_Key)
	client := &http.Client{}
	res, err := client.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (l *LaMetric) EncodeApiKey(key string) string {
	data := "dev:" + key
	return b64.StdEncoding.EncodeToString([]byte(data))
}

func (l *LaMetric) GetState() (DeviceState, error) {
	var d DeviceState

	res, err := l.MakeRequest("GET", "device", nil)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal([]byte(res), &d)
	return d, err
}

func (l *LaMetric) SendNotification(n Notification) (string, error) {
	json, err := json.Marshal(n)
	if err != nil {
		return "", err
	}
	return l.MakeRequest("POST", "device/notifications", json)
}

func (l *LaMetric) GetNotifications() ([]Notification, error) {
	var n []Notification

	res, err := l.MakeRequest("GET", "device/notifications", nil)
	if err != nil {
		return n, err
	}
	err = json.Unmarshal([]byte(res), &n)
	return n, err
}

func (l *LaMetric) DeleteNotification(id string) (Success, error) {
	var s Success

	res, err := l.MakeRequest("DELETE", "device/notifications/"+id, nil)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal([]byte(res), &s)
	return s, err
}

func (l *LaMetric) GetApps() (Apps, error) {
	var a Apps

	res, err := l.MakeRequest("GET", "device/apps", nil)
	if err != nil {
		return a, err
	}
	err = json.Unmarshal([]byte(res), &a)
	return a, err
}

func (l *LaMetric) NextApp() (string, error) {
	return l.MakeRequest("PUT", "device/apps/next", nil)
}

func (l *LaMetric) PrevApp() (string, error) {
	return l.MakeRequest("PUT", "device/apps/prev", nil)
}

func (l *LaMetric) GetApp(id string) (App, error) {
	var a App

	res, err := l.MakeRequest("GET", "device/apps/"+id, nil)
	if err != nil {
		return a, err
	}
	err = json.Unmarshal([]byte(res), &a)
	return a, err
}
