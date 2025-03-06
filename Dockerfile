FROM busybox
WORKDIR /srv
COPY ./bin/swap_api /srv/swap_api
CMD ["/srv/swap_api"]
