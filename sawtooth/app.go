package sawtooth

import (
	"log"
	"os"

	"github.com/c360dev/sawmint/restapi"
	"github.com/c360dev/sawmint/restapi/operations"
	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
)

//Restapi simply tells the REST API to begin to exist
func Restapi() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewSawtoothRESTAPI(swaggerSpec)
	goapisrv := restapi.NewServer(api)
	defer goapisrv.Shutdown()

	//Specifies the port to operate on
	goapisrv.Port = 8008

	parser := flags.NewParser(goapisrv, flags.Default)
	parser.ShortDescription = "Sawtooth REST API"
	parser.LongDescription = "_This HTTP pragmatic REST API is built on top of Sawtooth's\nexisting ZMQ/Protobuf infrastructure, simplifying client\ninteraction with the blockchain by exposing endpoints that\nuse common HTTP/JSON standards._\n"

	goapisrv.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	goapisrv.ConfigureAPI()
	if err := goapisrv.Serve(); err != nil {
		log.Fatalln(err)
	}
}
