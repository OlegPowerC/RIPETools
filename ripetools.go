package RIPEtools

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const RIPEWHOIS_URL = "https://stat.ripe.net/data/whois/data.json?resource="

type RipeData struct {
	Records [][](map[string]string)
}

type RipeDataExt struct {
	Version  string     `json:"version"`
	Messages [][]string `json:"messages"`
	Data     RipeData   `json:"data"`
}

func GetCountry(Ipaddr string) (err error, country string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	var resp *http.Response

	urlstring := fmt.Sprintf("%s%s", RIPEWHOIS_URL, Ipaddr)

	resp, err = client.Get(urlstring)
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()
	var Ddata RipeDataExt
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return err, ""
		}

		UmErr := json.Unmarshal(bodyBytes, &Ddata)

		if UmErr == nil {
			for _, rec := range Ddata.Data.Records {
				for _, recin := range rec {
					if vl, ok := recin["key"]; ok {
						if strings.ToLower(vl) == "country" {
							if vl, ok = recin["value"]; ok {
								return nil, vl
							}
						}
					}
				}
			}
		} else {
			return err, ""
		}
	}
	return fmt.Errorf("%s", "country not found"), ""
}
