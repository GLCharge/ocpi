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

// checks if the ChargingProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargingProfile{}

// ChargingProfile struct for ChargingProfile
type ChargingProfile struct {
	StartDateTime *string `json:"start_date_time,omitempty"`
	Duration *int32 `json:"duration,omitempty"`
	ChargingRateUnit string `json:"charging_rate_unit"`
	MinChargingRate *float32 `json:"min_charging_rate,omitempty"`
	ChargingProfilePeriod *ChargingProfileChargingProfilePeriod `json:"charging_profile_period,omitempty"`
}

// NewChargingProfile instantiates a new ChargingProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargingProfile(chargingRateUnit string) *ChargingProfile {
	this := ChargingProfile{}
	this.ChargingRateUnit = chargingRateUnit
	return &this
}

// NewChargingProfileWithDefaults instantiates a new ChargingProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargingProfileWithDefaults() *ChargingProfile {
	this := ChargingProfile{}
	return &this
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *ChargingProfile) GetStartDateTime() string {
	if o == nil || IsNil(o.StartDateTime) {
		var ret string
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingProfile) GetStartDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartDateTime) {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *ChargingProfile) HasStartDateTime() bool {
	if o != nil && !IsNil(o.StartDateTime) {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given string and assigns it to the StartDateTime field.
func (o *ChargingProfile) SetStartDateTime(v string) {
	o.StartDateTime = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *ChargingProfile) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingProfile) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *ChargingProfile) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *ChargingProfile) SetDuration(v int32) {
	o.Duration = &v
}

// GetChargingRateUnit returns the ChargingRateUnit field value
func (o *ChargingProfile) GetChargingRateUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChargingRateUnit
}

// GetChargingRateUnitOk returns a tuple with the ChargingRateUnit field value
// and a boolean to check if the value has been set.
func (o *ChargingProfile) GetChargingRateUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChargingRateUnit, true
}

// SetChargingRateUnit sets field value
func (o *ChargingProfile) SetChargingRateUnit(v string) {
	o.ChargingRateUnit = v
}

// GetMinChargingRate returns the MinChargingRate field value if set, zero value otherwise.
func (o *ChargingProfile) GetMinChargingRate() float32 {
	if o == nil || IsNil(o.MinChargingRate) {
		var ret float32
		return ret
	}
	return *o.MinChargingRate
}

// GetMinChargingRateOk returns a tuple with the MinChargingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingProfile) GetMinChargingRateOk() (*float32, bool) {
	if o == nil || IsNil(o.MinChargingRate) {
		return nil, false
	}
	return o.MinChargingRate, true
}

// HasMinChargingRate returns a boolean if a field has been set.
func (o *ChargingProfile) HasMinChargingRate() bool {
	if o != nil && !IsNil(o.MinChargingRate) {
		return true
	}

	return false
}

// SetMinChargingRate gets a reference to the given float32 and assigns it to the MinChargingRate field.
func (o *ChargingProfile) SetMinChargingRate(v float32) {
	o.MinChargingRate = &v
}

// GetChargingProfilePeriod returns the ChargingProfilePeriod field value if set, zero value otherwise.
func (o *ChargingProfile) GetChargingProfilePeriod() ChargingProfileChargingProfilePeriod {
	if o == nil || IsNil(o.ChargingProfilePeriod) {
		var ret ChargingProfileChargingProfilePeriod
		return ret
	}
	return *o.ChargingProfilePeriod
}

// GetChargingProfilePeriodOk returns a tuple with the ChargingProfilePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingProfile) GetChargingProfilePeriodOk() (*ChargingProfileChargingProfilePeriod, bool) {
	if o == nil || IsNil(o.ChargingProfilePeriod) {
		return nil, false
	}
	return o.ChargingProfilePeriod, true
}

// HasChargingProfilePeriod returns a boolean if a field has been set.
func (o *ChargingProfile) HasChargingProfilePeriod() bool {
	if o != nil && !IsNil(o.ChargingProfilePeriod) {
		return true
	}

	return false
}

// SetChargingProfilePeriod gets a reference to the given ChargingProfileChargingProfilePeriod and assigns it to the ChargingProfilePeriod field.
func (o *ChargingProfile) SetChargingProfilePeriod(v ChargingProfileChargingProfilePeriod) {
	o.ChargingProfilePeriod = &v
}

func (o ChargingProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargingProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDateTime) {
		toSerialize["start_date_time"] = o.StartDateTime
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	toSerialize["charging_rate_unit"] = o.ChargingRateUnit
	if !IsNil(o.MinChargingRate) {
		toSerialize["min_charging_rate"] = o.MinChargingRate
	}
	if !IsNil(o.ChargingProfilePeriod) {
		toSerialize["charging_profile_period"] = o.ChargingProfilePeriod
	}
	return toSerialize, nil
}

type NullableChargingProfile struct {
	value *ChargingProfile
	isSet bool
}

func (v NullableChargingProfile) Get() *ChargingProfile {
	return v.value
}

func (v *NullableChargingProfile) Set(val *ChargingProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableChargingProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableChargingProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargingProfile(val *ChargingProfile) *NullableChargingProfile {
	return &NullableChargingProfile{value: val, isSet: true}
}

func (v NullableChargingProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargingProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

