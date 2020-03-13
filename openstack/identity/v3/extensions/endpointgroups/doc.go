/*
Package endpointgroups enables management of OpenStack Identity Endpoint Groups
and Endpoint associations.

Example to Get an Endpoint Group

	err := endpointgroups.Get(identityClient, endpointGroupID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to List all Endpoint Groups by name

	listOpts := endpointgropus.ListOpts{
		Name: "mygroup",
	}

	allPages, err := endpointgroups.List(identityClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allGroups, err := endpointgroups.ExtractEndpointGroups(allPages)
	if err != nil {
		panic(err)
	}

	for _, endpointgroup := range allGroups {
		fmt.Printf("%+v\n", endpointgroup)
	}

Example to Create an Endpoint Group

	createOpts := endpointgroups.CreateOpts{
		Name:        "my-ep-group",
		Description: "My endpoint group",
		Filters:     endpointgroups.EndpointFilter{
			Availability: gophercloud.AvailabilityPublic,
			ServiceID:    "1234",
			RegionID:     "5678",
		},
	}

	endpointGroup, err := endpointgroups.Create(identityClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}
*/
package endpointgroups
