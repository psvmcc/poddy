package commands

import (
	"log"
	"os"
	"poddy/pkg/types"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

func GetVolumes(c *cli.Context) error {
	if c.Args().Len() == 0 {
		var volumes types.Volume
		err := volumes.List(c.String("poddyendpoint"), c.String("poddytoken"), c.String("namespace"))
		if err != nil {
			log.Fatal(err)
		}
		header := []string{"Namespace", "Volume"}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(header)
		table.SetAutoWrapText(false)
		table.SetAutoFormatHeaders(true)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetHeaderLine(false)
		table.SetBorder(false)
		table.SetTablePadding("\t")
		table.SetNoWhiteSpace(true)
		table.AppendBulk(volumes)
		table.Render()
	} else {
		log.Fatalf("Unsupported arguments count: %d\n", c.Args().Len())
	}
	return nil
}
