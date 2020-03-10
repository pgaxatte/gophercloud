package endpointgroups

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

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
	_, r.Err = client.Post(rootURL(client), &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

// Get retrieves details on a single endpoint group, by ID.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToEndpointGroupUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts provides options for updating a role.
type UpdateOpts struct {
	// Name is the name of the new endpoint group.
	Name string `json:"name,omitempty"`

	// Filters is an EndpointFilter type describing the filter criteria
	Filters EndpointFilter `json:"filters,omitempty"`

	// Description is the description of the endpoint group
	Description string `json:"description,omitempty"`
}

// ToEndpointGroupUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToEndpointGroupUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "endpoint_group")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Update will update an existing endpoint group.
func Update(client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToEndpointGroupUpdateMap()
	if err != nil {
		r.Err = err
		return r
	}

	_, r.Err = client.Patch(resourceURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete deletes the specified endpoint group.
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
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
func (opts ListOpts) ToEndointGroupListQuery() (string, error) {
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

/*
// Endpoint group <-> project association
func Associate()              {}
func GetAssociation()         {}
func Dissociate()             {}
func ListAssociatedProjects() {}
func ListAssociatedEnpoints() {}

// Project <-> endpoint associations
func AssociateProjectToEndpoint()           {}
func IsEndpointAssociated()                 {}
func DissociateProjectToEndpoint()          {}
func ListProjectsAssociatedToEndpoint()     {}
func ListEndpointsAssociatedToProject()     {}
func ListEnpointGroupsAssociatedToProject() {}
*/
