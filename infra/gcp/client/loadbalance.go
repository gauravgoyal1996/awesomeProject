package client

import (
	"fmt"
	"log"
	"google.golang.org/api/compute/v1"
	"golang.org/x/net/context"
)

func GetLoadBalancerList(project string,computeService *compute.Service,ctx context.Context) []string{
	req := computeService.BackendServices.List(project)
	fmt.Println("list of loadbalancers")
	var list []string
	if err := req.Pages(ctx, func(page *compute.BackendServiceList) error {
		for _, backendService := range page.Items {
			list =append(list,fmt.Sprintf("%#v",backendService))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return list
}
