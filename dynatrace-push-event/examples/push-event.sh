#!/usr/bin/env bash

set -euo pipefail

DT_SOURCE=Jenkins
DT_DEPLOYMENT_NAME=k8s-staging
DT_DEPLOYMENT_VERSION=1.0.0
DT_EVENT_TYPE=CUSTOM_DEPLOYMENT
DT_ENTITY_IDS=APPLICATION-EA7C4B59F27D43EB
DT_CUSTOM_PROPERTIES='{"Project": "sockshop","Stage": "dev"}'
DT_TAGS='[{"meTypes":["HOST"],"tags":[{"context":"CONTEXTLESS","key":"hostenv","value":"prod"},{"context":"CONTEXTLESS","key":"hostapp","value":"carts"}]}]'

#docker build -t dynatraceace/dynatrace-push-event:1.0.0 ../
#-e EVENT_ENTITYIDS=$DT_ENTITY_IDS \

docker run --rm -e DT_ENV_URL="${DYNATRACE_ENVIRONMENT_URL}" \
  -e EVENT_SOURCE="${DT_SOURCE}" \
  -e EVENT_DEPLOYMENTNAME="${DT_DEPLOYMENT_NAME}" \
  -e EVENT_DEPLOYMENTVERSION="${DT_DEPLOYMENT_VERSION}" \
  -e EVENT_EVENTTYPE="${DT_EVENT_TYPE}" \
  -e EVENT_CUSTOMPROPERTIES="${DT_CUSTOM_PROPERTIES}" \
  -e EVENT_TAGRULES="${DT_TAGS}" \
  -e DT_API_TOKEN="${DYNATRACE_API_TOKEN}" luisescobar/dynatrace-push-event:1.0.0