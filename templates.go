package goubi

import (
	"strconv"
)

/*****************************************************************************/

// Represents an array of templates
type Templates struct {
	Templates []Template `json:"Templates"`
}

// Represents a template
type Template struct {
	Id          int     `json:"id,string"`
	Name        string  `json:"name"`
	Size        int     `json:"size,string"`
	CreatedOn   int     `json:"created_on,string"`
	MonthlyCost float64 `json:"monthly_cost"`
	ZoneId      int     `json:"zone_id,string"`
}

/*****************************************************************************/

// Deletes a user defined templates
func (c *Cloud) DeleteTemplate(template_id int) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.delete_template", map[string]string{"template_id": strconv.Itoa(template_id)}, &cs)
	return cs, err
}

// Lists the user made templates available
func (c *Cloud) ListTemplates() ([]Template, error) {
	tm := &Templates{}
	err := c.call("cloud.list_templates", &tm)
	return tm.Templates, err
}
