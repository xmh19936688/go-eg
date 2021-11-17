package setting

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Version from build flag, do NOT modify.
var Version = ""

type config struct {
	App struct {
		Port string `json:"port" yaml:"port"`
	} `json:"app" yaml:"app"`
	Auth struct {
		Secret string `json:"secret" yaml:"secret"`
		Url    string `json:"url" yaml:"url"`
	} `json:"auth" yaml:"auth"`
	Cache struct {
		Url string `json:"url" yaml:"url"`
	} `json:"cache" yaml:"cache"`
}

var Config config

func init() {
	Config.App.Port = "9000"
	Config.Auth.Secret = "admin"
	Config.Auth.Url = "http://localhost:9001/auth"
	Config.Cache.Url = "http://localhost:9002/"

	f, err := os.Open("config.yaml")
	if err != nil {
		log.Println("failed to open config.yaml:", err.Error())
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("failed to read from config.yaml:", err.Error())
		return
	}

	log.Println("config.yaml found:")
	log.Println(string(bs))

	err = yaml.Unmarshal(bs, &Config)
	if err != nil {
		log.Println("failed to read config:", err.Error())
	}
}
