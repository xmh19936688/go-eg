FROM alpine

COPY repeater /

WORKDIR /

EXPOSE 9000

ENTRYPOINT ["/repeater"]
