package dynatrace

import (
	"context"
	"fmt"
	"log"
	"net/url"

	dynatraceEnvironmentV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/environment/dynatrace"
	"github.com/kelseyhightower/envconfig"
)

// Server details
type Server struct {
	DtEnvURL   string `envconfig:"DT_ENV_URL"`
	DtAPIToken string `envconfig:"DT_API_TOKEN"`
}

// ProviderConfiguration stores the Dynatrace API client
type ProviderConfiguration struct {
	DynatraceEnvClientV1 *dynatraceEnvironmentV1.APIClient
	AuthEnvV1            context.Context
}

func initiateClient() (*ProviderConfiguration, error) {

	var dynatraceConnection Server

	err := envconfig.Process("DT_ENV_URL", &dynatraceConnection)
	if err != nil {
		log.Fatal(err.Error())
	}

	parsedDTUrl, err := url.Parse(dynatraceConnection.DtEnvURL)
	if err != nil {
		fmt.Println(err)
	}

	authEnvV1 := context.WithValue(
		context.Background(),
		dynatraceEnvironmentV1.ContextAPIKeys,
		map[string]dynatraceEnvironmentV1.APIKey{
			"Api-Token": {
				Key:    dynatraceConnection.DtAPIToken,
				Prefix: "Api-Token",
			},
		},
	)

	authEnvV1 = context.WithValue(authEnvV1, dynatraceEnvironmentV1.ContextServerVariables, map[string]string{
		"name":     parsedDTUrl.Host,
		"protocol": parsedDTUrl.Scheme,
	})

	envV1 := dynatraceEnvironmentV1.NewConfiguration()
	dynatraceEnvClientV1 := dynatraceEnvironmentV1.NewAPIClient(envV1)

	return &ProviderConfiguration{
		DynatraceEnvClientV1: dynatraceEnvClientV1,
		AuthEnvV1:            authEnvV1,
	}, err

}

func getErrorMessage(err error) string {
	var errorMessage string

	if apiErr, ok := err.(dynatraceEnvironmentV1.GenericOpenAPIError); ok {
		return fmt.Sprintf("%v: %s", err, apiErr.Body())
	}

	if errURL, ok := err.(*url.Error); ok {
		return fmt.Sprintf("%v:%s", err, errURL)
	}

	return errorMessage

}
