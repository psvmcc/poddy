package commands

import (
	"fmt"
	"log"
	"poddy/pkg/types"

	"github.com/urfave/cli/v2"
)

func Info(c *cli.Context) error {
	fmt.Printf("Client version %s [%s]\n", types.Version, types.Commit)

	var info types.Info
	err := info.Get(c.String("poddyendpoint"), c.String("poddytoken"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server version %s [%s]\n", info.Version, info.Commit)
	return nil
}
