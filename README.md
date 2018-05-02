# DNS Proxy
A simple DNS proxy written in go based on [github.com/miekg/dns](https://github.com/miekg/dns)

## How to use it


## Docker

```shell
$ docker run -p 153:53/udp katakonst/go-dns-proxy:latest -use-outbound -json-config='{
    "defaultDns": "8.8.8.8:53",
    "servers": {
        "google.com" : "8.8.8.8:53"
    },
    "domains": {
        "test.com." : "8.8.8.8"
    }
}'
```

## Download executables

[Download](https://github.com/katakonst/go-dns-proxy/releases)

## Go get

```shell
$ go get github.com/katakonst/go-dns-proxy
$ go-dns-proxy -use-outbound -json-config='{
    "defaultDns": "8.8.8.8:53",
    "servers": {
        "google.com" : "8.8.8.8:53"
    },
    "domains": {
        "test.com." : "8.8.8.8"
    }
}'
```

## Arguments

```
	-file		    config filename
	-log-level		log level(info or error)
	-expiration		cache expiration time in seconds
	-use-outbound	use outbound address as host for server
    -config-json    configs as json
```

## Config file format

```json
{
    "host": "192.168.1.4:53",
    "defaultDns": "8.8.8.8:53",
    "servers": {
        "google.com" : "8.8.8.8:53"
    },
    "domains": {
        ".*.com." : "8.8.8.8"
    }
}
```