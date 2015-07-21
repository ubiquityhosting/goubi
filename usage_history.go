package goubi

import (
	"strconv"
)

/*****************************************************************************/

// Represents an array of usage history records
type UsageHistory struct {
	UsageRecords []UsageRecord `json:"UsageHistory"`
}

// Represents a usage history record
type UsageRecord struct {
	Id        int `json:"id,string"`
	RAM       int `json:"ram,string"`
	Disk      int `json:"disk,string"`
	Backup    int `json:"bu,string"`
	Templates int `json:"tmp,string"`
	CPU       int `json:"cpu,string"`
	Bandwidth int `json:"bw,string"`
	Timestamp int `json:"timestamp,string"`
}

// Represents the optional parameters for listing usage history
type UsageHistorySearch struct {
	Start int
	End   int
}

/*****************************************************************************/

// Lists the billing breakdown of past virtual machine usage based on a period defined by the customer (start, end).
func (c *Cloud) UsageHistory(vm_id int, uhs ...*UsageHistorySearch) ([]UsageRecord, error) {
	uh := &UsageHistory{}

	var params = map[string]string{}
	params["vm_id"] = strconv.Itoa(vm_id)

	if len(uhs) > 0 {

		if uhs[0].Start != 0 {
			params["start"] = strconv.Itoa(uhs[0].Start)
		}
		if uhs[0].End != 0 {
			params["end"] = strconv.Itoa(uhs[0].End)
		}
	}

	err := c.callWithParams("cloud.usage_history", params, &uh)
	return uh.UsageRecords, err
}
