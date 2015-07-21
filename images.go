package goubi

/*****************************************************************************/

// Represents an array of cloud instance images
type Images struct {
	Images []Image `json:"Images"`
}

// Represents a cloud instance image
type Image struct {
	Id         int     `json:"id,string"`
	Name       string  `json:"name"`
	CatId      int     `json:"cat_id,string"`
	CatDesc    string  `json:"cat_desc"`
	Cost       float64 `json:"cost"`
	Disk       int     `json:"disk,string"`
	RAM        int     `json:"ram,string"`
	IPs        int     `json:"ips,string"`
	UdId       int     `json:"ud_id,string"`
	UdTemplate string  `json:"ud_template"`
}

/*****************************************************************************/

// Lists the available images you are able to use
func (c *Cloud) ListImages() ([]Image, error) {
	im := &Images{}
	err := c.call("cloud.list_images", &im)
	return im.Images, err
}
