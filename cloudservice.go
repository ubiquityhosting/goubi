package goubi

import (
	"strconv"
)

/*****************************************************************************/

// Represents a standard cloud action response
type CloudStatus struct {
	Status        string `json:"status"`
	StatusMessage string `json:"statusmsg"`
}

// Interface for cloud API methods
type cloudService interface {
	AddKey(*AddKeyParams) (int, error)
	BackupChangeSchedule(vm_id int, schedule string, name string, number_of_backups int) (*CloudStatus, error)
	BackupConvertToTemplate(vm_id int, backup_id int, name string) (*CloudStatus, error)
	BackupCreate(vm_id int, name string) (*CloudStatus, error)
	BackupDelete(vm_id int, backup_id int) (*CloudStatus, error)
	BackupList(vm_id int) ([]Backup, error)
	BackupRestore(vm_id int, backup_id int) (*CloudStatus, error)
	Create(*CreateVMParams) (*CreateVM, error)
	CreateCustomFlavor(*CreateCustomFlavorParams) (int, error)
	DeleteCustomFlavor(custom_flavors_id int) (bool, error)
	DeleteTemplate(template_id int) (*CloudStatus, error)
	Destroy(vm_id int) (*CloudStatus, error)
	EditLabel(vm_id int, new_label string) (*CloudStatus, error)
	Get(vm_id int) (*Instance, error)
	GetBalance() (float64, error)
	GetRecoveryModeStatus(vm_id int) (bool, error)
	GetWelcomeEmails(vm_id int) ([]WelcomeEmail, error)
	ListCustomFlavors(...*CustomFlavorSearch) ([]CustomFlavor, error)
	ListFlavors() ([]Flavor, error)
	ListImages() ([]Image, error)
	ListKeys() ([]Key, error)
	ListTemplates() ([]Template, error)
	ListVms() ([]Instance, error)
	ListZones() ([]Zone, error)
	Reboot(vm_id int) (*CloudStatus, error)
	Rebuild(vm_id int, image_id int) (*CloudStatus, error)
	Recover(vm_id int) (*CloudStatus, error)
	RemoveKey(sshkey_id int) (bool, error)
	ResetPassword(vm_id int) (string, error)
	Resize(vm_id int, flavor_id int) (*CloudStatus, error)
	Start(vm_id int) (*CloudStatus, error)
	Stop(vm_id int) (*CloudStatus, error)
	Unrecover(vm_id int) (*CloudStatus, error)
	UpdateCustomFlavor(custom_flavors_id int, ccfp *CreateCustomFlavorParams) (int, error)
	UsageHistory(vm_id int, uhs ...*UsageHistorySearch) ([]UsageRecord, error)
}

// Structure for cloud API methods
type Cloud struct {
	apiClient postClient
}

// Represents the returned sum of current hourly billing charges
type Balance struct {
	Amount float64 `json:"amount,string"`
}

// Represents the returned recovery mode status of a virtual machine
type RecoveryModeStatus struct {
	Status bool `json:"RecoveryModeStatus"`
}

// Represents the reset password result of a cloud instance
type NewPassword struct {
	Password string `json:"rootpassword"`
}

/*****************************************************************************/

func (c *Cloud) call(method string, i interface{}) error {
	ubireq := c.apiClient.makeRequest(method)
	res, err := c.apiClient.call(ubireq)
	if err != nil {
		return err
	}
	err = unmarshalToStruct(res, &i)
	return err
}

func (c *Cloud) callWithParams(method string, params map[string]string, i interface{}) error {
	ubireq := c.apiClient.makeRequest(method, params)
	res, err := c.apiClient.call(ubireq)
	if err != nil {
		return err
	}
	err = unmarshalToStruct(res, &i)
	return err
}

/*****************************************************************************/

// Queues a new Ubiquity Cloud VM to be built. The new VM will be associated with your client account once created and
// your welcome information will be emailed to the email associated with your account. A new VM for your account will
// not be created until the returned invoice is paid.
func (c *Cloud) Create(cvmp *CreateVMParams) (*CreateVM, error) {

	cv := &CreateVMBase{}
	params := map[string]string{}
	params["hostname"] = cvmp.Hostname
	params["zone_id"] = strconv.Itoa(cvmp.ZoneID)
	params["image_id"] = strconv.Itoa(cvmp.ImageID)
	params["flavor_id"] = strconv.Itoa(cvmp.FlavorID)
	params["key_id"] = strconv.Itoa(cvmp.KeyID)
	params["userdata"] = cvmp.Userdata
	params["docker_machine"] = strconv.FormatBool(cvmp.DockerMachine)

	err := c.callWithParams("cloud.create", params, &cv)
	return cv.VM, err
}

// Processes the immediate deletion of the provided hourly VM
func (c *Cloud) Destroy(vm_id int) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.destroy", map[string]string{"vm_id": strconv.Itoa(vm_id)}, &cs)
	return cs, err
}

// Updates a label for an indicated VM
func (c *Cloud) EditLabel(vm_id int, new_label string) (*CloudStatus, error) {
	el := &CloudStatus{}
	err := c.callWithParams("cloud.edit_label", map[string]string{"vm_id": strconv.Itoa(vm_id), "new_label": new_label}, &el)
	return el, err
}

// Returns the sum of current hourly billing charges for this billing period
func (c *Cloud) GetBalance() (float64, error) {
	bal := &Balance{}
	err := c.call("cloud.get_balance", &bal)
	return bal.Amount, err
}

// Returns whether or not a virtual machine is in recovery mode
func (c *Cloud) GetRecoveryModeStatus(vm_id int) (bool, error) {
	rms := &RecoveryModeStatus{}
	err := c.callWithParams("cloud.get_recovery_mode_status", map[string]string{"vm_id": strconv.Itoa(vm_id)}, &rms)
	return rms.Status, err
}

// Rebuilds a VM
func (c *Cloud) Rebuild(vm_id int, image_id int) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.rebuild", map[string]string{"vm_id": strconv.Itoa(vm_id), "image_id": strconv.Itoa(image_id)}, &cs)
	return cs, err
}

// Attempts to gracefully reboot the given VM
func (c *Cloud) Reboot(vm_id int) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.reboot", map[string]string{"vm_id": strconv.Itoa(vm_id)}, &cs)
	return cs, err
}

// Reboots the virtual machine into a custom recovery kernel environment that can aid in repairing the virtual machine
func (c *Cloud) Recover(vm_id int) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.recover", map[string]string{"vm_id": strconv.Itoa(vm_id)}, &cs)
	return cs, err
}

// Attempts to reset the root password on the given VM
func (c *Cloud) ResetPassword(vm_id int) (string, error) {
	pw := &NewPassword{}
	err := c.callWithParams("cloud.reset_password", map[string]string{"vm_id": strconv.Itoa(vm_id)}, &pw)
	return pw.Password, err
}

// Updates a VM to a new flavor - VMs can only be resized to a larger flavor
func (c *Cloud) Resize(vm_id int, flavor_id int) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.resize", map[string]string{"vm_id": strconv.Itoa(vm_id), "flavor_id": strconv.Itoa(flavor_id)}, &cs)
	return cs, err
}

// Starts up the indicated VM
func (c *Cloud) Start(vm_id int) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.start", map[string]string{"vm_id": strconv.Itoa(vm_id)}, &cs)
	return cs, err
}

// Stops the indicated VM
func (c *Cloud) Stop(vm_id int) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.stop", map[string]string{"vm_id": strconv.Itoa(vm_id)}, &cs)
	return cs, err
}

// Reboots the virtual machine back to the normal booted environment
func (c *Cloud) Unrecover(vm_id int) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.unrecover", map[string]string{"vm_id": strconv.Itoa(vm_id)}, &cs)
	return cs, err
}
