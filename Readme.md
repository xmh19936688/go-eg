# step

## image

1. docker pull golang:1.16

## container

1. docker run --name go-env-16 -itd golang:1.16 bash

## clone

1. git clone http://172.16.99.4:10080/xumenghao/go-eg.git

## build

1. chmod +x ./build.sh
2. ./build.sh

## test

```shell
IP=192.168.103.61
cp ./config.template.yaml ./config.yaml
sed -i "s/localhost/${IP}/g" ./config.yaml

docker run -d --name xmh-ui -p 9000:9000 -v $(pwd)/config.yaml:/config.yaml xmh-ui:1
docker run -d --name xmh-auther -p 9001:9000 -v $(pwd)/config.yaml:/config.yaml xmh-auther:1
docker run -d --name xmh-cacher -p 9002:9000 -v $(pwd)/config.yaml:/config.yaml xmh-cacher:1

open "http://${IP}:9000/" in browser
```

## clean

```shell
docker stop xmh-ui xmh-auther xmh-cacher
docker rm xmh-ui xmh-auther xmh-cacher
docker rmi xmh-ui:1 xmh-auther:1 xmh-cacher:1
rm -f ./config.yaml ./ui ./auther ./cacher
```
