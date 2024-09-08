package commands

import (
	"fmt"
	"log"
	"os"
	"poddy/pkg/types"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

func GetConfigMap(c *cli.Context) error {
	if c.Args().Len() == 0 {
		var configmaps types.ConfigMap
		err := configmaps.List(c.String("poddyendpoint"), c.String("poddytoken"), c.String("namespace"))
		if err != nil {
			log.Fatal(err)
		}
		header := []string{"Namespace", "Configmap"}
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
		table.AppendBulk(configmaps)
		table.Render()
	} else if c.Args().Len() == 1 {
		var configmap types.ConfigMap
		cm, err := configmap.Get(c.String("poddyendpoint"), c.String("poddytoken"), c.String("namespace"), c.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(cm))
	} else {
		log.Fatalf("Unsupported arguments count: %d\n", c.Args().Len())
	}
	return nil
}

func DeleteConfigMap(c *cli.Context) error {
	if c.Args().Len() == 1 {
		var configmap types.ConfigMap
		err := configmap.Delete(c.String("poddyendpoint"), c.String("poddytoken"), c.String("namespace"), c.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatalf("Unsupported arguments count: %d\n", c.Args().Len())
	}
	return nil
}
