FROM frolvlad/alpine-glibc

COPY log log

COPY config config

RUN apk add --no-cache bash

ADD main /

EXPOSE 8081

CMD ["/main"]