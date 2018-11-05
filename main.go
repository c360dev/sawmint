//This is the ABCI to Sawtooth application.  It's an ABCI client and REST server, among other things.
//We'll be splitting things into bytes and feeding them to Tendermint.

package main

import (
	"os"

	"github.com/c360dev/sawmint/sawtooth"
	"github.com/tendermint/tendermint/abci/example/counter"
	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/libs/log"
)

func main() {

	//entrypoint for tendermint abci server
	var serial bool
	serial = true
	app := counter.NewCounterApplication(serial)
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	//starts a tendermint counter
	sawmint, err := server.NewServer("127.0.0.1:26658", "socket", app)
	if err != nil {
		return
	}
	//keeps starting the tendemrint counter.  This stanza and the one above are acting strangely, with regard to error handling.  Should be considered broken till fixed (that's when it says return err, instead of just return....)
	sawmint.SetLogger(logger.With("module", "abci-server"))
	if err := sawmint.Start(); err != nil {
		return
	}

	sawtooth.Restapi()

}
