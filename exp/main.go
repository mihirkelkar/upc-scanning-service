package main

import "fmt"
import "github.com/go-redis/redis"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	client.Del("bss:123456")

	field_map := make(map[string]interface{})
	field_map["productname"] = "Thrive Market Cassava Chips"
	field_map["upc"] = "11234456567"
	field_map["searchterm"] = "11234456567"
	field_map["catalog"] = true
	client.HMSet("bss:123456", field_map).Result()
	if err != nil {
		fmt.Println(err)
	}

	result, err := client.HMGet("bss:123456", "productname", "upc", "searchterm", "catalog").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
