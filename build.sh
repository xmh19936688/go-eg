docker run -d --name  go-env-16 golang:1.16
docker start go-env-16

docker exec go-env-16 rm -rf /go/src/go-eg
docker cp . go-env-16:/go/src/go-eg/

docker exec go-env-16 sh -c "cd /go/src/go-eg && CGO_ENABLED=0 go build /go/src/go-eg/cmd/auther"
docker exec go-env-16 sh -c "cd /go/src/go-eg && CGO_ENABLED=0 go build /go/src/go-eg/cmd/cacher"
docker exec go-env-16 sh -c "cd /go/src/go-eg && CGO_ENABLED=0 go build /go/src/go-eg/cmd/ui"

docker cp go-env-16:/go/src/go-eg/auther .
docker cp go-env-16:/go/src/go-eg/cacher .
docker cp go-env-16:/go/src/go-eg/ui .

docker build . -f auther.Dockerfile -t xmh-auther:1
docker build . -f cacher.Dockerfile -t xmh-cacher:1
docker build . -f ui.Dockerfile -t xmh-ui:1
