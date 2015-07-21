package goubi

import (
	"errors"
	"strings"
	"testing"
)

var generic_cloud_status *CloudStatus = &CloudStatus{Status: "success", StatusMessage: "stuff and things"}
var generic_cloud_status_json []byte = []byte(`{"status": "success", "statusmsg": "stuff and things"}`)

var cloud *Cloud = new(Cloud)

type mockUbiClient struct{}

func (c *mockUbiClient) makeRequest(method string, params ...map[string]string) *ubiRequest {
	url := getAPIURI(method)
	if len(params) > 0 {
		ubireq := &ubiRequest{url, params[0]}
		return ubireq
	}
	ubireq := &ubiRequest{url, make(map[string]string, 0)}
	return ubireq
}

func (c *mockUbiClient) call(ubireq *ubiRequest) ([]byte, error) {
	var res []byte

	index := strings.Index(ubireq.URL, "=")
	method := ubireq.URL[index+1:]

	switch method {
	case "cloud.add_key":
		res = []byte(`{"result": 112233}`)
	case "cloud.backup_change_schedule":
		res = generic_cloud_status_json
	case "cloud.backup_convert_to_template":
		res = generic_cloud_status_json
	case "cloud.backup_create":
		res = generic_cloud_status_json
	case "cloud.backup_delete":
		res = generic_cloud_status_json
	case "cloud.backup_list":
		res = []byte(`{"Backups":[{"id": "22"}, {"id": "88"}]}`)
	case "cloud.backup_restore":
		res = generic_cloud_status_json
	case "cloud.create":
		res = []byte(`{"vm":{"invoice_id":"9999"}}`)
	case "cloud.create_custom_flavor":
		res = []byte(`{"result":"12"}`)
	case "cloud.delete_custom_flavor":
		res = []byte(`{"result":true}`)
	case "cloud.delete_template":
		res = generic_cloud_status_json
	case "cloud.destroy":
		res = generic_cloud_status_json
	case "cloud.edit_label":
		res = generic_cloud_status_json
	case "cloud.get":
		res = []byte(`{"vm":{"vm_id":"123"}}`)
	case "cloud.get_balance":
		res = []byte(`{"amount":"9.99"}`)
	case "cloud.get_recovery_mode_status":
		res = []byte(`{"result":false}`)
	case "cloud.get_welcome_emails":
		res = []byte(`{"WelcomeEmails":[{"Id":"12"},{"Id":"13"}]}`)
	case "cloud.list_custom_flavors":
		res = []byte(`{"CustomFlavors":[{"id": "22"}, {"id": "88"}]}`)
	case "cloud.list_flavors":
		res = []byte(`{"Flavors":[{"Id":"12"},{"Id":"13"}]}`)
	case "cloud.list_images":
		res = []byte(`{"Images":[{"id": "12"}, {"id": "13"}]}`)
	case "cloud.list_keys":
		res = []byte(`{"Keys":[{"id": "12"}, {"id": "13"}]}`)
	case "cloud.list_templates":
		res = []byte(`{"Templates":[{"id": "12"}, {"id": "13"}]}`)
	case "cloud.list_vms":
		res = []byte(`{"vms":[{"vm_id":"123"}]}`)
	case "cloud.list_zones":
		res = []byte(`{"Zones":[{"Id":"12"},{"Id":"13"}]}`)
	case "cloud.reboot":
		res = generic_cloud_status_json
	case "cloud.rebuild":
		res = generic_cloud_status_json
	case "cloud.recover":
		res = generic_cloud_status_json
	case "cloud.remove_key":
		res = []byte(`{"result":true}`)
	case "cloud.reset_password":
		res = []byte(`{"rootpassword":"axc32v4x65635s"}`)
	case "cloud.resize":
		res = generic_cloud_status_json
	case "cloud.start":
		res = generic_cloud_status_json
	case "cloud.stop":
		res = generic_cloud_status_json
	case "cloud.unrecover":
		res = generic_cloud_status_json
	case "cloud.update_custom_flavor":
		res = []byte(`{"result":"13"}`)
	case "cloud.usage_history":
		res = []byte(`{"UsageHistory":[{"id": "22"}, {"id": "88"}]}`)
	default:
		return res, errors.New("HTTP ERROR")
	}
	return res, nil
}

func TestSetup(t *testing.T) {
	mock_ubi_client := new(mockUbiClient)
	cloud.apiClient = mock_ubi_client

}

func TestAddKey(t *testing.T) {
	key := new(AddKeyParams)
	key.KeyName = "SuperSecret"
	key.PubKey = "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAhSUJobfJ57J8pRxch8Vv19U+51AMY393YOTJGx5WdWf6eJhGZq6k/oIbJHK8aSLXhhyPxL2JoF87xMgdoLKNSpe//tgnHNhQIr0ZPnEe+zc066J1gsSIMiVkRKOhDqftOILbTFpqv+Md1AZ6qd+u8XZHuVjOO7VKBwwGX60AVuM="
	//	key.Fingerprint = "10:10:84:a2:f0:df:94:41:af:88:e0:a0:10:ab:e0:8b"
	res, err := cloud.AddKey(key)
	exp := 112233
	evalResponse(res, err, exp, t)
}
func TestBackupChangeSchedule(t *testing.T) {
	res, err := cloud.BackupChangeSchedule(123, "monthly", "my_backup", 3)
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}
func TestBackupConvertToTemplate(t *testing.T) {
	res, err := cloud.BackupConvertToTemplate(123, 41, "MyTemplate")
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}
func TestBackupCreate(t *testing.T) {
	res, err := cloud.BackupCreate(123, "MyBackup")
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}
func TestBackupDelete(t *testing.T) {
	res, err := cloud.BackupDelete(123, 22)
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}
func TestBackupList(t *testing.T) {
	res, err := cloud.BackupList(123)
	exp := []Backup{{Id: 22}, {Id: 88}}
	evalResponse(res, err, exp, t)
}
func TestBackupRestore(t *testing.T) {
	res, err := cloud.BackupRestore(123, 22)
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}
func TestCreate(t *testing.T) {
	new_vm := new(CreateVMParams)
	new_vm.Hostname = "MyNewVM"
	new_vm.ZoneID = 44
	new_vm.ImageID = 1
	new_vm.FlavorID = 1
	res, err := cloud.Create(new_vm)
	exp := &CreateVM{InvoiceID: 9999}
	evalResponse(res, err, exp, t)
}
func TestCreateCustomFlavor(t *testing.T) {
	new_custom_flavor := new(CreateCustomFlavorParams)
	new_custom_flavor.Name = "CustomLinux"
	new_custom_flavor.Disk = 20
	new_custom_flavor.RAM = 512
	new_custom_flavor.VCPUs = 1
	res, err := cloud.CreateCustomFlavor(new_custom_flavor)
	exp := 12
	evalResponse(res, err, exp, t)
}
func TestDeleteCustomFlavor(t *testing.T) {
	res, err := cloud.DeleteCustomFlavor(12)
	exp := true
	evalResponse(res, err, exp, t)
}

func TestDeleteTemplate(t *testing.T) {
	res, err := cloud.DeleteTemplate(41)
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}
func TestDestroy(t *testing.T) {
	res, err := cloud.Destroy(123)
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}
func TestEditLabel(t *testing.T) {
	res, err := cloud.EditLabel(123, "RenamedServer")
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}
func TestGet(t *testing.T) {
	res, err := cloud.Get(123)
	exp := &Instance{VmID: 123}
	evalResponse(res, err, exp, t)
}

func TestGetBalance(t *testing.T) {
	res, err := cloud.GetBalance()
	exp := 9.99
	evalResponse(res, err, exp, t)
}

func TestGetRecoveryModeStatus(t *testing.T) {
	res, err := cloud.GetRecoveryModeStatus(123)
	exp := false
	evalResponse(res, err, exp, t)
}

func TestGetWelcomeEmails(t *testing.T) {
	res, err := cloud.GetWelcomeEmails(123)
	exp := []WelcomeEmail{{Id: 12}, {Id: 13}}
	evalResponse(res, err, exp, t)
}

func TestListCustomFlavors(t *testing.T) {
	cfs := new(CustomFlavorSearch)
	cfs.CustomFlavorsId = 44
	cfs.Name = "MyCustomFlavor"
	res, err := cloud.ListCustomFlavors(cfs)
	exp := []CustomFlavor{{Id: 22}, {Id: 88}}
	evalResponse(res, err, exp, t)

}

func TestListCustomFlavorsNoExtraParams(t *testing.T) {
	res, err := cloud.ListCustomFlavors()
	exp := []CustomFlavor{{Id: 22}, {Id: 88}}
	evalResponse(res, err, exp, t)

}

func TestListFlavors(t *testing.T) {
	res, err := cloud.ListFlavors()
	exp := []Flavor{{Id: 12}, {Id: 13}}
	evalResponse(res, err, exp, t)
}

func TestListImages(t *testing.T) {
	res, err := cloud.ListImages()
	exp := []Image{{Id: 12}, {Id: 13}}
	evalResponse(res, err, exp, t)
}
func TestListKeys(t *testing.T) {
	res, err := cloud.ListKeys()
	exp := []Key{{Id: 12}, {Id: 13}}
	evalResponse(res, err, exp, t)
}

func TestListTemplates(t *testing.T) {
	res, err := cloud.ListTemplates()
	exp := []Template{{Id: 12}, {Id: 13}}
	evalResponse(res, err, exp, t)
}

func TestListVms(t *testing.T) {
	res, err := cloud.ListVms()
	exp := []Instance{{VmID: 123}}
	evalResponse(res, err, exp, t)
}

func TestListZones(t *testing.T) {
	res, err := cloud.ListZones()
	exp := []Zone{{Id: 12}, {Id: 13}}
	evalResponse(res, err, exp, t)
}

func TestReboot(t *testing.T) {
	res, err := cloud.Reboot(123)
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}
func TestRebuild(t *testing.T) {
	res, err := cloud.Rebuild(123, 1)
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}
func TestRecover(t *testing.T) {
	res, err := cloud.Recover(123)
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}
func TestRemoveKey(t *testing.T) {
	res, err := cloud.RemoveKey(123)
	exp := true
	evalResponse(res, err, exp, t)
}
func TestResetPassword(t *testing.T) {
	res, err := cloud.ResetPassword(123)
	exp := "axc32v4x65635s"
	evalResponse(res, err, exp, t)
}

func TestResize(t *testing.T) {
	res, err := cloud.Resize(123, 1)
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}

func TestStart(t *testing.T) {
	res, err := cloud.Start(123)
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}

func TestStop(t *testing.T) {
	res, err := cloud.Stop(123)
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}

func TestUnrecover(t *testing.T) {
	res, err := cloud.Unrecover(123)
	exp := generic_cloud_status
	evalResponse(res, err, exp, t)
}

func TestUpdateCustomFlavor(t *testing.T) {
	cfp := new(CreateCustomFlavorParams)
	cfp.Name = "ReallyCustomLinux"
	cfp.Disk = 40
	cfp.RAM = 1024
	cfp.VCPUs = 2
	res, err := cloud.UpdateCustomFlavor(12, cfp)
	exp := 13
	evalResponse(res, err, exp, t)
}

func TestUsageHistory(t *testing.T) {
	res, err := cloud.UsageHistory(123, &UsageHistorySearch{Start: 1000, End: 1430839955})
	exp := []UsageRecord{{Id: 22}, {Id: 88}}
	evalResponse(res, err, exp, t)
}

func TestCallError(t *testing.T) {
	var i interface{} = nil
	err := cloud.call("", &i)
	if err == nil {
		t.Error("No Error, this is a fail")
	}
}

func TestCallWithParamsError(t *testing.T) {
	var i interface{} = nil
	err := cloud.callWithParams("", make(map[string]string), &i)
	if err == nil {
		t.Error("No Error, this is a fail")
	}
}
