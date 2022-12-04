package myip

import (
	"io/ioutil"
	"log"
	"net/http"
)

// GetMyIp print my public ip
func GetMyIp() {

	resp, err := http.Get("https://ipinfo.io/ip")

	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("myip - " + string(data))
}
