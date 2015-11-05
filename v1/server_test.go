package client_test

import (
	"fmt"
	"io/ioutil"

	"github.com/influxdb/enterprise-client/v1"
)

// Example of saving a server to Enterprise
func Example_saveServer() {
	c := client.New("token-goes-here")
	// override the URL for testing
	c.URL = "https://enterprise.staging.influxdata.com"

	s := client.Server{
		ClusterID: "clus1",
		Host:      "example.com",
		Product:   "jambox",
		Version:   "1.0",
		ServerID:  "serv1",
	}

	res, err := c.Save(s)
	fmt.Printf("err: %s\n", err)
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("b: %s\n", b)
}
