/*
Public EMMA API

Testing OperatingSystemsAPIService

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

func Test_emma_OperatingSystemsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OperatingSystemsAPIService GetOperatingSystem", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var operatingSystemId int32

		resp, httpRes, err := apiClient.OperatingSystemsAPI.GetOperatingSystem(context.Background(), operatingSystemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OperatingSystemsAPIService GetOperatingSystems", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OperatingSystemsAPI.GetOperatingSystems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
