package main

import "fmt"
import "github.com/kelseyhightower/envconfig"
import "net/http"
import "os"
import "encoding/json"
import "io/ioutil"
import "bytes"

// Structure to hold input options from environment variables
type DTOptions struct {
  BaseURL string
  APIToken string
  HTTPMethod string
  Name string
  Description string
  Type string
  Suppression string
  Scope json.RawMessage
  Schedule json.RawMessage
  WindowID string
}

// Structure to model the JSON body to send to Dynatrace
type BodyContent struct {
  // Note the json:"name" part is the JSON key field.
  // eg. {"name": "foo"}
  // Without this it'd be {"Name": "foo"}
  
  Name string `json:"name"`
  Description string `json:"description"`
  Type string `json:"type"`
  Suppression string `json:"suppression"`
  Scope json.RawMessage `json:"scope"`
  Schedule json.RawMessage `json:"schedule"`
  WindowID string
}

func main() {

  // Parse input options
  var dtoptions DTOptions
  error := envconfig.Process("dt", &dtoptions)
  
  if (error != nil) {
    fmt.Println("got an error")
  }
  
  // These params are mandatory regardless of METHOD type used.
  if (dtoptions.BaseURL == "") {
    fmt.Println("NO DT_BASEURL environment variable defined.")
    os.Exit(1)
  }
  if (dtoptions.APIToken == "") {
    fmt.Println("NO DT_APITOKEN environment variable defined.")
    os.Exit(1)
  }
  if (dtoptions.HTTPMethod == "") {
    fmt.Println("NO DT_HTTPMETHOD environment variable defined. Must be one of [GET, POST, PUT]")
    os.Exit(1)
  }
  
  dynatrace_api_url := dtoptions.BaseURL + "/api/config/v1/maintenanceWindows"
  
  // Create HTTP Client
  client := &http.Client{}
  
  // Check additional mandatory params for POST and PUT
  // Exit if we're missing any
  if ((dtoptions.HTTPMethod == "POST" || dtoptions.HTTPMethod == "PUT") && dtoptions.Name == "") {
    fmt.Println("NO DT_NAME environment variable defined. Please name your maintenance window.")
    os.Exit(1)
  }
  if ((dtoptions.HTTPMethod == "POST" || dtoptions.HTTPMethod == "PUT") && dtoptions.Type == "") {
    fmt.Println("NO DT_TYPE environment variable defined. Must be one of [ PLANNED, UNPLANNED ].")
    os.Exit(1)
  }
  if ((dtoptions.HTTPMethod == "POST" || dtoptions.HTTPMethod == "PUT") && dtoptions.Suppression == "") {
    fmt.Println("NO DT_SUPPRESSION environment variable defined. Must be one of [ DETECT_PROBLEMS_AND_ALERT, DETECT_PROBLEMS_DONT_ALERT, DONT_DETECT_PROBLEMS ].")
    os.Exit(1)
  }
  // Schedule is an array. Check length is not zero    
  if ((dtoptions.HTTPMethod == "POST" || dtoptions.HTTPMethod == "PUT") && len(dtoptions.Schedule) == 0) {
    fmt.Println("NO DT_SCHEDULE environment variable defined.")
    os.Exit(1)
  }
    
  // Check additional mandatory parameter for PUT
  // Exit if we're missing it
  if (dtoptions.HTTPMethod == "PUT") {
    if (dtoptions.WindowID == "") {
      fmt.Println("NO DT_WINDOWID environment variable defined. Please provide a maintenance window ID in UUID format.")
      os.Exit(1)
    }
      
    // We have a window ID so adjust the URL
    dynatrace_api_url += "/" + dtoptions.WindowID
    
  }

  fmt.Println("==================")
  fmt.Println("Calling Dynatrace Maintenance Window API")
  fmt.Println("Method: " + dtoptions.HTTPMethod)
  fmt.Println("URL: " + dynatrace_api_url)
  
  bodyContent := new(BodyContent)
  bodyContent.Name = dtoptions.Name
  bodyContent.Description = dtoptions.Description
  bodyContent.Type = dtoptions.Type
  bodyContent.Suppression = dtoptions.Suppression
  bodyContent.Scope = dtoptions.Scope
  bodyContent.Schedule = dtoptions.Schedule
    
  jsonOutput,_ := json.Marshal(bodyContent)

  req,_ := http.NewRequest(dtoptions.HTTPMethod, dynatrace_api_url, bytes.NewBuffer([]byte(jsonOutput)))
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Api-Token " + dtoptions.APIToken)

  // Send Request
  resp,_ := client.Do(req)
  defer resp.Body.Close()
    
  fmt.Println("Response Code:", resp.StatusCode)
  
  body,_ := ioutil.ReadAll(resp.Body)
 
  fmt.Printf("Body: %s", body)
}

//TODO
//func sendRequest(Request req) {

//}