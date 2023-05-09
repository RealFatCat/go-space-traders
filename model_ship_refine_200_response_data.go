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

// checks if the ShipRefine200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipRefine200ResponseData{}

// ShipRefine200ResponseData struct for ShipRefine200ResponseData
type ShipRefine200ResponseData struct {
	Cargo ShipCargo `json:"cargo"`
	Cooldown Cooldown `json:"cooldown"`
	Produced []ShipRefine200ResponseDataProducedInner `json:"produced"`
	Consumed []ShipRefine200ResponseDataProducedInner `json:"consumed"`
}

// NewShipRefine200ResponseData instantiates a new ShipRefine200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipRefine200ResponseData(cargo ShipCargo, cooldown Cooldown, produced []ShipRefine200ResponseDataProducedInner, consumed []ShipRefine200ResponseDataProducedInner) *ShipRefine200ResponseData {
	this := ShipRefine200ResponseData{}
	this.Cargo = cargo
	this.Cooldown = cooldown
	this.Produced = produced
	this.Consumed = consumed
	return &this
}

// NewShipRefine200ResponseDataWithDefaults instantiates a new ShipRefine200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipRefine200ResponseDataWithDefaults() *ShipRefine200ResponseData {
	this := ShipRefine200ResponseData{}
	return &this
}

// GetCargo returns the Cargo field value
func (o *ShipRefine200ResponseData) GetCargo() ShipCargo {
	if o == nil {
		var ret ShipCargo
		return ret
	}

	return o.Cargo
}

// GetCargoOk returns a tuple with the Cargo field value
// and a boolean to check if the value has been set.
func (o *ShipRefine200ResponseData) GetCargoOk() (*ShipCargo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cargo, true
}

// SetCargo sets field value
func (o *ShipRefine200ResponseData) SetCargo(v ShipCargo) {
	o.Cargo = v
}

// GetCooldown returns the Cooldown field value
func (o *ShipRefine200ResponseData) GetCooldown() Cooldown {
	if o == nil {
		var ret Cooldown
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *ShipRefine200ResponseData) GetCooldownOk() (*Cooldown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *ShipRefine200ResponseData) SetCooldown(v Cooldown) {
	o.Cooldown = v
}

// GetProduced returns the Produced field value
func (o *ShipRefine200ResponseData) GetProduced() []ShipRefine200ResponseDataProducedInner {
	if o == nil {
		var ret []ShipRefine200ResponseDataProducedInner
		return ret
	}

	return o.Produced
}

// GetProducedOk returns a tuple with the Produced field value
// and a boolean to check if the value has been set.
func (o *ShipRefine200ResponseData) GetProducedOk() ([]ShipRefine200ResponseDataProducedInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Produced, true
}

// SetProduced sets field value
func (o *ShipRefine200ResponseData) SetProduced(v []ShipRefine200ResponseDataProducedInner) {
	o.Produced = v
}

// GetConsumed returns the Consumed field value
func (o *ShipRefine200ResponseData) GetConsumed() []ShipRefine200ResponseDataProducedInner {
	if o == nil {
		var ret []ShipRefine200ResponseDataProducedInner
		return ret
	}

	return o.Consumed
}

// GetConsumedOk returns a tuple with the Consumed field value
// and a boolean to check if the value has been set.
func (o *ShipRefine200ResponseData) GetConsumedOk() ([]ShipRefine200ResponseDataProducedInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Consumed, true
}

// SetConsumed sets field value
func (o *ShipRefine200ResponseData) SetConsumed(v []ShipRefine200ResponseDataProducedInner) {
	o.Consumed = v
}

func (o ShipRefine200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipRefine200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cargo"] = o.Cargo
	toSerialize["cooldown"] = o.Cooldown
	toSerialize["produced"] = o.Produced
	toSerialize["consumed"] = o.Consumed
	return toSerialize, nil
}

type NullableShipRefine200ResponseData struct {
	value *ShipRefine200ResponseData
	isSet bool
}

func (v NullableShipRefine200ResponseData) Get() *ShipRefine200ResponseData {
	return v.value
}

func (v *NullableShipRefine200ResponseData) Set(val *ShipRefine200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableShipRefine200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableShipRefine200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipRefine200ResponseData(val *ShipRefine200ResponseData) *NullableShipRefine200ResponseData {
	return &NullableShipRefine200ResponseData{value: val, isSet: true}
}

func (v NullableShipRefine200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipRefine200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


