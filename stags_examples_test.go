package stags

import (
	"encoding/json"
	"io/ioutil"
)

func ExampleSTags() {
	type Configuration struct {
		Root   string `env:"CFG_ROOT" default:"/usr/share/doc"`
		Server struct {
			Address string `env:"CFG_ADDR" default:"localhost:8080"`
		}
	}

	var cfg Configuration

	s := New(cfg)

	content, err := ioutil.ReadFile("file.conf")
	if err != nil {
		panic(err)
	}

 	err = json.Unmarshal(content, &cfg)
	if err != nil {
		panic(err)
	} 

	if cfg.Root == "" {
		cfg.Root = s.Get("default", "Root")
	}

	if cfg.Server.Address == "" {
		cfg.Server.Address = s.Get("default", "Server", "Address")
	}

}
