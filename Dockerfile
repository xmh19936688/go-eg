FROM alpine

#COPY repeater /
COPY cacher /
COPY static /static/

WORKDIR /

EXPOSE 9000

#ENTRYPOINT ["/repeater"]
ENTRYPOINT ["/cacher"]
