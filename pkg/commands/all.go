package commands

import (
	"fmt"
	"log"
	"os"
	"poddy/pkg/types"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

func GetAll(c *cli.Context) error {
	if c.Args().Len() == 0 {

		var err error
		var header []string

		var configmaps types.ConfigMap
		var volumes types.Volume

		err = configmaps.List(c.String("poddyendpoint"), c.String("poddytoken"), c.String("namespace"))
		if err != nil {
			log.Fatal(err)
		}
		header = []string{"Namespace", "Configmap"}
		tableCM := tablewriter.NewWriter(os.Stdout)
		tableCM.SetHeader(header)
		tableCM.SetAutoWrapText(false)
		tableCM.SetAutoFormatHeaders(true)
		tableCM.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		tableCM.SetAlignment(tablewriter.ALIGN_LEFT)
		tableCM.SetCenterSeparator("")
		tableCM.SetColumnSeparator("")
		tableCM.SetRowSeparator("")
		tableCM.SetHeaderLine(false)
		tableCM.SetBorder(false)
		tableCM.SetTablePadding("\t")
		tableCM.SetNoWhiteSpace(true)
		tableCM.AppendBulk(configmaps)
		tableCM.Render()

		fmt.Println()

		err = volumes.List(c.String("poddyendpoint"), c.String("poddytoken"), c.String("namespace"))
		if err != nil {
			log.Fatal(err)
		}
		header = []string{"Namespace", "Volume"}
		tableV := tablewriter.NewWriter(os.Stdout)
		tableV.SetHeader(header)
		tableV.SetAutoWrapText(false)
		tableV.SetAutoFormatHeaders(true)
		tableV.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		tableV.SetAlignment(tablewriter.ALIGN_LEFT)
		tableV.SetCenterSeparator("")
		tableV.SetColumnSeparator("")
		tableV.SetRowSeparator("")
		tableV.SetHeaderLine(false)
		tableV.SetBorder(false)
		tableV.SetTablePadding("\t")
		tableV.SetNoWhiteSpace(true)
		tableV.AppendBulk(volumes)
		tableV.Render()

	} else {
		log.Fatalf("Unsupported arguments count: %d\n", c.Args().Len())
	}
	return nil
}
