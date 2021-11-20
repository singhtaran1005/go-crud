package main

import (
	"bytes"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Number int `yaml:"number"`
}

// Earlier version of code with log.Fatalf.
func ParseConfig(c []byte) Config {
	cfg := Config{}
	dec := yaml.NewDecoder(bytes.NewReader(c))
	dec.KnownFields(true)
	log.Fatalf("Decode failed %v", dec.Decode(&cfg))
	return cfg
}

// New version of code with log.Errorf.
func ParseConfigNew(c []byte) Config {
	cfg := Config{}
	dec := yaml.NewDecoder(bytes.NewReader(c))
	dec.KnownFields(true)
	if err := dec.Decode(&cfg); err != nil {
		log.Errorf("Decode failed %v", err)
		return Config{}
	}
	return cfg
}

func main() {
	config := []byte("number: 'abc'")
	// cfg := ParseConfig(config)
	cfg := ParseConfigNew(config)
	fmt.Println("This will print if not fatal", cfg.Number)
}
