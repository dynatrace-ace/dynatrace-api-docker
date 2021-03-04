FROM golang:alpine3.13 AS builder
WORKDIR /go/src/github.com/dynatrace-ace/dynatrace-api-docker/dynatrace-push-event/
COPY . .
RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dynatrace-push-event .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/dynatrace-ace/dynatrace-api-docker/dynatrace-push-event/dynatrace-push-event .
CMD ["./dynatrace-push-event"]