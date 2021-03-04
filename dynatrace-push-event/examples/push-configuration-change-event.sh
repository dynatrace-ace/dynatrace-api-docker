#!/usr/bin/env bash

set -euo pipefail

DT_SOURCE='Ansible Tower'
DT_DESCRIPTION="promotion rate set to 20 %"
DT_CONFIGURATION="promotion rate"
DT_EVENT_TYPE="CUSTOM_CONFIGURATION"
DT_ORIGINAL_CONFIGURATION="0%"
DT_CUSTOM_PROPERTIES='{"remediationAction":"https://10.42.241.27/api/v2/job_templates/36/launch"}'
DT_TAG_MATCH_RULES='[{"meTypes":["HOST"],"tags":[{"context":"CONTEXTLESS","key":"hostenv","value":"prod"},{"context":"CONTEXTLESS","key":"hostapp","value":"carts"}]}]'

docker run --rm -e DT_ENV_URL="${DYNATRACE_ENVIRONMENT_URL}" \
  -e SOURCE="${DT_SOURCE}" \
  -e DESCRIPTION="${DT_DESCRIPTION}" \
  -e CONFIGURATION="${DT_CONFIGURATION}" \
  -e ORIGINAL="${DT_ORIGINAL_CONFIGURATION}" \
  -e EVENTTYPE="${DT_EVENT_TYPE}" \
  -e CUSTOM_PROPERTIES="${DT_CUSTOM_PROPERTIES}" \
  -e TAG_MATCH_RULES="${DT_TAG_MATCH_RULES}" \
  -e DT_API_TOKEN="${DYNATRACE_API_TOKEN}" dynatraceace/dynatrace-push-event:1.0.0