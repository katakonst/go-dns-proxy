FROM golang:alpine as pacpac
LABEL mantainer="Pirvu Catalin <pirvu.catalin94@gmail.com>"
RUN apk add --update --no-cache git
RUN go get github.com/katakonst/go-dns-proxy
RUN go build -o go-dns-proxy

FROM alpine:latest
RUN apk --no-cache add ca-certificates
EXPOSE 53/udp
WORKDIR /root/
COPY --from=pacpac  /go/src/katakonst/go-dns-proxy .
ENTRYPOINT ["./go-dns-proxy"]