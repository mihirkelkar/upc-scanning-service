package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type RedisConfig struct {
	Address  string `json string`
	Password string `json string`
	Database int    `json int`
}

//NewConfig : Returns a new config object
func NewConfig() *RedisConfig {
	return &RedisConfig{}
}

//ReadConfigJson : reads a json file of configuration and stores it in a
//config object.
func (rr *RedisConfig) ReadConfigJson(configname string) error {
	data, err := ioutil.ReadFile(configname)
	if err != nil {
		return err
	}
	var v interface{}
	json.Unmarshal(data, &v)
	vi, _ := v.(map[string]interface{})
	rr.Address = vi["address"].(string)
	rr.Password = vi["password"].(string)
	rr.Database = int(vi["database"].(float64))
	return nil
}

//ReturnConfig : Returns a map that can be used in other places
func (rr *RedisConfig) ReturnConfig() map[string]string {
	configmap := make(map[string]string)
	configmap["address"] = rr.Address
	configmap["password"] = rr.Password
	configmap["database"] = strconv.Itoa(rr.Database)
	return configmap
}
