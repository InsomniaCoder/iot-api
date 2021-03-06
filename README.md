# Instruction

This project is developped using Golang.

It basically is a CRUD API for iot devices providing persistence to the database, and publish/subscibe with Kafka.

Another purpose of this project is to act as an initilized project structure for future Go App development.

I will try to incorporate technologies, concepts, good practices in including

- Externalized configuration using Viper
- API development using Gin
- Hexagonal architecture / Interface driven inspired by https://github.com/bxcodec/go-clean-arch
- [Todo] Unit testing and integration testing
- ORM using Gorm
- Connection with Kafka via segmentio/kafka-go
- Dockerfile / Docker-compose inspired by https://github.com/bxcodec/go-clean-arch and KubeBuilder


## Local

- run `make local` to run API internally
- run `make run` to rebuilt the project and start the service along with database using docker-compose
- run `make docker` to build a Docker image of this service

## API usage

### Endpoints

#### POST /api/v1/sensors

request body:

```
{
 "soil_moisture"   : 10.2,
}
```

##### using curl to try

```
curl --header "Content-Type: application/json" --request POST --data '{"soil_moisture":10.2}' http://localhost:8082/api/v1/sensors
```

#### GET /api/v1/sensors
##### using curl to try
```
curl --header "Content-Type: application/json" --request GET http://localhost:8082/api/v1/sensors
```


#### POST /api/v1/images/search

request body:

```
{
 "startTime"   : "2021-07-07T00:00:00Z",
 "endTime"     : "2021-07-08T00:00:00Z"
}
```

##### using curl to try
```
curl --header "Content-Type: application/json" --request POST --data '{"startTime":"2021-07-07T00:00:00Z","endTime":"2021-07-08T00:00:00Z"}'  http://localhost:8082/api/v1/images/search
```

#### POST /api/v1/images/capture

##### using curl to try
```
curl --header "Content-Type: application/json" --request POST http://localhost:8082/api/v1/images/capture
```