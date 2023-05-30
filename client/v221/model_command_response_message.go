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

// checks if the CommandResponseMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommandResponseMessage{}

// CommandResponseMessage struct for CommandResponseMessage
type CommandResponseMessage struct {
	Language string `json:"language"`
	Text string `json:"text"`
}

// NewCommandResponseMessage instantiates a new CommandResponseMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandResponseMessage(language string, text string) *CommandResponseMessage {
	this := CommandResponseMessage{}
	this.Language = language
	this.Text = text
	return &this
}

// NewCommandResponseMessageWithDefaults instantiates a new CommandResponseMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandResponseMessageWithDefaults() *CommandResponseMessage {
	this := CommandResponseMessage{}
	return &this
}

// GetLanguage returns the Language field value
func (o *CommandResponseMessage) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *CommandResponseMessage) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *CommandResponseMessage) SetLanguage(v string) {
	o.Language = v
}

// GetText returns the Text field value
func (o *CommandResponseMessage) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *CommandResponseMessage) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *CommandResponseMessage) SetText(v string) {
	o.Text = v
}

func (o CommandResponseMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommandResponseMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["language"] = o.Language
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

type NullableCommandResponseMessage struct {
	value *CommandResponseMessage
	isSet bool
}

func (v NullableCommandResponseMessage) Get() *CommandResponseMessage {
	return v.value
}

func (v *NullableCommandResponseMessage) Set(val *CommandResponseMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandResponseMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandResponseMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandResponseMessage(val *CommandResponseMessage) *NullableCommandResponseMessage {
	return &NullableCommandResponseMessage{value: val, isSet: true}
}

func (v NullableCommandResponseMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandResponseMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

