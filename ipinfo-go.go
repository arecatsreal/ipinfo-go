package ipinfogo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Ipinfo struct {
	Ip       string   `json:"ip"`
	City     string   `json:"city"`
	Country  string   `json:"country"`
	Loc      string   `json:"loc"`
	Org      string   `json:"org"`
	Postal   string   `json:"postal"`
	Timezone string   `json:"timezone"`
	Readme   string   `json:"readme"`
}

func IpinfoRun() Ipinfo {
	web, err := http.Get("https://ipinfo.io")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(web.Body)
	var getInfo string
	for i := 0; scanner.Scan() && i < 11; i++ {
		// fmt.Println(scanner.Text())
		getInfo = fmt.Sprintln(getInfo + scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	web.Body.Close()
	getInfo = strings.TrimSpace(getInfo)

	out := Ipinfo{}
	json.Unmarshal([]byte(getInfo), &out)
	
	return out
}
