package client

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/compute/v1"
	"log"
)

func GetInstanceList(project string,zone string,computeService *compute.Service,ctx context.Context)[]string {
	req := computeService.Instances.List(project, zone)
	var list []string
	if err := req.Pages(ctx, func(page *compute.InstanceList) error {
		for _, instance := range page.Items {
			list = append(list,fmt.Sprintf("%#v",instance))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return list
}

