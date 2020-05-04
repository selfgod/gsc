package Base

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var GConfig *Config

type ConfigBase struct {
	Env string
	Dev Config
	Pro Config
}

type Config struct {
	Db     Db
	Redis  Redis
	Secret string
}

type Db struct {
	Host     string
	Name     string
	Username string
	Password string
}

type Redis struct {
	Host     string
	Username string
	Password string
}

func LoadConfig() *Config {
	c, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("config:%v\n", string(c))
	ret := new(ConfigBase)
	err = json.Unmarshal(c, ret)
	if err != nil {
		panic(err.Error())
	}
	if ret.Env == "dev" {
		GConfig = &ret.Dev
	} else {
		GConfig = &ret.Pro
	}
	return GConfig
}
