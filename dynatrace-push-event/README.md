# Dynatrace Push Event

This container allows you to push the following event types to dynatrace:

- CUSTOM_ANNOTATION
- CUSTOM_DEPLOYMENT
- CUSTOM_CONFIGURATION
- CUSTOM_INFO

## Example

### Deployment event

```bash
docker run --rm \
-e DYNATRACE_BASE_URL={your-environment-url} \
-e DYNATRACE_API_TOKEN={your-api-token} \
-e EVENT_TYPE=CUSTOM_DEPLOYMENT \
-e SOURCE=Jenkins \
-e DEPLOYMENT_NAME=k8s-deploy-staging \
-e DEPLOYMENT_VERSION=1.2.0 \
-e DEPLOYMENT_PROJECT=sockshop \
-e DEPLOYMENT_REMEDIATION_ACTION=https://10.42.241.27/api/v2/job_templates/36/launch \
-e CI_BACK_LINK=http://10.51.241.100/job/sockshop/job/orders/job/master/1/ \
-e TAG_RULE="[{\"meTypes\":\"SERVICE\",\"tags\":[{\"context\":\"CONTEXTLESS\",\"key\":\"app\",\"value\":\"carts\"},{\"context\":\"CONTEXTLESS\",\"key\":\"env\",\"value\":\"dev\"}]}]" \
-e CUSTOM_PROPERTIES="\"BuildNumber\":\"1\",\"GitCommit\":\"6c2bda106a0e15880c0a8a2755f2c112ac2b3698\"" \
dynatraceacm/push-event:0.1.0
```

### Configuration change event

```bash
docker run --rm \
-e DYNATRACE_BASE_URL={your-environment-url} \
-e DYNATRACE_API_TOKEN={your-api-token} \
-e EVENT_TYPE=CUSTOM_CONFIGURATION \
-e SOURCE="Ansible Tower" \
-e DESCRIPTION="promotion rate set to 20 %" \
-e CONFIGURATION="promotion rate" \
-e ORIGINAL_CONFIGURATION="0%" \
-e TAG_RULE="[{\"meTypes\":\"SERVICE\",\"tags\":[{\"context\":\"CONTEXTLESS\",\"key\":\"app\",\"value\":\"carts\"},{\"context\":\"CONTEXTLESS\",\"key\":\"env\",\"value\":\"dev\"}]}]" \
-e CUSTOM_PROPERTIES="\"remediationAction\":\"https://10.42.241.27/api/v2/job_templates/36/launch\"" \
dynatraceacm/push-event:0.1.0
```

### Annotation Event for a Performance Test

```bash
docker run --rm \
-e DYNATRACE_BASE_URL={your-environment-url} \
-e DYNATRACE_API_TOKEN={your-api-token} \
-e EVENT_TYPE="CUSTOM_ANNOTATION" \
-e SOURCE="NeoLoad" \
-e ANNOTATION_TYPE="Performance Test" \
-e ANNOTATION_DESCRIPTION="Performance Test executed" \
-e START="2020-05-03T15:00:00Z" \
-e END="2020-05-03T15:10:00Z" \
-e TAG_RULE="[{\"meTypes\":\"SERVICE\",\"tags\":[{\"context\":\"CONTEXTLESS\",\"key\":\"iis-app\",\"value\":\"Default Web Site:80\"},{\"context\":\"CONTEXTLESS\",\"key\":\"env\",\"value\":\"dev\"}]}]" \
dynatraceacm/push-event:0.1.0
```
