package client

import (
	"fmt"
	"golang.org/x/net/context"
	//"fmt"
	//"golang.org/x/oauth2/google"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
	"log"
	//"golang.org/x/net/context"
	"net/http"
)

func GetCloudSQLInfo(projectID string,c *http.Client,ctx context.Context) []string {
	sqladminService, err := sqladmin.New(c)
	if err != nil {
log.Fatal(err)
}
	req := sqladminService.Instances.List(projectID)
	var DatabaseList []string
	if err := req.Pages(ctx, func(page *sqladmin.InstancesListResponse) error {
		for _, databaseInstance := range page.Items {
			DatabaseList =append(DatabaseList,fmt.Sprintf("%#v",databaseInstance))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return DatabaseList
}
