/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spacetraders

import (
	"encoding/json"
)

// checks if the ShipReactor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipReactor{}

// ShipReactor The reactor of the ship. The reactor is responsible for powering the ship's systems and weapons.
type ShipReactor struct {
	Symbol string `json:"symbol"`
	Name string `json:"name"`
	Description string `json:"description"`
	// Condition is a range of 0 to 100 where 0 is completely worn out and 100 is brand new.
	Condition *int32 `json:"condition,omitempty"`
	PowerOutput int32 `json:"powerOutput"`
	Requirements ShipRequirements `json:"requirements"`
}

// NewShipReactor instantiates a new ShipReactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipReactor(symbol string, name string, description string, powerOutput int32, requirements ShipRequirements) *ShipReactor {
	this := ShipReactor{}
	this.Symbol = symbol
	this.Name = name
	this.Description = description
	this.PowerOutput = powerOutput
	this.Requirements = requirements
	return &this
}

// NewShipReactorWithDefaults instantiates a new ShipReactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipReactorWithDefaults() *ShipReactor {
	this := ShipReactor{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ShipReactor) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ShipReactor) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ShipReactor) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *ShipReactor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShipReactor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShipReactor) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ShipReactor) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ShipReactor) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ShipReactor) SetDescription(v string) {
	o.Description = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ShipReactor) GetCondition() int32 {
	if o == nil || IsNil(o.Condition) {
		var ret int32
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipReactor) GetConditionOk() (*int32, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ShipReactor) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given int32 and assigns it to the Condition field.
func (o *ShipReactor) SetCondition(v int32) {
	o.Condition = &v
}

// GetPowerOutput returns the PowerOutput field value
func (o *ShipReactor) GetPowerOutput() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PowerOutput
}

// GetPowerOutputOk returns a tuple with the PowerOutput field value
// and a boolean to check if the value has been set.
func (o *ShipReactor) GetPowerOutputOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerOutput, true
}

// SetPowerOutput sets field value
func (o *ShipReactor) SetPowerOutput(v int32) {
	o.PowerOutput = v
}

// GetRequirements returns the Requirements field value
func (o *ShipReactor) GetRequirements() ShipRequirements {
	if o == nil {
		var ret ShipRequirements
		return ret
	}

	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
func (o *ShipReactor) GetRequirementsOk() (*ShipRequirements, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// SetRequirements sets field value
func (o *ShipReactor) SetRequirements(v ShipRequirements) {
	o.Requirements = v
}

func (o ShipReactor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipReactor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	toSerialize["powerOutput"] = o.PowerOutput
	toSerialize["requirements"] = o.Requirements
	return toSerialize, nil
}

type NullableShipReactor struct {
	value *ShipReactor
	isSet bool
}

func (v NullableShipReactor) Get() *ShipReactor {
	return v.value
}

func (v *NullableShipReactor) Set(val *ShipReactor) {
	v.value = val
	v.isSet = true
}

func (v NullableShipReactor) IsSet() bool {
	return v.isSet
}

func (v *NullableShipReactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipReactor(val *ShipReactor) *NullableShipReactor {
	return &NullableShipReactor{value: val, isSet: true}
}

func (v NullableShipReactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipReactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


