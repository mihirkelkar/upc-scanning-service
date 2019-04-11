package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

//Config : A configuration object that contains the redis config and
// the third party api config.
type Config struct {
	redisConfig *redisConfig
	apiConfig   *thirdPartyAPIConfig
}

//redisConfig : Store the configuration for Redis.
//note that this is an unexported configuration structure with unexported fields.
type redisConfig struct {
	address  string `json:"address"`
	password string `json:"password"`
	database int    `json:"database"`
}

//thirdPartyAPIConfig : store the configuration for the api.
//note that this is an unexported config structure with unexported fields.
type thirdPartyAPIConfig struct {
	apikey string
}

//NewConfig : Returns a new config object
func NewConfig() *Config {
	return &Config{}
}

//ReadRedisJSON : reads a json file of configuration and stores it in a
//config object.
func (c *Config) ReadRedisJSON(configname string) error {
	data, err := ioutil.ReadFile(configname)
	if err != nil {
		return err
	}
	var v interface{}
	json.Unmarshal(data, &v)
	vi, _ := v.(map[string]interface{})
	c.redisConfig.address = vi["address"].(string)
	c.redisConfig.password = vi["password"].(string)
	c.redisConfig.database = int(vi["database"].(float64))
	return nil
}

//ReturnConfig : Returns a map that can be used in other places
func (c *Config) ReturnConfig() map[string]string {
	configmap := make(map[string]string)
	configmap["address"] = c.redisConfig.address
	configmap["password"] = c.redisConfig.address
	configmap["database"] = strconv.Itoa(c.redisConfig.database)
	return configmap
}

func (c *Config) SetApiKey(apikey string) {
	c.apiConfig.apikey = apikey
}

func (c *Config) GetApiKey() string {
	return c.apiConfig.apikey
}
