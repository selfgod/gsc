package Base

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ConfigBase struct {
	Env string
	Dev Config
	Pro Config
}

type Config struct {
	Db    Db
	Redis Redis
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
		return &ret.Dev
	}
	return &ret.Pro
}
