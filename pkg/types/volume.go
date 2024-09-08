package types

import (
	"fmt"
	"poddy/pkg/helpers"
)

type Volume [][]string

func (n *Volume) List(serverEndpoint, token, namespace string) error {
	url := fmt.Sprintf("%s/api/v1/namespaces/%s/volumes", serverEndpoint, namespace)
	return helpers.QueryFiles(url, token, fmt.Sprintf("poddy/%s (%s)", Version, Commit), n)
}
