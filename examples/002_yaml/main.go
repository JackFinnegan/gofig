package main

import (
	"fmt"

	"go.krak3n.codes/gofig"
	"go.krak3n.codes/gofig/parsers/yaml"
)

// Config is a struct to unpack configuration into.
type Config struct {
	Foo struct {
		Bar struct {
			Baz string `gofig:"baz"`
		} `gofig:"bar"`
	} `gofig:"foo"`
	Fizz struct {
		Buzz map[string]string `gofig:"buzz"`
	} `gofig:"fizz"`
	A struct {
		B map[string][]int `gofig:"b"`
	} `gofig:"a"`
	C struct {
		D map[string]map[string][]int `gofig:"d"`
	} `gofig:"c"`
}

const blob string = `
foo:
  bar:
    baz: bar
fizz:
  buzz:
    hello: world
    bill: ben`

func main() {
	var cfg Config

	// Initialise gofig with the struct config values will be placed into
	gfg, err := gofig.New(&cfg)
	gofig.Must(err)

	// Create a parser
	parser := yaml.New()

	// Parse in order
	gofig.Must(gfg.Parse(
		gofig.FromFile(parser, "./config.yaml"),
		gofig.FromString(parser, blob)))

	fmt.Println(fmt.Sprintf("%+v", cfg))
}
