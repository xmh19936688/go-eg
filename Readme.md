# step

## image

1. docker pull golang:1.16

## container

1. docker run --name go-env-16 -itd golang:1.16 bash

## clone

1. git clone http://172.16.99.4:10080/xumenghao/go-eg.git

## go build

1. docker cp ./go-eg go-env-16:/go/src/
2. docker exec -it go-env-16 bash
3. cd /go/src/go-eg
4. CGO_ENABLED=0 go build ./cmd/cacher/
5. exit
6. docker cp go-env-16:/go/src/go-eg/cacher

## docker build

1. cd go-eg
2. docker build . -t xx:xx

## test

```sh
docker run -d --name xx -p 9000:9000 xx:xx

curl --location --request GET 'http://192.168.103.61:9000/' \
--header 'Content-Type: text/plain' \
--data 'your-data'
```