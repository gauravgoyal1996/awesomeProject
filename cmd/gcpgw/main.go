package main

import (
	"awesomeProject/infra/gcp"
)

func main() {

	// Project ID for this request.
	project := "sacred-alloy-250406"

	// Project Zone for this request.
	zone := "us-central1-b"

	gcp.Data(project, zone)
}
