#!/usr/bin/env bash
#
# Dynatrace Maintenance Window
#
# This code supports the following API operations:
# - GET
# - POST
#
# Required enviroment variables
# - DYNATRACE_BASE_URL "https://abc123.live.dynatrace.com"
# - DYNATRACE_API_TOKEN
# - METHOD [ GET, POST ]
# - (if METHOD=POST) NAME "Dummy Maintenance Window"
# - (if METHOD=POST) DESCRIPTION "Maintenance window description..."
# - (if METHOD=POST) TYPE [ PLANNED, UNPLANNED ]
# - (if METHOD=POST) SUPPRESSION [ DETECT_PROBLEMS_AND_ALERT, DETECT_PROBLEMS_DONT_ALERT, DONT_DETECT_PROBLEMS ]
# - (if METHOD=POST) SCOPE
# - (if METHOD=POST) SCHEDULE

echo "----------------------------------------------------"
echo "Running container"


# Required parameters
DYNATRACE_BASE_URL=${DYNATRACE_BASE_URL:?'DYNATRACE_BASE_URL variable missing.'}
DYNATRACE_API_TOKEN=${DYNATRACE_API_TOKEN:?'DYNATRACE_API_TOKEN variable missing.'}
METHOD=${METHOD:?'METHOD variable is missing. Must be one of [GET, POST]'}

# Calculated values
DYNATRACE_API_URL="$DYNATRACE_BASE_URL/api/config/v1/maintenanceWindows"

if [[ "${DEBUG}" == "true" ]]; then
    ARGS+=( --verbose )
  fi

############################
# GET LOGIC
############################
if [[ "${METHOD}" == "GET" ]]; then
  ARGS=(
    --request GET
    --silent
    --url "${DYNATRACE_API_URL}"
    --header "Content-type: application/json"
    --header "Authorization: Api-Token ${DYNATRACE_API_TOKEN}"
  )
fi

############################
# POST LOGIC
############################

if [[ "${METHOD}" == "POST" ]]; then
  NAME=${NAME?'NAME variable missing.'}
  DESCRIPTION=${DESCRIPTION?'DESCRIPTION variable missing.'}
  TYPE=${TYPE?'TYPE variable missing. Must be one of [PLANNED, UNPLANNED]'}
  SUPPRESSION=${SUPPRESSION?'SUPPRESSION variable missing. Must be one of [ DETECT_PROBLEMS_AND_ALERT, DETECT_PROBLEMS_DONT_ALERT, DONT_DETECT_PROBLEMS ]'}
  SCOPE=${SCOPE?'SCOPE variable missing. Please provide a scope for this maintenance window.'}
  SCHEDULE=${SCHEDULE?'SCHEDULE variable missing.'}

  # build up POST_BODY variable
  POST_BODY="{\"name\":\"${NAME}\""
  POST_BODY+=", \"description\":\"${DESCRIPTION}\""
  POST_BODY+=", \"type\":\"${TYPE}\""
  POST_BODY+=", \"suppression\":\"${SUPPRESSION}\""
  POST_BODY+=", \"scope\": ${SCOPE}"
  POST_BODY+=", \"schedule\": ${SCHEDULE}"
  # add the closing bracket
  POST_BODY="${POST_BODY}}"

  ARGS=(
    --request POST
    --silent
    --url "${DYNATRACE_API_URL}"
    --header "Content-type: application/json"
    --header "Authorization: Api-Token ${DYNATRACE_API_TOKEN}"
    --data "${POST_BODY}"
  )
fi

echo "----------------------------------------------------"
echo "Calling Dynatrace Maintenance Window API"
echo "DYNATRACE_BASE_URL = ${DYNATRACE_BASE_URL}"
echo "METHOD = ${METHOD}"
if [[ "${METHOD}" == "POST" ]]; then
  echo "POST_BODY = $POST_BODY"
fi
echo "----------------------------------------------------"
status=$(curl "${ARGS[@]}")
echo "API call results = $status"
echo "----------------------------------------------------"
if [[ "$status" == *"error"* ]]; then
  echo "Error: Failed: $status"
  exit 1
else
  echo "Success: $status"
fi