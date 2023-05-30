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

// checks if the CommandResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommandResponse{}

// CommandResponse struct for CommandResponse
type CommandResponse struct {
	Result string `json:"result"`
	Timeout int32 `json:"timeout"`
	Message *CommandResponseMessage `json:"message,omitempty"`
}

// NewCommandResponse instantiates a new CommandResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandResponse(result string, timeout int32) *CommandResponse {
	this := CommandResponse{}
	this.Result = result
	this.Timeout = timeout
	return &this
}

// NewCommandResponseWithDefaults instantiates a new CommandResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandResponseWithDefaults() *CommandResponse {
	this := CommandResponse{}
	return &this
}

// GetResult returns the Result field value
func (o *CommandResponse) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CommandResponse) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CommandResponse) SetResult(v string) {
	o.Result = v
}

// GetTimeout returns the Timeout field value
func (o *CommandResponse) GetTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *CommandResponse) GetTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *CommandResponse) SetTimeout(v int32) {
	o.Timeout = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CommandResponse) GetMessage() CommandResponseMessage {
	if o == nil || IsNil(o.Message) {
		var ret CommandResponseMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResponse) GetMessageOk() (*CommandResponseMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CommandResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given CommandResponseMessage and assigns it to the Message field.
func (o *CommandResponse) SetMessage(v CommandResponseMessage) {
	o.Message = &v
}

func (o CommandResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommandResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["timeout"] = o.Timeout
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCommandResponse struct {
	value *CommandResponse
	isSet bool
}

func (v NullableCommandResponse) Get() *CommandResponse {
	return v.value
}

func (v *NullableCommandResponse) Set(val *CommandResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandResponse(val *CommandResponse) *NullableCommandResponse {
	return &NullableCommandResponse{value: val, isSet: true}
}

func (v NullableCommandResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

