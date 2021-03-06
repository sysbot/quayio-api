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

// AddLabel Adds a label to a manifest
type AddLabel struct {
	// The value for the label
	Value string `json:"value"`
	// The key for the label
	Key string `json:"key"`
}

// NewAddLabel instantiates a new AddLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLabel(value string, key string) *AddLabel {
	this := AddLabel{}
	this.Value = value
	this.Key = key
	return &this
}

// NewAddLabelWithDefaults instantiates a new AddLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLabelWithDefaults() *AddLabel {
	this := AddLabel{}
	return &this
}

// GetValue returns the Value field value
func (o *AddLabel) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AddLabel) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AddLabel) SetValue(v string) {
	o.Value = v
}

// GetKey returns the Key field value
func (o *AddLabel) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AddLabel) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AddLabel) SetKey(v string) {
	o.Key = v
}

func (o AddLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableAddLabel struct {
	value *AddLabel
	isSet bool
}

func (v NullableAddLabel) Get() *AddLabel {
	return v.value
}

func (v *NullableAddLabel) Set(val *AddLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLabel(val *AddLabel) *NullableAddLabel {
	return &NullableAddLabel{value: val, isSet: true}
}

func (v NullableAddLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
