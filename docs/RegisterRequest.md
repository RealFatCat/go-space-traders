# RegisterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Faction** | **string** | The faction you choose determines your headquarters. | 
**Symbol** | **string** | How other agents will see your ships and information. | 

## Methods

### NewRegisterRequest

`func NewRegisterRequest(faction string, symbol string, ) *RegisterRequest`

NewRegisterRequest instantiates a new RegisterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterRequestWithDefaults

`func NewRegisterRequestWithDefaults() *RegisterRequest`

NewRegisterRequestWithDefaults instantiates a new RegisterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFaction

`func (o *RegisterRequest) GetFaction() string`

GetFaction returns the Faction field if non-nil, zero value otherwise.

### GetFactionOk

`func (o *RegisterRequest) GetFactionOk() (*string, bool)`

GetFactionOk returns a tuple with the Faction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaction

`func (o *RegisterRequest) SetFaction(v string)`

SetFaction sets Faction field to given value.


### GetSymbol

`func (o *RegisterRequest) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *RegisterRequest) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *RegisterRequest) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


