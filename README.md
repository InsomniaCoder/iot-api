# Instruction

This project is developped using Golang.

It basically is a CRUD API for iot devices.

Another purpose of this project is to act as an initilized project structure for future Go API development.

I will try to incorporate technologies, concepts, good practices in including

- Externalized configuration using Viper
- API development using Gin
- Hexagonal architecture / Interface driven inspired by https://github.com/bxcodec/go-clean-arch
- [Todo] Unit testing and integration testing
- ORM using Gorm
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