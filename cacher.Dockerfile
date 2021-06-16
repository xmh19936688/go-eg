FROM alpine

COPY cacher /
COPY static /static/

WORKDIR /

EXPOSE 9000

ENTRYPOINT ["/cacher"]
