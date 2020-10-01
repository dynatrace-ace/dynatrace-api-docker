# Dynatrace Maintenance Window


# Prerequisites

```
go get "github.com/kelseyhightower/envconfig"
```

# GET Usage

Set your environment variables accordingly. For example:
```
set DT_BASEURL=https://abc123.live.dynatrace.com
set DT_APITOKEN=***
set DT_HTTPMETHOD=GET
```

Then execute:

```
bin/dt-maintenance-window.exe
or
go run app.go
```

# POST Usage
Set some additional environment variables:
```
set DT_NAME=Dummy Maintenance Window
set DT_DESCRIPTION=description here
set DT_TYPE=PLANNED
set DT_SUPPRESSION=DETECT_PROBLEMS_AND_ALERT
set DT_SCOPE={ "entities": [], "matches": [{ "type": "SERVICE", "tags": [{ "key": "customer-a", "context": "CONTEXTLESS" }] }, { "type": "APPLICATION", "tags": [{ "key": "customer-a", "context": "CONTEXTLESS" }] }] }
set DT_SCHEDULE={ "recurrenceType": "ONCE", "start": "2020-09-29 11:00", "end": "2020-09-29 11:15", "zoneId": "Australia/Brisbane" }
```

Then execute:
```
bin/dt-maintenance-window.exe
or
go run app.go
```

# PUT Usage
The `DT_WINDOWID` environment variable is mandatory for the PUT command
```
set DT_WINDOWID=12345678-1234-1234-1234-123456789012
```

Then execute:
```
bin/dt-maintenance-window.exe
or
go run app.go
```

# Build Notes

Build for linux:
```
set GOOS=linux
set GOOS=amd64
go build
```

Build for windows:
```
set GOOS=windows
set GOARCH=amd64
go build
```