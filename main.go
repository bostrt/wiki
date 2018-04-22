package main

import (
	"github.com/namsral/flag"
)

var (
	cfg Config
)

func main() {
	var (
		config string
		data   string
		bind   string
		brand  string
	)

	flag.StringVar(&config, "config", "", "config file")
	flag.StringVar(&data, "data", "./data", "path to data")
	flag.StringVar(&bind, "bind", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.StringVar(&brand, "brand", "Wiki", "branding at top of each page")
	flag.Parse()

	// TODO: Abstract the Config and Handlers better
	cfg.data = data
	cfg.brand = brand

	NewServer(bind, cfg).ListenAndServe()
}
