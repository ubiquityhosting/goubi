package goubi

/*****************************************************************************/

// Represents an array of zones
type Zones struct {
	Zones []Zone `json:"Zones"`
}

// Represents a zone (a geographical location of where the instance is hosted)
type Zone struct {
	Id     int    `json:"id,string"`
	Name   string `json:"name"`
	Public int    `json:"public,string"`
}

/*****************************************************************************/

// Lists the available zones you are able to use
func (c *Cloud) ListZones() ([]Zone, error) {
	zn := &Zones{}
	err := c.call("cloud.list_zones", &zn)
	return zn.Zones, err
}
