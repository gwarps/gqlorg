package config

import (
	"io/ioutil"
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

var AppConfig Config

var DgraphClient *dgo.Dgraph

type Config struct {
	Dgraph struct {
		Host string `yaml:"host"`
	} `yaml:"dgraph"`
}

func ReadConfig() error {
	configYaml, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Println("Config read error")
		return err
	}

	var config Config

	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		log.Println("Config unmarshal error")
	}

	AppConfig = config
	return err
}

func InitDgraphConnection() error {
	log.Println("Connecting to Dgraph server: " + AppConfig.Dgraph.Host)
	d, err := grpc.Dial(AppConfig.Dgraph.Host, grpc.WithInsecure())
	if err != nil {
		log.Println("Error in creating Dgraph connection")
		return err
	}

	DgraphClient = dgo.NewDgraphClient(api.NewDgraphClient(d))

	log.Println("Connected to Dgraph")

	return nil
}
