#!/usr/bin/env bash

set -euo pipefail

DT_SOURCE="NeoLoad"
DT_ANNOTATION_DESCRIPTION="Performance Test executed" 
DT_ANNOTATION_TYPE="Performance Test"
DT_START="2021-03-03T15:00:00Z"
DT_END="2021-03-03T15:10:00Z"
DT_EVENT_TYPE="CUSTOM_ANNOTATION"
DT_ENTITY_IDS="APPLICATION-EA7C4B59F27D43EB, HOST-D7C1DD84AFAD0DD1"
DT_ORIGINAL_CONFIGURATION="0%"
DT_CUSTOM_PROPERTIES='{"Virtual Users":"5", "Loops":"20"}'

docker run --rm -e DT_ENV_URL="${DYNATRACE_ENVIRONMENT_URL}" \
  -e SOURCE="${DT_SOURCE}" \
  -e START_TIME="${DT_START}" \
  -e END_TIME="${DT_END}" \
  -e ANNOTATIONDESCRIPTION="${DT_ANNOTATION_DESCRIPTION}" \
  -e ANNOTATIONTYPE="${DT_ANNOTATION_TYPE}" \
  -e EVENTTYPE="${DT_EVENT_TYPE}" \
  -e ENTITY_IDS="${DT_ENTITY_IDS}" \
  -e CUSTOM_PROPERTIES="${DT_CUSTOM_PROPERTIES}" \
  -e DT_API_TOKEN="${DYNATRACE_API_TOKEN}" dynatraceace/dynatrace-push-event:1.0.0