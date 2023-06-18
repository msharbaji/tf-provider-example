package main

import (
	"context"
	"flag"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/msharbaji/tf-provider-example/internal/provider"
)

var (
	version = "0.0.0"
)

//go:generate terraform fmt -recursive ../examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --rendered-website-dir ../docs --examples-dir ../examples --provider-name example
func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "malsharbaji.com/providers/example",
		Debug:   debug,
	}

	// Logger
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	err := providerserver.Serve(context.Background(), provider.New(version), opts)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start the provider")
	}
}
