package conf

import (
	"encoding/json"
	"io/ioutil"
	"time"

	conf "github.com/mikeqiao/ant/config"
	"github.com/mikeqiao/ant/log"
)

func InitConfig() {
	data, err := ioutil.ReadFile("conf/config.json")

	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &conf.Config)
	if err != nil {
		log.Fatal("%v", err)
	}
	conf.Config.HTTPTimeout = 3 * time.Second
	conf.Config.ConnectInterval = 3 * time.Second
	conf.Config.LittleEndian = false

}
