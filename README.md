# 💡GoFig

[![Go Version][goversion-image]][goversion-url]
[![Documentation][doc-image]][doc-url]
[![Workflow Status][workflow-image]][workflow-url]
[![Go Report Card][report-image]][report-url]

GoFig is a configuration loading library for Go. It aims to provide a simple, flexible and
decoupled API for all your configuration loading needs.

* **Status**: PoC (Proof of Concept)

## Example

``` go
package main

import (
	"log"

	"go.krak3n.codes/gofig"
	"go.krak3n.codes/gofig/parsers/env"
	"go.krak3n.codes/gofig/parsers/yaml"
	"go.krak3n.codes/gofig/notifiers/fsnotify"
)

type Config  struct {
	Foo struct {
		Bar string `gofig:"bar"`
	} `gofig:"foo"`
	Fizz struct {
		Buzz string `gofig:"buzz"`
	} `gofig:"fizz"`
}

func main() {
	var cfg Config

	// Initialise gofig with the destination struct
	gfg, err := gofig.New(&cfg)
	gofig.Must(err)

	// Setsup a yaml parser with file notification support
	yml := gofig.FromFileWithNotify(yaml.New(), fsnotify.New("./config.yaml"))

	// Setup a notification channel to send notification of configuration updates
	notifyCh := make(chan error, 1)
	gfg.Notify(notifyCh, yml)

	// Parse the yaml file and then environment variables
	gofig.Must(gfg.Parse(yml, env.New(env.HasAndTrimPrefix("GOFIG"))))

	// Watch for configuration changes to reload your application
	for {
		log.Printf("configuration: %+v\n", cfg)

		if err := <- notifyCh; err != nil {
			log.Fatal(err) // Failed to read or parse the configuration change
		}
	}
}
```

## Parsers

GoFig implements it's parsers as sub modules. Currently it supports:

* [Environment Variables][env-url]
* [JSON][json-url]
* [TOML][toml-url]
* [YAML][yaml-url]

# Roadmap

* [ ] Test Suite / Code Coverage reporting
* [ ] Helpful errors
* [ ] Support pointer values
* [ ] Default Values via a struct tag, e.g: `gofig:"foo,default=bar"`
* [ ] Support `omitempty` for pointer values which should not be initialised to their zero value.
* [ ] Support notification of config changes via `Notifier` interface
* [x] Implement File notifier on changes to files via `fsnotify`
* [ ] Add support for:
  * [ ] ETCD Parser / Notifier
  * [ ] Consul Parser / Notifier

[workflow-image]: https://img.shields.io/github/workflow/status/krak3n/gofig/GoFig?style=flat&logo=github&logoColor=white&label=Workflow
[workflow-url]: https://github.com/krak3n/gofig/actions?query=workflow%3AGoFig
[goversion-image]: https://img.shields.io/badge/Go-1.13+-00ADD8.svg?style=flat&logo=go&logoColor=white
[goversion-url]: https://golang.org/
[doc-image]: https://img.shields.io/badge/Documentation-pkg.go.dev-00ADD8.svg?style=flat&logo=go&logoColor=white
[doc-url]: https://pkg.go.dev/go.krak3n.codes/gofig
[report-image]: https://goreportcard.com/badge/github.com/krak3n/gofig?style=flat-square
[report-url]: https://goreportcard.com/report/github.com/krak3n/gofig
[env-url]: ./parsers/env
[json-url]: ./parsers/json
[toml-url]: ./parsers/toml
[yaml-url]: ./parsers/yaml
