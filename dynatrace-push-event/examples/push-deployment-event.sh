#!/usr/bin/env bash

set -euo pipefail

DT_SOURCE="Jenkins"
DT_DEPLOYMENT_NAME="k8s-staging"
DT_DEPLOYMENT_PROJECT="Sockshop"
DT_DEPLOYMENT_VERSION="1.0.0"
DT_EVENT_TYPE="CUSTOM_DEPLOYMENT"
DT_CUSTOM_PROPERTIES='{"Stage":"dev"}'
DT_TAG_MATCH_RULES='[{"meTypes":["HOST"],"tags":[{"context":"CONTEXTLESS","key":"hostenv","value":"prod"},{"context":"CONTEXTLESS","key":"hostapp","value":"carts"}]}]'

docker run --rm -e DT_ENV_URL="${DYNATRACE_ENVIRONMENT_URL}" \
  -e SOURCE="${DT_SOURCE}" \
  -e DEPLOYMENTNAME="${DT_DEPLOYMENT_NAME}" \
  -e DEPLOYMENTVERSION="${DT_DEPLOYMENT_VERSION}" \
  -e EVENTTYPE="${DT_EVENT_TYPE}" \
  -e DEPLOYMENTPROJECT="${DT_DEPLOYMENT_PROJECT}" \
  -e CUSTOM_PROPERTIES="${DT_CUSTOM_PROPERTIES}" \
  -e TAG_MATCH_RULES="${DT_TAG_MATCH_RULES}" \
  -e DT_API_TOKEN="${DYNATRACE_API_TOKEN}" dynatraceace/dynatrace-push-event:1.0.0