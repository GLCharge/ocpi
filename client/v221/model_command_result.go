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

// checks if the CommandResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommandResult{}

// CommandResult struct for CommandResult
type CommandResult struct {
	Result string `json:"result"`
	Message *CommandResponseMessage `json:"message,omitempty"`
}

// NewCommandResult instantiates a new CommandResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandResult(result string) *CommandResult {
	this := CommandResult{}
	this.Result = result
	return &this
}

// NewCommandResultWithDefaults instantiates a new CommandResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandResultWithDefaults() *CommandResult {
	this := CommandResult{}
	return &this
}

// GetResult returns the Result field value
func (o *CommandResult) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CommandResult) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CommandResult) SetResult(v string) {
	o.Result = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CommandResult) GetMessage() CommandResponseMessage {
	if o == nil || IsNil(o.Message) {
		var ret CommandResponseMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResult) GetMessageOk() (*CommandResponseMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CommandResult) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given CommandResponseMessage and assigns it to the Message field.
func (o *CommandResult) SetMessage(v CommandResponseMessage) {
	o.Message = &v
}

func (o CommandResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommandResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCommandResult struct {
	value *CommandResult
	isSet bool
}

func (v NullableCommandResult) Get() *CommandResult {
	return v.value
}

func (v *NullableCommandResult) Set(val *CommandResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandResult(val *CommandResult) *NullableCommandResult {
	return &NullableCommandResult{value: val, isSet: true}
}

func (v NullableCommandResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

