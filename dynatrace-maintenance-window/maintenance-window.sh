#!/usr/bin/env bash
#
# Dynatrace Maintenance Window
#
# This code supports the following API operations:
# - GET
# - POST
# - PUT
#
# Required enviroment variables
# - DYNATRACE_BASE_URL "https://abc123.live.dynatrace.com"
# - DYNATRACE_API_TOKEN
# - METHOD [ GET, POST, PUT ]
# - (if METHOD=POST or PUT) NAME "Dummy Maintenance Window"
# - (if METHOD=POST or PUT) DESCRIPTION "Maintenance window description..."
# - (if METHOD=POST or PUT) TYPE [ PLANNED, UNPLANNED ]
# - (if METHOD=POST or PUT) SUPPRESSION [ DETECT_PROBLEMS_AND_ALERT, DETECT_PROBLEMS_DONT_ALERT, DONT_DETECT_PROBLEMS ]
# - (if METHOD=POST or PUT) SCOPE
# - (if METHOD=POST or PUT) SCHEDULE
# - (if METHOD=PUT) WINDOW_ID

echo "----------------------------------------------------"
echo "Running container"


# Required parameters
DYNATRACE_BASE_URL=${DYNATRACE_BASE_URL:?'DYNATRACE_BASE_URL variable missing.'}
DYNATRACE_API_TOKEN=${DYNATRACE_API_TOKEN:?'DYNATRACE_API_TOKEN variable missing.'}
METHOD=${METHOD:?'METHOD variable is missing. Must be one of [GET, POST, PUT]'}

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
    --request "${METHOD}"
    --silent
    --url "${DYNATRACE_API_URL}"
    --header "Content-type: application/json"
    --header "Authorization: Api-Token ${DYNATRACE_API_TOKEN}"
  )
fi

############################
# POST LOGIC
############################

if [[ "${METHOD}" == "POST" ]] || [[ "${METHOD}" == "PUT" ]]; then
  NAME=${NAME?'NAME variable missing.'}
  DESCRIPTION=${DESCRIPTION?'DESCRIPTION variable missing.'}
  TYPE=${TYPE?'TYPE variable missing. Must be one of [PLANNED, UNPLANNED]'}
  SUPPRESSION=${SUPPRESSION?'SUPPRESSION variable missing. Must be one of [ DETECT_PROBLEMS_AND_ALERT, DETECT_PROBLEMS_DONT_ALERT, DONT_DETECT_PROBLEMS ]'}
  SCOPE=${SCOPE?'SCOPE variable missing. Please provide a scope for this maintenance window.'}
  SCHEDULE=${SCHEDULE?'SCHEDULE variable missing.'}
  
  if [[ "${METHOD}" == "PUT" ]]; then
    WINDOW_ID=${WINDOW_ID?'WINDOW_ID variable missing. You must provide a maintenance window ID.'}
    DYNATRACE_API_URL+="/${WINDOW_ID}"
  fi

  # build up BODY variable
  BODY="{\"name\":\"${NAME}\""
  BODY+=", \"description\":\"${DESCRIPTION}\""
  BODY+=", \"type\":\"${TYPE}\""
  BODY+=", \"suppression\":\"${SUPPRESSION}\""
  BODY+=", \"scope\": ${SCOPE}"
  BODY+=", \"schedule\": ${SCHEDULE}"
  # add the closing bracket
  BODY="${BODY}}"

  ARGS=(
    --request "${METHOD}"
    --silent
    --url "${DYNATRACE_API_URL}"
    --header "Content-type: application/json"
    --header "Authorization: Api-Token ${DYNATRACE_API_TOKEN}"
    --data "${BODY}"
  )
fi

echo "----------------------------------------------------"
echo "Calling Dynatrace Maintenance Window API"
echo "DYNATRACE_API_URL = ${DYNATRACE_API_URL}"
echo "METHOD = ${METHOD}"
if [[ "${METHOD}" == "POST" ]] || [[ "${METHOD}" == "PUT" ]]; then
  echo "BODY = $BODY"
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