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

// checks if the CancelReservation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelReservation{}

// CancelReservation struct for CancelReservation
type CancelReservation struct {
	ResponseUrl string `json:"response_url"`
	ReservationId string `json:"reservation_id"`
}

// NewCancelReservation instantiates a new CancelReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelReservation(responseUrl string, reservationId string) *CancelReservation {
	this := CancelReservation{}
	this.ResponseUrl = responseUrl
	this.ReservationId = reservationId
	return &this
}

// NewCancelReservationWithDefaults instantiates a new CancelReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelReservationWithDefaults() *CancelReservation {
	this := CancelReservation{}
	return &this
}

// GetResponseUrl returns the ResponseUrl field value
func (o *CancelReservation) GetResponseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseUrl
}

// GetResponseUrlOk returns a tuple with the ResponseUrl field value
// and a boolean to check if the value has been set.
func (o *CancelReservation) GetResponseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseUrl, true
}

// SetResponseUrl sets field value
func (o *CancelReservation) SetResponseUrl(v string) {
	o.ResponseUrl = v
}

// GetReservationId returns the ReservationId field value
func (o *CancelReservation) GetReservationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value
// and a boolean to check if the value has been set.
func (o *CancelReservation) GetReservationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReservationId, true
}

// SetReservationId sets field value
func (o *CancelReservation) SetReservationId(v string) {
	o.ReservationId = v
}

func (o CancelReservation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelReservation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["response_url"] = o.ResponseUrl
	toSerialize["reservation_id"] = o.ReservationId
	return toSerialize, nil
}

type NullableCancelReservation struct {
	value *CancelReservation
	isSet bool
}

func (v NullableCancelReservation) Get() *CancelReservation {
	return v.value
}

func (v *NullableCancelReservation) Set(val *CancelReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelReservation(val *CancelReservation) *NullableCancelReservation {
	return &NullableCancelReservation{value: val, isSet: true}
}

func (v NullableCancelReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

