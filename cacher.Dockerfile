FROM alpine

COPY cacher /
COPY static /static/
COPY config.template.yaml /config.yaml

WORKDIR /

EXPOSE 9000

ENTRYPOINT ["/cacher"]
