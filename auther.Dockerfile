FROM alpine

COPY auther /
COPY config.template.yaml /config.yaml

WORKDIR /

EXPOSE 9000

ENTRYPOINT ["/auther"]
