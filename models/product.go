package models

import (
	"strconv"

	"github.com/go-redis/redis"
)

type Product struct {
	ProductName string `json:"productname"`
	Upc         string `json:"upc"`
	SearchTerm  string `json:"searchterm"`
	Catalog     int    `json:"catalog"`
}

func (p *Product) populateStruct(result []interface{}) {
	//if the upc is not filled, then the result is assumed empty.
	if result[1] == nil {
		return
	}
	p.ProductName = result[0].(string)
	p.Upc = result[1].(string)
	p.SearchTerm = result[2].(string)
	p.Catalog, _ = strconv.Atoi(result[3].(string))
}

//isEmpty : If the UPC field of a product is empty, the product is
//assumed empty
func (p *Product) isEmpty() bool {
	if p.Upc == "" {
		return true
	}
	return false
}

//ProductRedis : The Product Redis Interface is a
// collection of methods that insert, update or retrieve
// a product from the redis server.
// we will have a new struct fulfil this interface that
// can actually access Redis.

//When writing tests, do not test the actually connections with
type ProductDB interface {
	ByUpc(string) (*Product, error)
}

//This will be a client that eventually connects to Redis.
//We can also write a higher layer that still implements
//the ProductRedis interface along with a bunch of other
//validation functions
type productDB struct {
	client *redis.Client
}

func (p *productDB) buildKey(upc string) string {
	return "bss:" + upc
}

func (p *productDB) ByUpc(upc string) (*Product, error) {
	upc = p.buildKey(upc)
	product, err := p.client.HMGet(upc, "productname", "upc", "searchterm", "catalog").Result()
	if err != nil {
		return nil, err
	}

	prd := Product{}
	prd.populateStruct(product)
	if prd.isEmpty() {
		//no response was returned from Redis.
		return nil, nil
	}
	return &prd, nil
}

//TODO : Write a productValidator struct that implements the ProductRedis
//interface. the productvalidator class will validate the
//set functions when adding things to redis.
//The productValidator struct will be exposed as the struct that implements
//ProductRedis
//so that the entore service interface has access to

//ProductService : The ProductService interface can be exported to all other files.
//as a single encapusulation

type ProductService interface {
	ProductDB
}

type productService struct {
	ProductDB
}

//Creates a service for interacting with products in Redis.
func NewProductService(redisClient *redis.Client) (ProductService, error) {
	prdRedis := &productDB{client: redisClient}
	prdService := &productService{ProductDB: prdRedis}
	return prdService, nil
}
