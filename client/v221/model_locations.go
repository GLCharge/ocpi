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

// checks if the Locations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Locations{}

// Locations struct for Locations
type Locations struct {
	Data *LocationsData `json:"data,omitempty"`
	StatusCode float32 `json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
	TimeStamp *string `json:"timeStamp,omitempty"`
}

// NewLocations instantiates a new Locations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocations(statusCode float32) *Locations {
	this := Locations{}
	this.StatusCode = statusCode
	return &this
}

// NewLocationsWithDefaults instantiates a new Locations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationsWithDefaults() *Locations {
	this := Locations{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Locations) GetData() LocationsData {
	if o == nil || IsNil(o.Data) {
		var ret LocationsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locations) GetDataOk() (*LocationsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Locations) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LocationsData and assigns it to the Data field.
func (o *Locations) SetData(v LocationsData) {
	o.Data = &v
}

// GetStatusCode returns the StatusCode field value
func (o *Locations) GetStatusCode() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *Locations) GetStatusCodeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *Locations) SetStatusCode(v float32) {
	o.StatusCode = v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *Locations) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locations) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *Locations) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *Locations) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *Locations) GetTimeStamp() string {
	if o == nil || IsNil(o.TimeStamp) {
		var ret string
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locations) GetTimeStampOk() (*string, bool) {
	if o == nil || IsNil(o.TimeStamp) {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *Locations) HasTimeStamp() bool {
	if o != nil && !IsNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given string and assigns it to the TimeStamp field.
func (o *Locations) SetTimeStamp(v string) {
	o.TimeStamp = &v
}

func (o Locations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Locations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["status_code"] = o.StatusCode
	if !IsNil(o.StatusMessage) {
		toSerialize["status_message"] = o.StatusMessage
	}
	if !IsNil(o.TimeStamp) {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	return toSerialize, nil
}

type NullableLocations struct {
	value *Locations
	isSet bool
}

func (v NullableLocations) Get() *Locations {
	return v.value
}

func (v *NullableLocations) Set(val *Locations) {
	v.value = val
	v.isSet = true
}

func (v NullableLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocations(val *Locations) *NullableLocations {
	return &NullableLocations{value: val, isSet: true}
}

func (v NullableLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

