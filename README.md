# sme-ray
sme-ray

## docker run

### Protobufs

make proto

### Run

make build

sudo make run

sudo docker-compose up

MICRO_REGISTRY=etcdv3 MICRO_REGISTRY_ADDRESS=10.19.20.72:2379 MICRO_API_NAMESPACE=snc.gc.api micro api --handler=http

sudo docker-compose down

### HTTP

service.http