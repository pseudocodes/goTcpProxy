// Author: zheng-ji.info

package main

import (
	//"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ProxyConfig Type
type ProxyConfig struct {
	Bind         string    `yaml:"bind"`
	WaitQueueLen int       `yaml:"wait_queue_len"`
	MaxConn      int       `yaml:"max_conn"`
	Timeout      int       `yaml:"timeout"`
	FailOver     int       `yaml:"failover"`
	Backend      []string  `yaml:"backend"`
	Log          LogConfig `yaml:"log"`
	Stats        string    `yaml:"stats"`
}

// LogConfig Type
type LogConfig struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

func parseConfigFile(filepath string) error {
	if config, err := ioutil.ReadFile(filepath); err == nil {
		if err = yaml.Unmarshal(config, &pConfig); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}
