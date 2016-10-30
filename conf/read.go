package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Name string `yaml:"Name"`
	Env  Env    `yaml:"Env"`
}

type Env struct {
	Extension string `yaml:"Extension"`
	RootPath  string `yaml:"RootPath"`
}

func main() {
	f, _ := filepath.Abs("./config.yml")
	yf, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
	}

	var config Config
	err = yaml.Unmarshal(yf, &config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Name: ", config.Name)
	fmt.Println("Env: ", config.Env)
}
