package dynatrace

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	dynatraceEnvironmentV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/environment/dynatrace"
	"github.com/kelseyhightower/envconfig"
)

// EventFields used to get environment variable field for non string types
type EventFields struct {
	CustomProperties string `envconfig:"CUSTOM_PROPERTIES"`
	TimeseriesIDs    string `envconfig:"TIMESERIES_IDS"`
	StartTime        string `envconfig:"START_TIME"`
	EndTime          string `envconfig:"END_TIME"`
	TimeoutMinutes   string `envconfig:"TIME_OUT_MINUTES"`
	EntityIDs        string `envconfig:"ENTITY_IDS"`
	TagMatchRules    string `envconfig:"TAG_MATCH_RULES"`
	AllowDavisMerge  string `envconfig:"ALLOW_DAVIS_MERGE"`
}

func buildEvent() dynatraceEnvironmentV1.EventCreation {

	dtEvent := dynatraceEnvironmentV1.EventCreation{}

	err := envconfig.Process("", &dtEvent)
	if err != nil {
		log.Fatal(err.Error())
	}

	eventFields := EventFields{}

	err = envconfig.Process("", &eventFields)
	if err != nil {
		log.Fatal(err.Error())
	}

	if dtEvent.EventType != "" {
		dtEvent.SetEventType(dtEvent.EventType)
	}

	if eventFields.StartTime != "" {
		startTime, err := time.Parse(time.RFC3339, eventFields.StartTime)
		if err != nil {
			log.Fatal(err.Error())
		}
		startTimeEpoch := startTime.UnixNano() / int64(time.Millisecond)
		dtEvent.SetStart(startTimeEpoch)
	}

	if eventFields.EndTime != "" {
		endTime, err := time.Parse(time.RFC3339, eventFields.EndTime)
		if err != nil {
			log.Fatal(err.Error())
		}
		endTimeEpoch := endTime.UnixNano() / int64(time.Millisecond)
		dtEvent.SetEnd(endTimeEpoch)
	}

	if eventFields.TimeoutMinutes != "" {
		timeoutMinutes, err := strconv.ParseInt(eventFields.TimeoutMinutes, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		dtEvent.SetTimeoutMinutes(int32(timeoutMinutes))
	}

	if dtEvent.Source != "" {
		dtEvent.SetSource(dtEvent.Source)
	}

	if dtEvent.AnnotationType != nil {
		dtEvent.SetAnnotationType(*dtEvent.AnnotationType)
	}

	if dtEvent.AnnotationDescription != nil {
		dtEvent.SetAnnotationDescription(*dtEvent.AnnotationDescription)
	}

	if dtEvent.Description != nil {
		dtEvent.SetDescription(*dtEvent.Description)
	}

	if dtEvent.DeploymentName != nil {
		dtEvent.SetDeploymentName(*dtEvent.DeploymentName)
	}

	if dtEvent.DeploymentVersion != nil {
		dtEvent.SetDeploymentVersion(*dtEvent.DeploymentVersion)
	}

	if eventFields.TimeseriesIDs != "" {
		timeseriesIDsRaw := eventFields.TimeseriesIDs
		timeseriesIDs := strings.ReplaceAll(timeseriesIDsRaw, " ", "")
		timeseriesIDsArr := strings.Split(timeseriesIDs, ",")
		dtEvent.SetTimeseriesIds(timeseriesIDsArr)
	}

	if dtEvent.DeploymentProject != nil {
		dtEvent.SetDeploymentProject(*dtEvent.DeploymentProject)
	}

	if dtEvent.CiBackLink != nil {
		dtEvent.SetCiBackLink(*dtEvent.CiBackLink)
	}

	if dtEvent.RemediationAction != nil {
		dtEvent.SetRemediationAction(*dtEvent.RemediationAction)
	}

	if dtEvent.Original != nil {
		dtEvent.SetOriginal(*dtEvent.Original)
	}

	if dtEvent.Changed != nil {
		dtEvent.SetChanged(*dtEvent.Changed)
	}

	if dtEvent.Configuration != nil {
		dtEvent.SetConfiguration(*dtEvent.Configuration)
	}

	if dtEvent.Title != nil {
		dtEvent.SetTitle(*dtEvent.Title)
	}

	if eventFields.AllowDavisMerge != "" {
		allowDavisMerge, err := strconv.ParseBool(eventFields.AllowDavisMerge)
		if err != nil {
			log.Fatal(err.Error())
		}
		dtEvent.SetAllowDavisMerge(allowDavisMerge)
	}

	if eventFields.CustomProperties != "" {
		m := make(map[string]interface{})

		json.Unmarshal(([]byte)(eventFields.CustomProperties), &m)
		if err != nil {
			panic(err)
		}

		dtEvent.SetCustomProperties(m)
	}

	if eventFields.TagMatchRules != "" || eventFields.EntityIDs != "" {

		dtAttachRules := dynatraceEnvironmentV1.PushEventAttachRules{}

		if eventFields.TagMatchRules != "" {
			var tagRule []dynatraceEnvironmentV1.TagMatchRule

			json.Unmarshal(([]byte)(eventFields.TagMatchRules), &tagRule)
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

	eventPushResponse, err := json.Marshal(dtEvent)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nEvent JSON body: %s \n", eventPushResponse)

	resp, _, err := dynatraceEnvClientV1.EventsApi.PostEvent(envConfigV1).EventCreation(dtEvent).Execute()
	if err != nil {
		fmt.Println(getErrorMessage(err))
	}

	eventPushResponse, err = json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nResponse from `EventsApi.PostEvent`: %s \n", eventPushResponse)
}
