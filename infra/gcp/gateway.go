package gcp

import (
	"awesomeProject/infra/gcp/client"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"log"
	"net/http"
)

func Data(project string,zone string){
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	InstanceData(project, zone, computeService,ctx)
	LoadBalancerData(project,computeService,ctx)
	CloudStorageData(project,ctx)
	CloudSQLData(project,c,ctx)
	DiskData(project, zone, computeService,ctx)

}

func InstanceData(project string,zone string,computeService *compute.Service,ctx context.Context){
	fmt.Println("Instance Data")
	list := client.GetInstanceList(project, zone, computeService,ctx)
	fmt.Println("list of instances")
	fmt.Println(list)
}

func LoadBalancerData(project string,computeService *compute.Service,ctx context.Context) {
	fmt.Println("LoadBalancer Data")
	list := client.GetLoadBalancerList(project,computeService,ctx)
	fmt.Println(list)
}

func CloudStorageData(project string,ctx context.Context) {
	 buckets := client.GetCloudStorageInfo(project,ctx)
	 fmt.Println("list of gc buckets")
	 fmt.Println(buckets)
}

func CloudSQLData(project string,c *http.Client,ctx context.Context) {
	CloudSQLList := client.GetCloudSQLInfo(project,c,ctx)
	fmt.Println("list of Cloud SQL databases")
	fmt.Println(CloudSQLList)
}

func DiskData(project string,zone string,computeService *compute.Service,ctx context.Context){
	fmt.Println("Disks Data")
	list := client.GetDiskList(project, zone, computeService,ctx)
	fmt.Println("list of disks")
	fmt.Println(list)
}
