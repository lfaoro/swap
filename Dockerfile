FROM alpine:latest
WORKDIR /srv
COPY ./bin/swap /srv/swap
RUN apk add --no-cache ca-certificates
ENTRYPOINT ["/srv/swap"]
CMD []