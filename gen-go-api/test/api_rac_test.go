/*
authentik

Testing RacApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_api_RacApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RacApiService RacConnectionTokensDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var connectionTokenUuid string

		httpRes, err := apiClient.RacApi.RacConnectionTokensDestroy(context.Background(), connectionTokenUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RacApiService RacConnectionTokensList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RacApi.RacConnectionTokensList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RacApiService RacConnectionTokensPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var connectionTokenUuid string

		resp, httpRes, err := apiClient.RacApi.RacConnectionTokensPartialUpdate(context.Background(), connectionTokenUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RacApiService RacConnectionTokensRetrieve", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var connectionTokenUuid string

		resp, httpRes, err := apiClient.RacApi.RacConnectionTokensRetrieve(context.Background(), connectionTokenUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RacApiService RacConnectionTokensUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var connectionTokenUuid string

		resp, httpRes, err := apiClient.RacApi.RacConnectionTokensUpdate(context.Background(), connectionTokenUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RacApiService RacConnectionTokensUsedByList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var connectionTokenUuid string

		resp, httpRes, err := apiClient.RacApi.RacConnectionTokensUsedByList(context.Background(), connectionTokenUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RacApiService RacEndpointsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RacApi.RacEndpointsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RacApiService RacEndpointsDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pbmUuid string

		httpRes, err := apiClient.RacApi.RacEndpointsDestroy(context.Background(), pbmUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RacApiService RacEndpointsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RacApi.RacEndpointsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RacApiService RacEndpointsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pbmUuid string

		resp, httpRes, err := apiClient.RacApi.RacEndpointsPartialUpdate(context.Background(), pbmUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RacApiService RacEndpointsRetrieve", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pbmUuid string

		resp, httpRes, err := apiClient.RacApi.RacEndpointsRetrieve(context.Background(), pbmUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RacApiService RacEndpointsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pbmUuid string

		resp, httpRes, err := apiClient.RacApi.RacEndpointsUpdate(context.Background(), pbmUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RacApiService RacEndpointsUsedByList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pbmUuid string

		resp, httpRes, err := apiClient.RacApi.RacEndpointsUsedByList(context.Background(), pbmUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
