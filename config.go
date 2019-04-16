package main

import (
	"encoding/json"
	"fmt"
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
	apiurl string
}

//NewConfig : Returns a new config object
func NewConfig() *Config {
	return &Config{redisConfig: &redisConfig{}, apiConfig: &thirdPartyAPIConfig{}}
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
	c.apiConfig.apikey = vi["apikey"].(string)
	c.apiConfig.apiurl = vi["apiurl"].(string)
	return nil
}

//ReturnConfig : Returns a map that can be used in other places
func (c *Config) ReturnConfig() map[string]string {
	configmap := make(map[string]string)
	configmap["address"] = c.redisConfig.address
	configmap["password"] = c.redisConfig.password
	configmap["database"] = strconv.Itoa(c.redisConfig.database)
	configmap["apiurl"] = c.GetCompleteAPIURL()
	return configmap
}

//GetAPIKey : Returns the API key
func (c *Config) GetAPIKey() string {
	return c.apiConfig.apikey
}

//GetCompleteAPIURL : Returns the complete formatted API url
func (c *Config) GetCompleteAPIURL() string {
	return fmt.Sprintf(c.apiConfig.apiurl, c.GetAPIKey())
}
