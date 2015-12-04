package client_test

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"testing"

	"github.com/influxdb/usage-client/v1"
	"github.com/stretchr/testify/require"
)

func Test_Usage_Path(t *testing.T) {
	r := require.New(t)
	u := client.Usage{Product: "influxdb"}
	r.Equal("/usage/influxdb", u.Path())
}

func Test_NewUsage(t *testing.T) {
	r := require.New(t)
	u := client.NewUsage("influxdb", []client.UsageData{
		{
			Tags: client.Tags{
				"k1": "v1",
			},
			Values: client.Values{
				"num": 123,
			},
		},
		{
			Tags: client.Tags{
				"k2": "v2",
			},
			Values: client.Values{
				"str": "hello",
			},
		},
	})

	r.Equal(u.Product, "influxdb")
	r.Len(u.Data, 3)

	// Only asserting on the interesting tags -
	// the values are an implementation detail out of scope of tests.
	r.Equal(u.Data[0].Tags["os"], runtime.GOOS)
	r.Equal(u.Data[0].Tags["arch"], runtime.GOARCH)

	r.Equal(u.Data[1], client.UsageData{
		Tags: client.Tags{
			"k1": "v1",
		},
		Values: client.Values{
			"num": 123,
		},
	})
	r.Equal(u.Data[2], client.UsageData{
		Tags: client.Tags{
			"k2": "v2",
		},
		Values: client.Values{
			"str": "hello",
		},
	})
}

// Example of saving Usage data to the Usage API
func Example_saveUsage() {
	c := client.New("token-goes-here")
	// override the URL for testing
	c.URL = "https://usage.staging.influxdata.com"

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
