package goubi

import (
	"strconv"
)

/*****************************************************************************/

// Represents an array of backups
type Backups struct {
	Backups []Backup `json:"Backups"`
}

// Represents a backup
type Backup struct {
	Id             int    `json:"id,string"`
	RawSize        int    `json:"raw_size,string"`
	CompressedSize int    `json:"compressed_size,string"`
	ImageAlias     string `json:"image_alias"`
	CreatedOn      int    `json:"created_on,string"`
	Scheduled      string `json:"scheduled"`
}

/*****************************************************************************/

// Changes the backup schedule of the provided VM
func (c *Cloud) BackupChangeSchedule(vm_id int, schedule string, name string, number_of_backups int) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams(
		"cloud.backup_change_schedule",
		map[string]string{"vm_id": strconv.Itoa(vm_id), "schedule": schedule, "name": name, "number_of_backups": strconv.Itoa(number_of_backups)},
		&cs)
	return cs, err
}

func (c *Cloud) BackupConvertToTemplate(vm_id int, backup_id int, name string) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.backup_convert_to_template",
		map[string]string{"vm_id": strconv.Itoa(vm_id), "backup_id": strconv.Itoa(backup_id), "name": name},
		&cs)
	return cs, err
}

// Creates a backup of the provided VM
func (c *Cloud) BackupCreate(vm_id int, name string) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.backup_create", map[string]string{"vm_id": strconv.Itoa(vm_id), "name": name}, &cs)
	return cs, err
}

// Removes a backup of an indicated VM
func (c *Cloud) BackupDelete(vm_id int, backup_id int) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.backup_delete",
		map[string]string{"vm_id": strconv.Itoa(vm_id), "backup_id": strconv.Itoa(backup_id)},
		&cs)
	return cs, err
}

// Lists the available backups for the provided VM
func (c *Cloud) BackupList(vm_id int) ([]Backup, error) {
	bu := &Backups{}
	err := c.callWithParams("cloud.backup_list", map[string]string{"vm_id": strconv.Itoa(vm_id)}, &bu)
	return bu.Backups, err
}

// Restores a backup to an indicated VM
func (c *Cloud) BackupRestore(vm_id int, backup_id int) (*CloudStatus, error) {
	cs := &CloudStatus{}
	err := c.callWithParams("cloud.backup_restore", map[string]string{"vm_id": strconv.Itoa(vm_id), "backup_id": strconv.Itoa(backup_id)}, &cs)
	return cs, err
}
