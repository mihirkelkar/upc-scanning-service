package models

import (
	"errors"

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
}

//NewServices: Generates a new service that can be used by other
//models.
func NewServices() (*Services, error) {
	//create a new redis client.
	redisClient, err := newRedisClient()
	if err != nil {
		return nil, err
	}

	prdService, prdSrvErr := NewProductService(redisClient)
	if prdSrvErr != nil {
		return nil, prdSrvErr
	}

	return &Services{ProductService: prdService}, nil

}

func newRedisClient() (*redis.Client, error) {
	//TODO : Need to take this out into a config package
	//set the config package from a command line in the main.go file
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//test the client. If something is amiss, return an error.
	_, err := client.Ping().Result()
	if err != nil {
		return nil, errors.New("Error: Could not connect to Redis")
	}
	return client, nil
}
