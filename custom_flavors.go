package goubi

import (
	"strconv"
)

/*****************************************************************************/

// Represents an array of custom flavors
type CustomFlavors struct {
	CustomFlavors []CustomFlavor `json:"CustomFlavors"`
}

// Represents a custom flavor
type CustomFlavor struct {
	Id    int     `json:"id,string"`
	Name  string  `json:"name"`
	RAM   int     `json:"ram,string"`
	Disk  int     `json:"disk,string"`
	VCPUs int     `json:"vcpus,string"`
	RamGb float64 `json:"ram_gig"`
	CPU   int     `json:"cpu,string"`
}

// Represents the optional parameters for listing custom flavors
type CustomFlavorSearch struct {
	Name            string
	CustomFlavorsId int
}

// Represents the needed parameters for the creation of a new custom flavor
type CreateCustomFlavorParams struct {
	Name  string
	RAM   int
	Disk  int
	VCPUs int
}

// Represents the result of the operation of removing an SSH key
type DeleteCustomFlavorResults struct {
	Results bool `json:"result"`
}

// Represents the returned ID from the creation of a new custom flavor
type CustomFlavorId struct {
	Results int `json:"result,string"`
}

/*****************************************************************************/

// Creates a new Custom User Flavor to be used on a user's Private Cloud.
// Returns the Custom Flavor ID of the newly created flavor.
func (c *Cloud) CreateCustomFlavor(ccfp *CreateCustomFlavorParams) (int, error) {
	cf := &CustomFlavorId{}
	params := map[string]string{}
	params["name"] = ccfp.Name
	params["ram"] = strconv.Itoa(ccfp.RAM)
	params["disk"] = strconv.Itoa(ccfp.Disk)
	params["vcpus"] = strconv.Itoa(ccfp.VCPUs)
	err := c.callWithParams("cloud.create_custom_flavor", params, &cf)
	return cf.Results, err
}

// Deletes a Custom Flavor. Returns TRUE on success.
func (c *Cloud) DeleteCustomFlavor(custom_flavors_id int) (bool, error) {
	cs := &DeleteCustomFlavorResults{}
	err := c.callWithParams("cloud.delete_custom_flavor", map[string]string{"custom_flavors_id": strconv.Itoa(custom_flavors_id)}, &cs)
	return cs.Results, err
}

// Returns a list of Custom Flavors to be used on a Private Cloud.
//
// If no parameters are passed (using a CustomFlavorSearch struct) it will return all
// of the custom flavors owned by you.
//
// If a CustomFlavorSearch struct is defined and passed, one or two optional parameters
// may then be passed to allow searching on a name or the custom flavor ID, within all
// of the custom flavors owned by you.
//
// The CustomFlavorSearch struct is defined as follows:
//
// 		CustomFlavorSearch{
//			Name: "customname"
//			CustomFlavorsId: 17
//		}
//
func (c *Cloud) ListCustomFlavors(cfs ...*CustomFlavorSearch) ([]CustomFlavor, error) {
	cf := &CustomFlavors{}

	if len(cfs) > 0 {
		var params = map[string]string{}
		if cfs[0].Name != "" {
			params["name"] = cfs[0].Name
		}
		if cfs[0].CustomFlavorsId != 0 {
			params["custom_flavors_id"] = strconv.Itoa(cfs[0].CustomFlavorsId)
		}
		err := c.callWithParams("cloud.list_custom_flavors", params, &cf)
		return cf.CustomFlavors, err
	}

	err := c.call("cloud.list_custom_flavors", &cf)
	return cf.CustomFlavors, err
}

// Updates a Custom User Flavor to be used on a Private Cloud.
// Returns the Custom Flavor ID of the newly updated flavor.
func (c *Cloud) UpdateCustomFlavor(custom_flavors_id int, ccfp *CreateCustomFlavorParams) (int, error) {
	cf := &CustomFlavorId{}
	params := map[string]string{}
	params["name"] = ccfp.Name
	params["ram"] = strconv.Itoa(ccfp.RAM)
	params["disk"] = strconv.Itoa(ccfp.Disk)
	params["vcpus"] = strconv.Itoa(ccfp.VCPUs)
	params["custom_flavors_id"] = strconv.Itoa(custom_flavors_id)
	err := c.callWithParams("cloud.update_custom_flavor", params, &cf)
	return cf.Results, err
}
