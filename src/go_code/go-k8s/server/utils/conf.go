package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var ConfMap *map[string]interface{} = confParse()

func confParse() *map[string]interface{} {
	yamlFile, err := ioutil.ReadFile("conf/user.yaml")
	if err != nil {
		log.Fatalf("parse config file conf/user.yaml failed, Error: %v", err)
	}

	result := make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, result)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("parse configfile conf/user.yaml success")
	return &result
	// return result, err
	// log.Println("conf", result)
}

// func main() {
// 	ConfParse()
// }
