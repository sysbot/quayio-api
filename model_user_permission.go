/*
 * Quay Frontend
 *
 * This API allows you to perform many of the operations required to work with Quay repositories, users, and organizations. You can find out more at <a href=\"https://quay.io\">Quay</a>.
 *
 * API version: v1
 * Contact: support@quay.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UserPermission Description of a user permission.
type UserPermission struct {
	// Role to use for the user
	Role string `json:"role"`
}

// NewUserPermission instantiates a new UserPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPermission(role string) *UserPermission {
	this := UserPermission{}
	this.Role = role
	return &this
}

// NewUserPermissionWithDefaults instantiates a new UserPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPermissionWithDefaults() *UserPermission {
	this := UserPermission{}
	return &this
}

// GetRole returns the Role field value
func (o *UserPermission) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UserPermission) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UserPermission) SetRole(v string) {
	o.Role = v
}

func (o UserPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableUserPermission struct {
	value *UserPermission
	isSet bool
}

func (v NullableUserPermission) Get() *UserPermission {
	return v.value
}

func (v *NullableUserPermission) Set(val *UserPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPermission(val *UserPermission) *NullableUserPermission {
	return &NullableUserPermission{value: val, isSet: true}
}

func (v NullableUserPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}