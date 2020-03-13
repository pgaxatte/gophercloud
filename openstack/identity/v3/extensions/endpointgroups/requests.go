package endpointgroups

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Get retrieves details on a single endpoint group, by ID.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToEndpointGroupListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// Name filters the response by endpoint group name.
	Name string `q:"name"`
}

// ToEndpointGroupListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToEndpointGroupListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List enumerates the endpoint groups
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(client)
	if opts != nil {
		query, err := opts.ToEndpointGroupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return EndpointGroupPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToEndpointGroupCreateMap() (map[string]interface{}, error)
}

// CreateOpts provides options used to create an endpoint group.
type CreateOpts struct {
	// Name is the name of the new endpoint group.
	Name string `json:"name" required:"true"`

	// Filters is an EndpointFilter type describing the filter criteria
	Filters EndpointFilter `json:"filters" required:"true"`

	// Description is the description of the endpoint group
	Description string `json:"description,omitempty"`
}

// ToEndpointGroupCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToEndpointGroupCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "endpoint_group")
}

// Create creates a new endpoint group
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToEndpointGroupCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(rootURL(client), &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
