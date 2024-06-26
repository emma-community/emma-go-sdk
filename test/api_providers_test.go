/*
Public EMMA API

Testing ProvidersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package emma

import (
	"context"
	openapiclient "github.com/emma-community/emma-go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_emma_ProvidersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProvidersAPIService GetProvider", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var providerId int32

		resp, httpRes, err := apiClient.ProvidersAPI.GetProvider(context.Background(), providerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProvidersAPIService GetProviders", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProvidersAPI.GetProviders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
