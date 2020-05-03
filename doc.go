// Package gofig provides a library for loading configuration into a struct type. It also provides
// notification functionality (for parsers supporting it) for when configuration changes whilst the
// application is running, allowing you to hot reload your application when configuration changes.
//
// At it's core Gofig takes no 3rd party dependencies, parsers are implemented as their own sub
// modules, which may take 3rd party dependencies so you only get what you decide to use.
//
// Gofig It aims to provide a simple set of interfaces and API's to make it easy for users to implement
// their own parsers beyond those bundled within the parsers package.
//
// Example.
//
//   package main
//
//   import (
//       "go.krak3n.codes/gofig"
//       "go.krak3n.codes/gofig/parsers/toml" // because why aren't you using TOML?
//   )
//
//   type Config struct {
//       Foo string `gofig:"foo"`
//       Bar string `gofig:"bar"`
//   }
//
//   func main() {
//	     var cfg Config
//
//	     // Initialise gofig with the struct config values will be placed into
//	     gfg, err := gofig.New(&cfg)
//	     gofig.Must(err)
//
//	     // Parse so environment variables
//       gofig.Must(gfg.Parse(gfg.FromFile(toml.New(), "/pah/to/cfg.toml")))
//
//	     // Use the config
//	     fmt.Println("Foo:", cfg.Foo, "Bar:", cfg.Bar)
//   }
package gofig
