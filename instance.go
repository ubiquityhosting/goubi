package goubi

import (
	"strconv"
)

/*****************************************************************************/

// Represents an array of cloud instances
type Instances struct {
	Vms []Instance `json:"vms"`
}

// Represents the return value of a cloud instance
type InstanceBase struct {
	Vm *Instance `json:"vm"`
}

// Represents a cloud instance
type Instance struct {
	VmID                         int    `json:"vm_id,string"`
	ServiceID                    int    `json:"service_id,string"`
	ZoneID                       int    `json:"zone_id,string"`
	FlavorID                     int    `json:"flavor_id,string"`
	MainIPaddress                string `json:"mainipaddress"`
	IPaddresses                  string `json:"ipaddresses"`
	InternalIPs                  string `json:"internalips"`
	Type                         string `json:"type"`
	State                        string `json:"state"`
	Memory                       int    `json:"memory"`
	Hdd                          int    `json:"hdd"`
	TrafficGraph                 string `json:"trafficgraph"`
	LoadGraph                    string `json:"loadgraph"`
	MemoryGraph                  string `json:"memorygraph"`
	Hostname                     string `json:"hostname"`
	Template                     string `json:"template"`
	MAC                          string `json:"mac"`
	Swapburst                    int    `json:"swap-burst"`
	ZoneName                     string `json:"zone_name"`
	TemplateID                   int    `json:"template_id,string"`
	RootPassword                 string `json:"rootpassword"`
	UserPassword                 string `json:"userpassword"`
	VncIP                        string `json:"vncip"`
	VNCPort                      int    `json:"vncport,string"`
	VNCPassword                  string `json:"vncpassword"`
	OrderID                      int    `json:"order_id,string"`
	InRecovery                   int    `json:"in_recovery"`
	Locked                       int    `json:"locked,string"`
	LockedReason                 string `json:"locked_reason"`
	NearOverageSentOn            string `json:"near_overage_sent_on"`
	FullOverageSentOn            string `json:"full_overage_sent_on"`
	IsBareMetal                  int    `json:"is_bare_metal,string"`
	HddFriendly                  string `json:"hdd_friendly"`
	MemoryFriendly               string `json:"memory_friendly"`
	BandwidthUsed                int    `json:"bandwidth_used,string"`
	BandwidthTotal               int    `json:"bandwidth_total,string"`
	BandwidthFree                int    `json:"bandwidth_free,string"`
	BandwidthPercentUsed         int    `json:"bandwidth_percent_used,string"`
	BandwidthUsedFriendly        string `json:"bandwidth_used_friendly"`
	BandwidthTotalFriendly       string `json:"bandwidth_total_friendly"`
	BandwidthFreeFriendly        string `json:"bandwidth_free_friendly"`
	BandwidthPercentUsedFriendly string `json:"bandwidth_percent_used_friendly"`
}

// Represents the return value of a pending deployment for a new cloud instance
type CreateVMBase struct {
	VM *CreateVM `json:"vm"`
}

// Represents a pending deployment for a new cloud instance
type CreateVM struct {
	InvoiceID int     `json:"invoice_id,string"`
	ServiceID int     `json:"service_id,string"`
	Balance   float64 `json:"balance"`
	OrderID   int     `json:"order_id,string"`
}

// Represents the needed parameters for the creation of a cloud instance deployment
type CreateVMParams struct {
	Hostname      string
	ImageID       int
	FlavorID      int
	ZoneID        int
	KeyID         int
	Userdata      string
	DockerMachine bool
}

/*****************************************************************************/

// Returns information about the provided VM
func (c *Cloud) Get(vm_id int) (*Instance, error) {
	cs := &InstanceBase{}
	err := c.callWithParams("cloud.get", map[string]string{"vm_id": strconv.Itoa(vm_id)}, &cs)
	return cs.Vm, err
}

// Returns identifying information about a customer's Ubiquity Cloud VMs
func (c *Cloud) ListVms() ([]Instance, error) {
	cs := &Instances{}
	err := c.call("cloud.list_vms", &cs)
	return cs.Vms, err
}
