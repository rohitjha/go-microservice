# go-microservice

## Run all unit and component tests

```bash
go test ./...
```

## Build the container

```bash
docker build -t "go-microservice" .
docker images "go-microservice"
```

## Run the container

```bash
docker run -it -p 3001:3001 go-microservice
```

## Testing using curl

```bash
curl -kv https://localhost:3001/products
curl -kv https://localhost:3001/products/Latte
curl -kv -X PUT https://localhost:3001/products -d '{"name":"Pie","description":"pi","price":3.1415,"sku":"123456"}' -H 'Content-Type: application/json'
curl -kv https://localhost:3001/products
curl -kv -X POST https://localhost:3001/products/Pie -d '{"description":"pie"}' -H 'Content-Type: application/json'
curl -kv https://localhost:3001/products/Pie
curl -kv -X DELETE https://localhost:3001/products/Pie
curl -kv https://localhost:3001/products/Pie
curl -kv https://localhost:3001/products
```
