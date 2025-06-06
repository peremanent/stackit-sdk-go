/*
STACKIT Model Serving API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package modelserving

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

func Test_modelserving_DefaultApiService(t *testing.T) {

	t.Run("Test DefaultApiService CreateToken", func(t *testing.T) {
		_apiUrlPath := "/v1/projects/{projectId}/regions/{regionId}/tokens"
		regionIdValue := "regionId"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"regionId"+"}", url.PathEscape(ParameterValueToString(regionIdValue, "regionId")), -1)
		projectIdValue := uuid.NewString()
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(_apiUrlPath, func(w http.ResponseWriter, req *http.Request) {
			data := CreateTokenResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for modelserving_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		regionId := regionIdValue
		projectId := projectIdValue

		resp, reqErr := apiClient.CreateToken(context.Background(), regionId, projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if IsNil(resp) {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DeleteToken", func(t *testing.T) {
		_apiUrlPath := "/v1/projects/{projectId}/regions/{regionId}/tokens/{tId}"
		regionIdValue := "regionId"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"regionId"+"}", url.PathEscape(ParameterValueToString(regionIdValue, "regionId")), -1)
		projectIdValue := uuid.NewString()
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		tIdValue := uuid.NewString()
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"tId"+"}", url.PathEscape(ParameterValueToString(tIdValue, "tId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(_apiUrlPath, func(w http.ResponseWriter, req *http.Request) {
			data := MessageResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for modelserving_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		regionId := regionIdValue
		projectId := projectIdValue
		tId := tIdValue

		resp, reqErr := apiClient.DeleteToken(context.Background(), regionId, projectId, tId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if IsNil(resp) {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetChatModel", func(t *testing.T) {
		_apiUrlPath := "/v1/regions/{regionId}/chat/models/{modelId}"
		regionIdValue := "regionId"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"regionId"+"}", url.PathEscape(ParameterValueToString(regionIdValue, "regionId")), -1)
		modelIdValue := uuid.NewString()
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"modelId"+"}", url.PathEscape(ParameterValueToString(modelIdValue, "modelId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(_apiUrlPath, func(w http.ResponseWriter, req *http.Request) {
			data := GetChatModelResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for modelserving_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		regionId := regionIdValue
		modelId := modelIdValue

		resp, reqErr := apiClient.GetChatModel(context.Background(), regionId, modelId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if IsNil(resp) {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetEmbeddingModel", func(t *testing.T) {
		_apiUrlPath := "/v1/regions/{regionId}/embedding/models/{modelId}"
		regionIdValue := "regionId"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"regionId"+"}", url.PathEscape(ParameterValueToString(regionIdValue, "regionId")), -1)
		modelIdValue := uuid.NewString()
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"modelId"+"}", url.PathEscape(ParameterValueToString(modelIdValue, "modelId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(_apiUrlPath, func(w http.ResponseWriter, req *http.Request) {
			data := GetEmbeddingsModelResp{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for modelserving_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		regionId := regionIdValue
		modelId := modelIdValue

		resp, reqErr := apiClient.GetEmbeddingModel(context.Background(), regionId, modelId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if IsNil(resp) {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetToken", func(t *testing.T) {
		_apiUrlPath := "/v1/projects/{projectId}/regions/{regionId}/tokens/{tId}"
		regionIdValue := "regionId"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"regionId"+"}", url.PathEscape(ParameterValueToString(regionIdValue, "regionId")), -1)
		projectIdValue := uuid.NewString()
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		tIdValue := uuid.NewString()
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"tId"+"}", url.PathEscape(ParameterValueToString(tIdValue, "tId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(_apiUrlPath, func(w http.ResponseWriter, req *http.Request) {
			data := GetTokenResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for modelserving_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		regionId := regionIdValue
		projectId := projectIdValue
		tId := tIdValue

		resp, reqErr := apiClient.GetToken(context.Background(), regionId, projectId, tId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if IsNil(resp) {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListModels", func(t *testing.T) {
		_apiUrlPath := "/v1/regions/{regionId}/models"
		regionIdValue := "regionId"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"regionId"+"}", url.PathEscape(ParameterValueToString(regionIdValue, "regionId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(_apiUrlPath, func(w http.ResponseWriter, req *http.Request) {
			data := ListModelsResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for modelserving_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		regionId := regionIdValue

		resp, reqErr := apiClient.ListModels(context.Background(), regionId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if IsNil(resp) {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListTokens", func(t *testing.T) {
		_apiUrlPath := "/v1/projects/{projectId}/regions/{regionId}/tokens"
		regionIdValue := "regionId"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"regionId"+"}", url.PathEscape(ParameterValueToString(regionIdValue, "regionId")), -1)
		projectIdValue := uuid.NewString()
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(_apiUrlPath, func(w http.ResponseWriter, req *http.Request) {
			data := ListTokenResp{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for modelserving_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		regionId := regionIdValue
		projectId := projectIdValue

		resp, reqErr := apiClient.ListTokens(context.Background(), regionId, projectId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if IsNil(resp) {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService PartialUpdateToken", func(t *testing.T) {
		_apiUrlPath := "/v1/projects/{projectId}/regions/{regionId}/tokens/{tId}"
		regionIdValue := "regionId"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"regionId"+"}", url.PathEscape(ParameterValueToString(regionIdValue, "regionId")), -1)
		projectIdValue := uuid.NewString()
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		tIdValue := uuid.NewString()
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"tId"+"}", url.PathEscape(ParameterValueToString(tIdValue, "tId")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(_apiUrlPath, func(w http.ResponseWriter, req *http.Request) {
			data := UpdateTokenResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for modelserving_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		regionId := regionIdValue
		projectId := projectIdValue
		tId := tIdValue

		resp, reqErr := apiClient.PartialUpdateToken(context.Background(), regionId, projectId, tId).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if IsNil(resp) {
			t.Fatalf("response not present")
		}
	})

}
