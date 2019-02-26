package apiCall

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Get() CustomerPoint {
	resp, err := http.Get("https://services.ebz-chn-dev.mkaws.com/PointRedeemServices/v1/customers/1100012821/point?RedeemType=CustomerPoint")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var point CustomerPoint
	err = json.Unmarshal(body, &point)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Get request result: %s\n", string(body))
	return point
}
