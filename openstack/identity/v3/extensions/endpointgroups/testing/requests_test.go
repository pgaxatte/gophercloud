package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/extensions/endpointgroups"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreateSuccessful(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/OS-EP-FILTER/endpoint_groups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestJSONRequest(t, r, `
      {
        "endpoint_group": {
		  "description": "endpoint group test",
		  "filters": {
		    "interface": "public",
		    "service_id": "1234",
			"region_id": "5678",
			"enabled": true
		  },
          "name": "endpointgroup1"
        }
      }
    `)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `
      {
        "endpoint_group": {
          "id": "ac4861",
		  "filters": {
		    "interface": "public",
		    "service_id": "1234",
			"region_id": "5678",
			"enabled": true
		  },
          "name": "endpointgroup1",
		  "description": "endpoint group test",
          "links": {
            "self": "https://localhost:5000/v3/OS-EP-FILTER/endpoint_groups/ac4861"
          }
        }
      }
    `)
	})

	enabled := true
	filters := endpointgroups.EndpointFilter{
		Availability: gophercloud.AvailabilityPublic,
		ServiceID:    "1234",
		RegionID:     "5678",
		Enabled:      &enabled,
	}

	actual, err := endpointgroups.Create(client.ServiceClient(), endpointgroups.CreateOpts{
		Name:        "endpointgroup1",
		Description: "endpoint group test",
		Filters:     filters,
	}).Extract()
	th.AssertNoErr(t, err)

	expected := &endpointgroups.EndpointGroup{
		ID:          "ac4861",
		Name:        "endpointgroup1",
		Description: "endpoint group test",
		Filters:     filters,
	}

	th.AssertDeepEquals(t, expected, actual)
}
