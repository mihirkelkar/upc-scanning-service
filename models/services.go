package models

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

/*
This was created to create all of our model's services at the moment.
This service will instantiate a redis at the moment and then return
the redis service's client back to the model that wants to use it.
*/

type Services struct {
	/*
	  This is the encapusulation that will eventually store
	  services related to all the models used here.
	*/
	ProductService ProductService
	db             *redis.Client
}

//NewServices: Generates a new service that can be used by other
//models.Accepts a map of configurations for redis.
func NewServices(config map[string]string) (*Services, error) {
	//create a new redis client.
	redisClient, err := newRedisClient(config)
	if err != nil {
		return nil, err
	}
	brcdService, err := NewBarcodeLookup(config["apiurl"])
	if err != nil {
		return nil, err
	}
	prdService, prdSrvErr := NewProductService(redisClient, brcdService)
	if prdSrvErr != nil {
		return nil, prdSrvErr
	}

	return &Services{ProductService: prdService, db: redisClient}, nil

}

func newRedisClient(config map[string]string) (*redis.Client, error) {
	//TODO : Need to take this out into a config package
	//set the config package from a command line in the main.go file
	db, _ := strconv.Atoi(config["database"])
	client := redis.NewClient(&redis.Options{
		Addr:     config["address"],
		Password: config["password"], // no password set
		DB:       db,                 // use default DB
	})

	//test the client. If something is amiss, return an error.
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error: Could not connect to Redis")
	}
	return client, nil
}

//Close : close connection to the redis client.
func (s *Services) Close() error {
	return s.db.Close()
}
