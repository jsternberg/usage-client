package client_test

import (
	"fmt"
	"io/ioutil"

	"github.com/influxdb/enterprise-client/v1"
)

// Example of saving Usage data to Enterprise
func Example_saveUsage() {
	c := client.New("token-goes-here")
	// override the URL for testing
	c.URL = "https://enterprise.staging.influxdata.com"

	u := client.Usage{
		Product: "influxdb",
		Data: []client.UsageData{
			{
				Tags: client.Tags{
					"version": "0.9.5",
					"arch":    "amd64",
					"os":      "linux",
				},
				Values: client.Values{
					"cluster_id":       "23423",
					"server_id":        "1",
					"num_databases":    3,
					"num_measurements": 2342,
					"num_series":       87232,
				},
			},
		},
	}

	res, err := c.Save(u)
	fmt.Printf("err: %s\n", err)
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("b: %s\n", b)
}
