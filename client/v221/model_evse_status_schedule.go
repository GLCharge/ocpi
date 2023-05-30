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

// checks if the EvseStatusSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvseStatusSchedule{}

// EvseStatusSchedule struct for EvseStatusSchedule
type EvseStatusSchedule struct {
	PeriodBegin string `json:"period_begin"`
	PeriodEnd *string `json:"period_end,omitempty"`
	Status string `json:"status"`
}

// NewEvseStatusSchedule instantiates a new EvseStatusSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvseStatusSchedule(periodBegin string, status string) *EvseStatusSchedule {
	this := EvseStatusSchedule{}
	this.PeriodBegin = periodBegin
	this.Status = status
	return &this
}

// NewEvseStatusScheduleWithDefaults instantiates a new EvseStatusSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvseStatusScheduleWithDefaults() *EvseStatusSchedule {
	this := EvseStatusSchedule{}
	return &this
}

// GetPeriodBegin returns the PeriodBegin field value
func (o *EvseStatusSchedule) GetPeriodBegin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeriodBegin
}

// GetPeriodBeginOk returns a tuple with the PeriodBegin field value
// and a boolean to check if the value has been set.
func (o *EvseStatusSchedule) GetPeriodBeginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodBegin, true
}

// SetPeriodBegin sets field value
func (o *EvseStatusSchedule) SetPeriodBegin(v string) {
	o.PeriodBegin = v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise.
func (o *EvseStatusSchedule) GetPeriodEnd() string {
	if o == nil || IsNil(o.PeriodEnd) {
		var ret string
		return ret
	}
	return *o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvseStatusSchedule) GetPeriodEndOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodEnd) {
		return nil, false
	}
	return o.PeriodEnd, true
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *EvseStatusSchedule) HasPeriodEnd() bool {
	if o != nil && !IsNil(o.PeriodEnd) {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given string and assigns it to the PeriodEnd field.
func (o *EvseStatusSchedule) SetPeriodEnd(v string) {
	o.PeriodEnd = &v
}

// GetStatus returns the Status field value
func (o *EvseStatusSchedule) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EvseStatusSchedule) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EvseStatusSchedule) SetStatus(v string) {
	o.Status = v
}

func (o EvseStatusSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvseStatusSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["period_begin"] = o.PeriodBegin
	if !IsNil(o.PeriodEnd) {
		toSerialize["period_end"] = o.PeriodEnd
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableEvseStatusSchedule struct {
	value *EvseStatusSchedule
	isSet bool
}

func (v NullableEvseStatusSchedule) Get() *EvseStatusSchedule {
	return v.value
}

func (v *NullableEvseStatusSchedule) Set(val *EvseStatusSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableEvseStatusSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableEvseStatusSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvseStatusSchedule(val *EvseStatusSchedule) *NullableEvseStatusSchedule {
	return &NullableEvseStatusSchedule{value: val, isSet: true}
}

func (v NullableEvseStatusSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvseStatusSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

