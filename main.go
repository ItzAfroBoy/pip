package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type IPType struct {
	QUERY  string
	ISP    string
	CITY   string
	ZIP    string
	REGION string
	CC     string `json:"CountryCode"`
}

type DNSType struct {
	DNS struct {
		GEO string
		IP  string
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	getLocation := flag.Bool("location", false, "Get location data")
	getDNS := flag.Bool("dns", false, "Get DNS data")
	getISP := flag.Bool("isp", false, "Get ISP data")
	flag.Parse()

	var IPData IPType
	var DNSData DNSType

	res, err := http.Get("http://ip-api.com/json/?fields=8758")
	check(err)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	check(err)
	err = json.Unmarshal(body, &IPData)
	check(err)

	DNSRes, err := http.Get("http://edns.ip-api.com/json/")
	check(err)
	defer DNSRes.Body.Close()
	DNSBody, err := io.ReadAll(DNSRes.Body)
	check(err)
	err = json.Unmarshal(DNSBody, &DNSData)
	check(err)

	IP := IPData.QUERY
	Location := fmt.Sprintf("%s [%s] | %s | %s", IPData.CITY, IPData.ZIP, IPData.REGION, IPData.CC)
	ISP := IPData.ISP
	DNS := fmt.Sprintf("%s [%s]", DNSData.DNS.GEO, DNSData.DNS.IP)

	fmt.Println(IP)
	if *getLocation {
		fmt.Println(Location)
	}
	if *getISP {
		fmt.Println(ISP)
	}
	if *getDNS {
		fmt.Println(DNS)
	}
}
