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

// checks if the TransferCargoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferCargoRequest{}

// TransferCargoRequest struct for TransferCargoRequest
type TransferCargoRequest struct {
	TradeSymbol string `json:"tradeSymbol"`
	Units int32 `json:"units"`
	ShipSymbol string `json:"shipSymbol"`
}

// NewTransferCargoRequest instantiates a new TransferCargoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferCargoRequest(tradeSymbol string, units int32, shipSymbol string) *TransferCargoRequest {
	this := TransferCargoRequest{}
	this.TradeSymbol = tradeSymbol
	this.Units = units
	this.ShipSymbol = shipSymbol
	return &this
}

// NewTransferCargoRequestWithDefaults instantiates a new TransferCargoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferCargoRequestWithDefaults() *TransferCargoRequest {
	this := TransferCargoRequest{}
	return &this
}

// GetTradeSymbol returns the TradeSymbol field value
func (o *TransferCargoRequest) GetTradeSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeSymbol
}

// GetTradeSymbolOk returns a tuple with the TradeSymbol field value
// and a boolean to check if the value has been set.
func (o *TransferCargoRequest) GetTradeSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeSymbol, true
}

// SetTradeSymbol sets field value
func (o *TransferCargoRequest) SetTradeSymbol(v string) {
	o.TradeSymbol = v
}

// GetUnits returns the Units field value
func (o *TransferCargoRequest) GetUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *TransferCargoRequest) GetUnitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *TransferCargoRequest) SetUnits(v int32) {
	o.Units = v
}

// GetShipSymbol returns the ShipSymbol field value
func (o *TransferCargoRequest) GetShipSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipSymbol
}

// GetShipSymbolOk returns a tuple with the ShipSymbol field value
// and a boolean to check if the value has been set.
func (o *TransferCargoRequest) GetShipSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipSymbol, true
}

// SetShipSymbol sets field value
func (o *TransferCargoRequest) SetShipSymbol(v string) {
	o.ShipSymbol = v
}

func (o TransferCargoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferCargoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tradeSymbol"] = o.TradeSymbol
	toSerialize["units"] = o.Units
	toSerialize["shipSymbol"] = o.ShipSymbol
	return toSerialize, nil
}

type NullableTransferCargoRequest struct {
	value *TransferCargoRequest
	isSet bool
}

func (v NullableTransferCargoRequest) Get() *TransferCargoRequest {
	return v.value
}

func (v *NullableTransferCargoRequest) Set(val *TransferCargoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferCargoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferCargoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferCargoRequest(val *TransferCargoRequest) *NullableTransferCargoRequest {
	return &NullableTransferCargoRequest{value: val, isSet: true}
}

func (v NullableTransferCargoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferCargoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


