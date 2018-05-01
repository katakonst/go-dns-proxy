# DNS Proxy
A simple DNS proxy written in go on top of [github.com/miekg/dns](https://github.com/miekg/dns)

## How to use it


# Docker
```
```
# Download executables
```
[Download](https://github.com/katakonst/go-dns-proxy/releases)
```

# GO get
```shell
$ go get github.com/katakonst/go-dns-proxy
```


## Arguments & Options

```
	-file		    config filename
	-log-level		log level
	-expiration		cache expiration time in seconds
	-use-outbound	use outbound address as host for server
```

## Config file format

```
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