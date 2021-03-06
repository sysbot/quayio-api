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

// NewPrototypeActivatingUser Repository creating user to whom the rule should apply
type NewPrototypeActivatingUser struct {
	// The username for the activating_user
	Name string `json:"name"`
}

// NewNewPrototypeActivatingUser instantiates a new NewPrototypeActivatingUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewPrototypeActivatingUser(name string) *NewPrototypeActivatingUser {
	this := NewPrototypeActivatingUser{}
	this.Name = name
	return &this
}

// NewNewPrototypeActivatingUserWithDefaults instantiates a new NewPrototypeActivatingUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewPrototypeActivatingUserWithDefaults() *NewPrototypeActivatingUser {
	this := NewPrototypeActivatingUser{}
	return &this
}

// GetName returns the Name field value
func (o *NewPrototypeActivatingUser) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NewPrototypeActivatingUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NewPrototypeActivatingUser) SetName(v string) {
	o.Name = v
}

func (o NewPrototypeActivatingUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNewPrototypeActivatingUser struct {
	value *NewPrototypeActivatingUser
	isSet bool
}

func (v NullableNewPrototypeActivatingUser) Get() *NewPrototypeActivatingUser {
	return v.value
}

func (v *NullableNewPrototypeActivatingUser) Set(val *NewPrototypeActivatingUser) {
	v.value = val
	v.isSet = true
}

func (v NullableNewPrototypeActivatingUser) IsSet() bool {
	return v.isSet
}

func (v *NullableNewPrototypeActivatingUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewPrototypeActivatingUser(val *NewPrototypeActivatingUser) *NullableNewPrototypeActivatingUser {
	return &NullableNewPrototypeActivatingUser{value: val, isSet: true}
}

func (v NullableNewPrototypeActivatingUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewPrototypeActivatingUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
