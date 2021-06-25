FROM golang:1.16 AS go-env
COPY . /go/src/go-env/
WORKDIR /go/src/go-env/
RUN ls && CGO_ENABLED=0 go build ./cmd/ui

FROM alpine
COPY --from=go-env /go/src/go-env/ui /
WORKDIR /
EXPOSE 9000
ENTRYPOINT ["/ui"]