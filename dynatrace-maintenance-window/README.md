# Dynatrace Maintenance Window Docker

# Build
```
docker build -t maintenance-window:0.1.0 .
```

# Usage

## GET All Maintenance Windows
```
docker run --rm ^
-e DYNATRACE_BASE_URL=https://abc123.live.dynatrace.com ^
-e DYNATRACE_API_TOKEN=*** ^
-e METHOD=GET ^
maintenance-window:0.1.0
```

## POST New Maintenance Window
This payload will tag `SERVICE` and `APPLICATION` entities which have a tag of `customer-a` (based on the scope).

```
docker run --rm ^
-e DYNATRACE_BASE_URL=https://abc123.live.dynatrace.com ^
-e DYNATRACE_API_TOKEN=*** ^
-e METHOD=POST ^
-e NAME="Dummy Maintenance Window" ^
-e DESCRIPTION="description here..." ^
-e TYPE="PLANNED" ^
-e SUPPRESSION="DETECT_PROBLEMS_AND_ALERT" ^
-e SCOPE=" { \"entities\": [ ], \"matches\": [{ \"type\": \"SERVICE\", \"tags\": [{ \"key\": \"customer-a\", \"context\": \"CONTEXTLESS\" }] }, { \"type\": \"APPLICATION\", \"tags\": [{ \"key\": \"customer-a\", \"context\": \"CONTEXTLESS\" }] }] } " ^
-e SCHEDULE=" { \"recurrenceType\": \"ONCE\", \"start\": \"2020-09-25 11:00\", \"end\": \"2020-09-25 11:15\", \"zoneId\": \"Australia/Brisbane\" } " ^
maintenance-window:0.1.0
```