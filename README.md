# upc-scanning-service
This is a barcode scanning service backend. The service maintains a cache of search terms for UPCs in Redis. For the products
that are not in the Private Label catalog, the service goes to barcodelookupapi.com , fetches the barcode info and returns a 
response
