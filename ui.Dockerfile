FROM alpine

COPY ui /
COPY static /static/
COPY config.template.yaml /config.yaml

WORKDIR /

EXPOSE 9000

ENTRYPOINT ["/ui"]
