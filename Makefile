BINARY=iot-api
test: 
	go test -v -cover -covermode=atomic ./...

local:
	go run app/main.go

iot-api:
	go build -o ${BINARY} app/*.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t ${BINARY} .

run:
	docker-compose up --build -d

stop:
	docker-compose down