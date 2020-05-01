FROM alpine:latest

RUN apk --no-cache add \
    bash

RUN apk add --update coreutils && rm -rf /var/cache/apk/*

RUN apk add curl

COPY push-event.sh ./

ENTRYPOINT ["./push-event.sh"]