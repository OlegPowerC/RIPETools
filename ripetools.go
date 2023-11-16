package RIPEtools

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const RIPEWHOIS_URL = "https://stat.ripe.net/data/whois/data.json?resource="

type ripeData struct {
	Records     [][](map[string]string) `json:"records"`
	IRR_Records [][](map[string]string) `json:"irr_records"`
}

type RIPEd struct {
	RipeData ripeData
}

func (rd *RIPEd) getValByKey(key string, findin [][]map[string]string) (err error, retval string) {
	for _, rec := range findin {
		for _, recin := range rec {
			if vl, ok := recin["key"]; ok {
				if strings.ToLower(vl) == key {
					if vl, ok = recin["value"]; ok {
						return nil, vl
					}
				}
			}
		}
	}
	return fmt.Errorf("%s", "value not found"), ""
}

func (rd *RIPEd) GetCountry() (err error, country string) {
	cnterr, cnt := rd.getValByKey("country", rd.RipeData.Records)
	if cnterr != nil {
		return fmt.Errorf("%s", "country not found"), ""
	}
	return nil, cnt
}

func (rd *RIPEd) GetNetwork() (err error, Network string) {
	neterr, netval := rd.getValByKey("route", rd.RipeData.IRR_Records)
	if neterr != nil {
		return fmt.Errorf("%s", "network not found"), ""
	}
	return nil, netval
}

func (rd *RIPEd) GetMaintainer() (err error, mntby string) {
	mnterr, mntval := rd.getValByKey("mnt-by", rd.RipeData.IRR_Records)
	if mnterr != nil {
		return fmt.Errorf("%s", "mnt not found"), ""
	}
	return nil, mntval
}

func (rd *RIPEd) GetDescription() (err error, mntby string) {
	mnterr, mntval := rd.getValByKey("descr", rd.RipeData.IRR_Records)
	if mnterr != nil {
		return fmt.Errorf("%s", "descr not found"), ""
	}
	return nil, mntval
}

func (rd *RIPEd) GetOriginAs() (err error, OriginAs uint) {
	origerr, origval := rd.getValByKey("origin", rd.RipeData.IRR_Records)
	if origerr != nil {
		return fmt.Errorf("%s", "origin not found"), 0
	}
	asnumm, asnumerr := strconv.Atoi(origval)
	if asnumerr != nil {
		return asnumerr, 0
	}
	return nil, uint(asnumm)
}

func (rd *RIPEd) GetData(Ipaddr string) (err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	var resp *http.Response

	urlstring := fmt.Sprintf("%s%s", RIPEWHOIS_URL, Ipaddr)

	resp, err = client.Get(urlstring)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//var Ddata RipeDataExt
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return err
		}
		var OwerallJson map[string]json.RawMessage
		UmErr := json.Unmarshal(bodyBytes, &OwerallJson)
		fmt.Println(UmErr)
		for key, val := range OwerallJson {
			if key == "data" {
				var DataJm ripeData
				DataUmErr := json.Unmarshal([]byte(val), &DataJm)
				if DataUmErr != nil {
					return DataUmErr
				}
				rd.RipeData = DataJm
			}
		}
	}
	return nil
}
