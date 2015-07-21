package goubi

import (
	"strconv"
)

/*****************************************************************************/

// Represents an array of SSH keys
type Keys struct {
	Keys []Key `json:"Keys"`
}

// Represents an SSH key
type Key struct {
	Id          int    `json:"id,string"`
	KeyName     string `json:"keyname"`
	PubKey      string `json:"pubkey"`
	Fingerprint string `json:"fingerprint"`
	DateAdded   int    `json:"date_added,string"`
}

// Represents the needed parameters for the creation of a new SSH key
type AddKeyParams struct {
	KeyName string
	PubKey  string
}

// Represents the returned ID from the creation of a new SSH key
type AddKeyId struct {
	Results int `json:"result"`
}

// Represents the result of the operation of removing an SSH key
type RemoveKeyResults struct {
	Results bool `json:"result"`
}

/*****************************************************************************/

// Adds a new SSH key to the list of keys you are able to use
func (c *Cloud) AddKey(akp *AddKeyParams) (int, error) {
	ak := &AddKeyId{}

	params := map[string]string{}
	params["keyname"] = akp.KeyName
	params["pubkey"] = akp.PubKey

	err := c.callWithParams("cloud.add_key", params, &ak)

	return ak.Results, err
}

// Lists the available SSH keys you are able to use
func (c *Cloud) ListKeys() ([]Key, error) {
	ky := &Keys{}
	err := c.call("cloud.list_keys", &ky)
	return ky.Keys, err
}

// Removes an SSH key
func (c *Cloud) RemoveKey(sshkey_id int) (bool, error) {
	rk := &RemoveKeyResults{}
	err := c.callWithParams("cloud.remove_key", map[string]string{"sshkey_id": strconv.Itoa(sshkey_id)}, &rk)
	return rk.Results, err
}
