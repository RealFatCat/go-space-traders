/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the FactionTrait type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FactionTrait{}

// FactionTrait struct for FactionTrait
type FactionTrait struct {
	// The unique identifier of the trait.
	Symbol string `json:"symbol"`
	// The name of the trait.
	Name string `json:"name"`
	// A description of the trait.
	Description string `json:"description"`
}

// NewFactionTrait instantiates a new FactionTrait object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFactionTrait(symbol string, name string, description string) *FactionTrait {
	this := FactionTrait{}
	this.Symbol = symbol
	this.Name = name
	this.Description = description
	return &this
}

// NewFactionTraitWithDefaults instantiates a new FactionTrait object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFactionTraitWithDefaults() *FactionTrait {
	this := FactionTrait{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *FactionTrait) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *FactionTrait) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *FactionTrait) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *FactionTrait) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FactionTrait) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FactionTrait) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *FactionTrait) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FactionTrait) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FactionTrait) SetDescription(v string) {
	o.Description = v
}

func (o FactionTrait) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FactionTrait) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullableFactionTrait struct {
	value *FactionTrait
	isSet bool
}

func (v NullableFactionTrait) Get() *FactionTrait {
	return v.value
}

func (v *NullableFactionTrait) Set(val *FactionTrait) {
	v.value = val
	v.isSet = true
}

func (v NullableFactionTrait) IsSet() bool {
	return v.isSet
}

func (v *NullableFactionTrait) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFactionTrait(val *FactionTrait) *NullableFactionTrait {
	return &NullableFactionTrait{value: val, isSet: true}
}

func (v NullableFactionTrait) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFactionTrait) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


