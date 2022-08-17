package utils

import (
	"log"
	"net"
	"strings"

	geo "github.com/oschwald/geoip2-golang"
)

type IPGeo struct {
	Reader *geo.Reader
}

func ToIP(remoteAddres string) string {
	if strings.Contains(remoteAddres, "[") {
		return strings.Split(remoteAddres, "]:")[0][1:]

	}
	return strings.Split(remoteAddres, ":")[0]
}

func (i *IPGeo) Load(path string) {
	reader, err := geo.Open(path)
	if err != nil {
		log.Println(err)
	} else {
		i.Reader = reader
	}
}

func (i *IPGeo) GetRecord(ip string) map[string]string {
	if i.Reader == nil {
		log.Println("load geolocation first")
		return nil
	}
	target := net.ParseIP(ip)
	record, err := i.Reader.Country(target)
	if err != nil {
		log.Fatal(err)
	}
	return record.Country.Names
}
