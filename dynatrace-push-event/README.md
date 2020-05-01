# Dynatrace Push Event

This container allows you to push the following event types to dynatrace:

- CUSTOM_ANNOTATION
- CUSTOM_DEPLOYMENT
- CUSTOM_CONFIGURATION
- CUSTOM_INFO

## Example

### Deployment Event

```bash
docker run --rm \
-e DYNATRACE_BASE_URL={your-environment-url} \
-e DYNATRACE_API_TOKEN={your-api-token} \
-e EVENT_TYPE=CUSTOM_DEPLOYMENT \
-e SOURCE=Jenkins \
-e DEPLOYMENT_NAME=k8s-deploy-staging \
-e DEPLOYMENT_VERSION=1.2.0 \
-e TAG_RULE="[{\"meTypes\":\"SERVICE\",\"tags\":[{\"context\":\"CONTEXTLESS\",\"key\":\"app\",\"value\":\"carts\"}]}]" \
-e CUSTOM_PROPERTIES="\"image\":\"carts\",\"tag\":\"0.1.1\"" \
dynatraceacm/push-event:0.1.0
```
