package goubi

/*****************************************************************************/

// Represents an array of cloud instance flavors
type Flavors struct {
	Flavors []Flavor `json:"Flavors"`
}

// Represents a cloud instance flavor
type Flavor struct {
	Id         int     `json:"id,string"`
	Name       string  `json:"name"`
	CPU        int     `json:"cpu,string"`
	RAM        float64 `json:"ram"`
	Disk       int     `json:"disk,string"`
	BW         int     `json:"bw,string"`
	HourlyCost float64 `json:"hourly_cost"`
	Cost       int     `json:"cost,string"`
}

/*****************************************************************************/

// Lists the available plans
func (c *Cloud) ListFlavors() ([]Flavor, error) {
	fl := &Flavors{}
	err := c.call("cloud.list_flavors", &fl)
	return fl.Flavors, err
}
