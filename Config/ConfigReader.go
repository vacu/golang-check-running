package Config

import (
    "encoding/json"
    "io/ioutil"
)

type Config struct {
    Daemons []string
    Interval int
}

func LoadConfig()(Config) {
    file, err := ioutil.ReadFile("config.json")
    if err != nil {
        panic(err.Error())
    }

    var config Config
    json.Unmarshal(file, &config)

    return config
}
