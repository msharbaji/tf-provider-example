package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/msharbaji/grpc-go-example/pkg/server"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"example": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(_ *testing.T) {

}

func TestMain(t *testing.M) {
	zerolog.SetGlobalLevel(zerolog.NoLevel)

	secrets := map[string]string{
		"my-key-id": "my-secret-key",
	}

	grpcServer, err := server.NewGrpcServer("localhost:50051", secrets)
	if err != nil {
		return
	}

	// Start the server
	go func() {
		err := grpcServer.Start()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to start grpc server")
		}
	}()

	// Run tests...
	t.Run()

}
