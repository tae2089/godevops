package myip

import (
	"io/ioutil"
	"log"
	"net/http"
)

// GetMyIp print my public ip
func GetMyIp() error {

	resp, err := http.Get("https://ipinfo.io/ip")

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println("myip - " + string(data))
	return nil
}
