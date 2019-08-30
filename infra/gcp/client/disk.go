package client

import (
"fmt"
"golang.org/x/net/context"
"google.golang.org/api/compute/v1"
"log"
)

func GetDiskList(project string,zone string,computeService *compute.Service,ctx context.Context)[]string {
	req := computeService.Disks.List(project,zone)
	var list []string
	if err := req.Pages(ctx, func(page *compute.DiskList) error {
		for _, disk := range page.Items {
			list = append(list,fmt.Sprintf("%#v",disk))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return list
}
