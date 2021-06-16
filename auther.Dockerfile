FROM alpine

COPY auther /

WORKDIR /

EXPOSE 9000

ENTRYPOINT ["/auther"]
