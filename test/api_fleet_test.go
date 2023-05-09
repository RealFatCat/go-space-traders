/*
SpaceTraders API

Testing FleetApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/realfatcat/go-space-traders"
)

func Test_sdk_FleetApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FleetApiService CreateChart", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.CreateChart(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService CreateShipShipScan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.CreateShipShipScan(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService CreateShipSystemScan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.CreateShipSystemScan(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService CreateShipWaypointScan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.CreateShipWaypointScan(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService CreateSurvey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.CreateSurvey(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService DockShip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.DockShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService ExtractResources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.ExtractResources(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService GetMyShip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.GetMyShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService GetMyShipCargo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.GetMyShipCargo(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService GetMyShips", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FleetApi.GetMyShips(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService GetShipCooldown", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.GetShipCooldown(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService GetShipNav", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.GetShipNav(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService Jettison", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.Jettison(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService JumpShip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.JumpShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService NavigateShip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.NavigateShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService OrbitShip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.OrbitShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService PatchShipNav", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.PatchShipNav(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService PurchaseCargo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.PurchaseCargo(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService PurchaseShip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FleetApi.PurchaseShip(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService RefuelShip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.RefuelShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService SellCargo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.SellCargo(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService ShipRefine", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.ShipRefine(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService TransferCargo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.TransferCargo(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetApiService WarpShip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipSymbol string

		resp, httpRes, err := apiClient.FleetApi.WarpShip(context.Background(), shipSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
