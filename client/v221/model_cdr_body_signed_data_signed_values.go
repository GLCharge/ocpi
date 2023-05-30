/*
OCPI modules

Specification for OCPIs modules handlers

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OCPI

import (
	"encoding/json"
)

// checks if the CdrBodySignedDataSignedValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdrBodySignedDataSignedValues{}

// CdrBodySignedDataSignedValues struct for CdrBodySignedDataSignedValues
type CdrBodySignedDataSignedValues struct {
	Nature string `json:"nature"`
	PlainData string `json:"plain_data"`
	SignedData string `json:"signed_data"`
}

// NewCdrBodySignedDataSignedValues instantiates a new CdrBodySignedDataSignedValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdrBodySignedDataSignedValues(nature string, plainData string, signedData string) *CdrBodySignedDataSignedValues {
	this := CdrBodySignedDataSignedValues{}
	this.Nature = nature
	this.PlainData = plainData
	this.SignedData = signedData
	return &this
}

// NewCdrBodySignedDataSignedValuesWithDefaults instantiates a new CdrBodySignedDataSignedValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdrBodySignedDataSignedValuesWithDefaults() *CdrBodySignedDataSignedValues {
	this := CdrBodySignedDataSignedValues{}
	return &this
}

// GetNature returns the Nature field value
func (o *CdrBodySignedDataSignedValues) GetNature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nature
}

// GetNatureOk returns a tuple with the Nature field value
// and a boolean to check if the value has been set.
func (o *CdrBodySignedDataSignedValues) GetNatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nature, true
}

// SetNature sets field value
func (o *CdrBodySignedDataSignedValues) SetNature(v string) {
	o.Nature = v
}

// GetPlainData returns the PlainData field value
func (o *CdrBodySignedDataSignedValues) GetPlainData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlainData
}

// GetPlainDataOk returns a tuple with the PlainData field value
// and a boolean to check if the value has been set.
func (o *CdrBodySignedDataSignedValues) GetPlainDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlainData, true
}

// SetPlainData sets field value
func (o *CdrBodySignedDataSignedValues) SetPlainData(v string) {
	o.PlainData = v
}

// GetSignedData returns the SignedData field value
func (o *CdrBodySignedDataSignedValues) GetSignedData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedData
}

// GetSignedDataOk returns a tuple with the SignedData field value
// and a boolean to check if the value has been set.
func (o *CdrBodySignedDataSignedValues) GetSignedDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedData, true
}

// SetSignedData sets field value
func (o *CdrBodySignedDataSignedValues) SetSignedData(v string) {
	o.SignedData = v
}

func (o CdrBodySignedDataSignedValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdrBodySignedDataSignedValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nature"] = o.Nature
	toSerialize["plain_data"] = o.PlainData
	toSerialize["signed_data"] = o.SignedData
	return toSerialize, nil
}

type NullableCdrBodySignedDataSignedValues struct {
	value *CdrBodySignedDataSignedValues
	isSet bool
}

func (v NullableCdrBodySignedDataSignedValues) Get() *CdrBodySignedDataSignedValues {
	return v.value
}

func (v *NullableCdrBodySignedDataSignedValues) Set(val *CdrBodySignedDataSignedValues) {
	v.value = val
	v.isSet = true
}

func (v NullableCdrBodySignedDataSignedValues) IsSet() bool {
	return v.isSet
}

func (v *NullableCdrBodySignedDataSignedValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdrBodySignedDataSignedValues(val *CdrBodySignedDataSignedValues) *NullableCdrBodySignedDataSignedValues {
	return &NullableCdrBodySignedDataSignedValues{value: val, isSet: true}
}

func (v NullableCdrBodySignedDataSignedValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdrBodySignedDataSignedValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

