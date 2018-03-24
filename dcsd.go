package main

import (
	"encoding/json"
	"fmt"
	"github.com/TalAntR/sok-dcs/dcs"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"net/http"
)

func yamlConfig() dcs.DcsConfig {
	var svc dcs.DcsConfig
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

// -H "Accept: application/json"
func dscHandler(w http.ResponseWriter, r *http.Request) {
	response := "{}"
	if r.Method == "GET" {
		response = "Hello! I'm dscd!"
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, response)
}

func main() {
	svc := yamlConfig()
	jsonSvc, err := json.Marshal(svc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonSvc))
	//http.HandleFunc("/", dscHandler)
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//    log.Fatal("ListenAndServe: ", err)
	//}
}
