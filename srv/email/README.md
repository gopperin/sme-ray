# run

## build

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w'

## cli 

MICRO_REGISTRY=etcdv3 MICRO_REGISTRY_ADDRESS=192.168.1.104:2579 MICRO_BROKER=nsq MICRO_BROKER_ADDRESS=192.168.1.104:4150 email

MICRO_REGISTRY=etcdv3 MICRO_REGISTRY_ADDRESS=192.168.1.104:2579 MICRO_BROKER=nats MICRO_BROKER_ADDRESS=192.168.1.104:4222 email

## docker

sudo docker build -t sme-ray_email:latest .

sudo docker-compose up

sudo docker-compose down

sudo docker rmi sme-ray_email