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

// checks if the ReserveNow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReserveNow{}

// ReserveNow struct for ReserveNow
type ReserveNow struct {
	ResponseUrl string `json:"response_url"`
	Token Token `json:"token"`
	ExpiryDate string `json:"expiry_date"`
	ReservationId string `json:"reservation_id"`
	LocationId string `json:"location_id"`
	EvseUid *string `json:"evse_uid,omitempty"`
	AuthorizationReference *string `json:"authorization_reference,omitempty"`
}

// NewReserveNow instantiates a new ReserveNow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReserveNow(responseUrl string, token Token, expiryDate string, reservationId string, locationId string) *ReserveNow {
	this := ReserveNow{}
	this.ResponseUrl = responseUrl
	this.Token = token
	this.ExpiryDate = expiryDate
	this.ReservationId = reservationId
	this.LocationId = locationId
	return &this
}

// NewReserveNowWithDefaults instantiates a new ReserveNow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReserveNowWithDefaults() *ReserveNow {
	this := ReserveNow{}
	return &this
}

// GetResponseUrl returns the ResponseUrl field value
func (o *ReserveNow) GetResponseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseUrl
}

// GetResponseUrlOk returns a tuple with the ResponseUrl field value
// and a boolean to check if the value has been set.
func (o *ReserveNow) GetResponseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseUrl, true
}

// SetResponseUrl sets field value
func (o *ReserveNow) SetResponseUrl(v string) {
	o.ResponseUrl = v
}

// GetToken returns the Token field value
func (o *ReserveNow) GetToken() Token {
	if o == nil {
		var ret Token
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ReserveNow) GetTokenOk() (*Token, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ReserveNow) SetToken(v Token) {
	o.Token = v
}

// GetExpiryDate returns the ExpiryDate field value
func (o *ReserveNow) GetExpiryDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value
// and a boolean to check if the value has been set.
func (o *ReserveNow) GetExpiryDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryDate, true
}

// SetExpiryDate sets field value
func (o *ReserveNow) SetExpiryDate(v string) {
	o.ExpiryDate = v
}

// GetReservationId returns the ReservationId field value
func (o *ReserveNow) GetReservationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value
// and a boolean to check if the value has been set.
func (o *ReserveNow) GetReservationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReservationId, true
}

// SetReservationId sets field value
func (o *ReserveNow) SetReservationId(v string) {
	o.ReservationId = v
}

// GetLocationId returns the LocationId field value
func (o *ReserveNow) GetLocationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value
// and a boolean to check if the value has been set.
func (o *ReserveNow) GetLocationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationId, true
}

// SetLocationId sets field value
func (o *ReserveNow) SetLocationId(v string) {
	o.LocationId = v
}

// GetEvseUid returns the EvseUid field value if set, zero value otherwise.
func (o *ReserveNow) GetEvseUid() string {
	if o == nil || IsNil(o.EvseUid) {
		var ret string
		return ret
	}
	return *o.EvseUid
}

// GetEvseUidOk returns a tuple with the EvseUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReserveNow) GetEvseUidOk() (*string, bool) {
	if o == nil || IsNil(o.EvseUid) {
		return nil, false
	}
	return o.EvseUid, true
}

// HasEvseUid returns a boolean if a field has been set.
func (o *ReserveNow) HasEvseUid() bool {
	if o != nil && !IsNil(o.EvseUid) {
		return true
	}

	return false
}

// SetEvseUid gets a reference to the given string and assigns it to the EvseUid field.
func (o *ReserveNow) SetEvseUid(v string) {
	o.EvseUid = &v
}

// GetAuthorizationReference returns the AuthorizationReference field value if set, zero value otherwise.
func (o *ReserveNow) GetAuthorizationReference() string {
	if o == nil || IsNil(o.AuthorizationReference) {
		var ret string
		return ret
	}
	return *o.AuthorizationReference
}

// GetAuthorizationReferenceOk returns a tuple with the AuthorizationReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReserveNow) GetAuthorizationReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationReference) {
		return nil, false
	}
	return o.AuthorizationReference, true
}

// HasAuthorizationReference returns a boolean if a field has been set.
func (o *ReserveNow) HasAuthorizationReference() bool {
	if o != nil && !IsNil(o.AuthorizationReference) {
		return true
	}

	return false
}

// SetAuthorizationReference gets a reference to the given string and assigns it to the AuthorizationReference field.
func (o *ReserveNow) SetAuthorizationReference(v string) {
	o.AuthorizationReference = &v
}

func (o ReserveNow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReserveNow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["response_url"] = o.ResponseUrl
	toSerialize["token"] = o.Token
	toSerialize["expiry_date"] = o.ExpiryDate
	toSerialize["reservation_id"] = o.ReservationId
	toSerialize["location_id"] = o.LocationId
	if !IsNil(o.EvseUid) {
		toSerialize["evse_uid"] = o.EvseUid
	}
	if !IsNil(o.AuthorizationReference) {
		toSerialize["authorization_reference"] = o.AuthorizationReference
	}
	return toSerialize, nil
}

type NullableReserveNow struct {
	value *ReserveNow
	isSet bool
}

func (v NullableReserveNow) Get() *ReserveNow {
	return v.value
}

func (v *NullableReserveNow) Set(val *ReserveNow) {
	v.value = val
	v.isSet = true
}

func (v NullableReserveNow) IsSet() bool {
	return v.isSet
}

func (v *NullableReserveNow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReserveNow(val *ReserveNow) *NullableReserveNow {
	return &NullableReserveNow{value: val, isSet: true}
}

func (v NullableReserveNow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReserveNow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

