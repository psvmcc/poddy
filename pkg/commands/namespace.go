package commands

import (
	"fmt"
	"log"
	"os"
	"poddy/pkg/types"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func GetNamespace(c *cli.Context) error {
	if c.Args().Len() == 0 {
		var namespaces types.Namespaces
		err := namespaces.Get(c.String("poddyendpoint"), c.String("poddytoken"))
		if err != nil {
			log.Fatal(err)
		}
		header := []string{"Namespace", "Access Type", "Podman Network"}
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
		table.AppendBulk(namespaces)
		table.Render()
	} else if c.Args().Len() == 1 {
		var namespace types.Namespace
		err := namespace.Get(c.String("poddyendpoint"), c.String("poddytoken"), c.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
		yamlData, _ := yaml.Marshal(namespace)
		fmt.Println(string(yamlData))
	} else {
		log.Fatalf("Unsupported arguments count: %d\n", c.Args().Len())
	}
	return nil
}
