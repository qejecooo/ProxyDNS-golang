package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	BlackList map[string]struct {
		IPv4 string `json:"ipv4"`
		IPv6 string `json:"ipv6"`
	} `json:"black_list"`
	UpstreamServer struct {
		IP   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"upstream_server"`
}

func ParseJson() (Config, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Printf("Error reading file: %v\n", err)
		return Config{}, err
	}

	// Parse the JSON data into a Config struct
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return Config{}, err
	}
	return config, nil
}
