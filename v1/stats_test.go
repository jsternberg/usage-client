package client_test

import (
	"fmt"
	"io/ioutil"

	"github.com/influxdb/enterprise-client/v1"
)

// Example of saving Stats data to Enterprise
func Example_saveStats() {
	c := client.New("token-goes-here")
	// override the URL for testing
	c.URL = "https://enterprise.staging.influxdata.com"

	st := client.Stats{
		ClusterID: "clus1",
		ServerID:  "serv1",
		Product:   "influxdb",
		Data: []client.StatsData{
			client.StatsData{
				Name: "engine",
				Tags: client.Tags{
					"path":    "/home/philip/.influxdb/data/_internal/monitor/1",
					"version": "bz1",
				},
				Values: client.Values{
					"blks_write":          39,
					"blks_write_bytes":    2421,
					"blks_write_bytes_c":  2202,
					"points_write":        39,
					"points_write_dedupe": 39,
				},
			},
		},
	}

	res, err := c.Save(st)
	fmt.Printf("err: %s\n", err)
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("b: %s\n", b)
}
