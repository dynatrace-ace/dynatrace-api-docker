# Dynatrace Maintenance Window Docker

# Build
```
docker build -t maintenance-window:0.1.0 .
```

# Usage

Either pass variables via command line `-e` parameters or via an environment variable file.

You'll need to escape the JSON if you use `-e` parameters for `SCOPE` and `SCHEDULE`.

For simplicity, use of the `--env-file` is recommended.

## GET All Maintenance Windows

Create your environment variable file. For example `env.txt`:

```
DYNATRACE_BASE_URL=https://abc123.live.dynatrace.com
DYNATRACE_API_TOKEN=***
METHOD=GET
```

Then execute with:
```
docker run --rm --env-file=env.txt maintenance-window:0.1.0
```

## POST New Maintenance Window
This payload will tag `SERVICE` and `APPLICATION` entities which have a tag of `customer-a` (based on the scope).

Modify the environment variable file (`env.txt`) to add additional values:

```
DYNATRACE_BASE_URL=https://abc123.live.dynatrace.com
DYNATRACE_API_TOKEN=***
METHOD=POST
NAME=Dummy Maintenance Window
DESCRIPTION=description here
TYPE=PLANNED
SUPPRESSION=DETECT_PROBLEMS_AND_ALERT
SCOPE={ "entities": [], "matches": [{ "type": "SERVICE", "tags": [{ "key": "customer-a", "context": "CONTEXTLESS" }] }, { "type": "APPLICATION", "tags": [{ "key": "customer-a", "context": "CONTEXTLESS" }] }] }
SCHEDULE={ "recurrenceType": "ONCE", "start": "2020-09-29 11:00", "end": "2020-09-29 11:15", "zoneId": "Australia/Brisbane" }
```

Then execute:
```
docker run --rm --env-file=env.txt maintenance-window:0.1.0
```

## PUT Maintenance Window
Use PUT to update an existing window or create a new one, if one doesn't already exist.

Requires a new `WINDOW_ID` parameter. Add this to the `env.txt` file:

```
DYNATRACE_BASE_URL=https://abc123.live.dynatrace.com
DYNATRACE_API_TOKEN=***
METHOD=POST
NAME=Dummy Maintenance Window
DESCRIPTION=description here
TYPE=PLANNED
SUPPRESSION=DETECT_PROBLEMS_AND_ALERT
SCOPE={ "entities": [], "matches": [{ "type": "SERVICE", "tags": [{ "key": "customer-a", "context": "CONTEXTLESS" }] }, { "type": "APPLICATION", "tags": [{ "key": "customer-a", "context": "CONTEXTLESS" }] }] }
SCHEDULE={ "recurrenceType": "ONCE", "start": "2020-09-29 11:00", "end": "2020-09-29 11:15", "zoneId": "Australia/Brisbane" }
WINDOW_ID=12345678-1234-1234-1234-123456789012
```

Then run, as usual, with:
```
docker run --rm --env-file=env.txt maintenance-window:0.1.0
```