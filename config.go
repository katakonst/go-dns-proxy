package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

type Config struct {
	DNSConfigs      map[string]interface{}
	CacheExpiration int64
	UseOutbound     bool
	LogLevel        string
}

func InitConfig() (Config, error) {
	fileName := flag.String("file", "config.json", "config filename")
	logLevel := flag.String("log-level", "info", "log level")
	expiration := flag.Int64("expiration", -1, "expiration time in seconds")
	useOutbound := flag.Bool("use-outbound", false, "use outbound address")
	cliConfigs := flag.String("json-config", "", "config in json format")
	flag.Parse()

	dnsConfigs := make(map[string]interface{})
	if *cliConfigs == "" {
		var err error
		dnsConfigs, err = parseFile(*fileName)
		if err != nil {
			return Config{}, err
		}
	} else {
		if err := json.Unmarshal([]byte(*cliConfigs), &dnsConfigs); err != nil {
			return Config{}, err
		}
	}

	return Config{
		DNSConfigs:      dnsConfigs,
		CacheExpiration: *expiration * 1000000000,
		UseOutbound:     *useOutbound,
		LogLevel:        *logLevel,
	}, nil
}

func parseFile(filePath string) (map[string]interface{}, error) {
	fileContents := make(map[string]interface{})
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &fileContents); err != nil {
		return nil, err
	}

	return fileContents, nil
}
