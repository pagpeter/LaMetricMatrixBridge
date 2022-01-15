package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func handle(err error) {
	if err != nil {
		log.Fatalf(err)
	}
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func makeRequest(url string) string {
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func getTimeLocation(locations []string) string {
	for _, location := range locations {
		res := makeRequest(location)
		if strings.Contains(res, "LaMetric") {
			re := regexp.MustCompile(`https:\/\/(.+?):443`)
			return re.FindStringSubmatch(res)[1]
		}
	}
	return ""
}

func getAllDevices() []string {
	var locations []string

	ssdpDiscover := "M-SEARCH * HTTP/1.1\r\n" +
		"HOST: 239.255.255.250:1900\r\n" +
		"MAN: \"ssdp:discover\"\r\n" +
		"MX: 1\r\n" +
		"ST: ssdp:all\r\n" +
		"\r\n"

	r := regexp.MustCompile(`(?i)location:[ ]*(.+)\r\n`)

	addr := &net.UDPAddr{IP: net.IPv4(239, 255, 255, 250), Port: 1900}
	ourAddr, _ := net.ResolveUDPAddr("udp", "192.168.178.30:10000")

	conn, err := net.ListenUDP("udp", ourAddr)
	defer conn.Close()
	handle(err)
	conn.WriteToUDP([]byte(ssdpDiscover), addr)
	handle(err)

	go func() {
		for {
			buf := make([]byte, 2048)
			n, _, _ := conn.ReadFromUDP(buf)
			matches := r.FindStringSubmatch(string(buf[:n]))
			if len(matches) > 1 {
				locations = append(locations, matches[1])
			}
		}
	}()
	time.Sleep(2 * time.Second)
	conn.Close()
	return removeDuplicateStr(locations)

}

func GetIPAddress() string {
	devices := getAllDevices()
	return getTimeLocation(devices)
}
