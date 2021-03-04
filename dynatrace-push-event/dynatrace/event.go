package dynatrace

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	dynatraceEnvironmentV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/environment/dynatrace"
	"github.com/kelseyhightower/envconfig"
)

// EventFields has no description
type EventFields struct {
	CustomProperties string `envconfig:"EVENT_CUSTOMPROPERTIES"`
	EntityIDs        string `envconfig:"EVENT_ENTITYIDS`
	TagRules         string `envconfig:"EVENT_TAGRULES`
}

func buildEvent() dynatraceEnvironmentV1.EventCreation {

	dtEvent := dynatraceEnvironmentV1.EventCreation{}

	err := envconfig.Process("EVENT", &dtEvent)
	if err != nil {
		log.Fatal(err.Error())
	}

	eventFields := EventFields{}

	err = envconfig.Process("EVENT", &eventFields)
	if err != nil {
		log.Fatal(err.Error())
	}

	if dtEvent.EventType != "" {
		dtEvent.SetEventType(dtEvent.EventType)
	}

	if dtEvent.Source != "" {
		dtEvent.SetSource(dtEvent.Source)
	}

	if dtEvent.DeploymentName != nil {
		dtEvent.SetDeploymentName(*dtEvent.DeploymentName)
	}

	if dtEvent.DeploymentVersion != nil {
		dtEvent.SetDeploymentVersion(*dtEvent.DeploymentVersion)
	}

	if eventFields.CustomProperties != "" {
		m := make(map[string]interface{})

		json.Unmarshal(([]byte)(eventFields.CustomProperties), &m)
		if err != nil {
			panic(err)
		}

		dtEvent.SetCustomProperties(m)
	}

	if eventFields.TagRules != "" || eventFields.EntityIDs != "" {

		dtAttachRules := dynatraceEnvironmentV1.PushEventAttachRules{}

		if eventFields.TagRules != "" {
			var tagRule []dynatraceEnvironmentV1.TagMatchRule

			json.Unmarshal(([]byte)(eventFields.TagRules), &tagRule)
			if err != nil {
				fmt.Println(err)
			}

			dtAttachRules.SetTagRule(tagRule)
		}

		if eventFields.EntityIDs != "" {
			entityIDsRaw := eventFields.EntityIDs
			entityIDs := strings.ReplaceAll(entityIDsRaw, " ", "")
			entityIDsArr := strings.Split(entityIDs, ",")
			dtAttachRules.SetEntityIds(entityIDsArr)
		}

		dtEvent.SetAttachRules(dtAttachRules)

	}

	return dtEvent
}

// PushEvent sends an event to Dynatrace
func PushEvent() {

	providerConf, err := initiateClient()
	dynatraceEnvClientV1 := providerConf.DynatraceEnvClientV1
	envConfigV1 := providerConf.AuthEnvV1

	dtEvent := buildEvent()

	resp, _, err := dynatraceEnvClientV1.EventsApi.PostEvent(envConfigV1).EventCreation(dtEvent).Execute()
	if err != nil {
		fmt.Println(getErrorMessage(err))
	}

	eventPushResponse, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nResponse from `EventsApi.PostEvent`: %s \n", eventPushResponse)
}
