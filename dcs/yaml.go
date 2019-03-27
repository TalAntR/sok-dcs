package dcs

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

func yamlConfig() DcsConfig {
	var svc DcsConfig
	yml, err := ioutil.ReadFile("dcs.d/dcsd.yml")
	if err != nil {
		log.Fatalf("Configuration is not found:  #%v ", err)
	}
	err = yaml.Unmarshal(yml, &svc)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return svc
}
