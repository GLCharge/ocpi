/*
OCPI modules

Specification for OCPIs modules handlers

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OCPI

import (
	"encoding/json"
	"fmt"
)

// EnvironmentalImpactCategoryType Categories of environmental impact values
type EnvironmentalImpactCategoryType string

// List of environmentalImpactCategoryType
const (
	NUCLEAR_WASTE EnvironmentalImpactCategoryType = "NUCLEAR_WASTE"
	CARBON_DIOXIDE EnvironmentalImpactCategoryType = "CARBON_DIOXIDE"
)

// All allowed values of EnvironmentalImpactCategoryType enum
var AllowedEnvironmentalImpactCategoryTypeEnumValues = []EnvironmentalImpactCategoryType{
	"NUCLEAR_WASTE",
	"CARBON_DIOXIDE",
}

func (v *EnvironmentalImpactCategoryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnvironmentalImpactCategoryType(value)
	for _, existing := range AllowedEnvironmentalImpactCategoryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnvironmentalImpactCategoryType", value)
}

// NewEnvironmentalImpactCategoryTypeFromValue returns a pointer to a valid EnvironmentalImpactCategoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnvironmentalImpactCategoryTypeFromValue(v string) (*EnvironmentalImpactCategoryType, error) {
	ev := EnvironmentalImpactCategoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnvironmentalImpactCategoryType: valid values are %v", v, AllowedEnvironmentalImpactCategoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnvironmentalImpactCategoryType) IsValid() bool {
	for _, existing := range AllowedEnvironmentalImpactCategoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to environmentalImpactCategoryType value
func (v EnvironmentalImpactCategoryType) Ptr() *EnvironmentalImpactCategoryType {
	return &v
}

type NullableEnvironmentalImpactCategoryType struct {
	value *EnvironmentalImpactCategoryType
	isSet bool
}

func (v NullableEnvironmentalImpactCategoryType) Get() *EnvironmentalImpactCategoryType {
	return v.value
}

func (v *NullableEnvironmentalImpactCategoryType) Set(val *EnvironmentalImpactCategoryType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentalImpactCategoryType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentalImpactCategoryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentalImpactCategoryType(val *EnvironmentalImpactCategoryType) *NullableEnvironmentalImpactCategoryType {
	return &NullableEnvironmentalImpactCategoryType{value: val, isSet: true}
}

func (v NullableEnvironmentalImpactCategoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentalImpactCategoryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
