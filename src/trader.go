package main

import (
  "net/http"
  "encoding/json"
  "log"
  "os"
//  "fmt"
)

type Config struct {
    ApiKey    string
}

func getConfig() Config {

// TODO: figure out how to have global logging var
  logging := true

  path, exists := os.LookupEnv("SF_CONFIG")
  if !exists {
    log.Fatal("No config path found in env SF_CONFIG")
  }
  
  if logging { log.Println("Found path " + path) }
  file, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  
  decoder := json.NewDecoder(file)
  config  := Config{}
  
  err     = decoder.Decode(&config)
  if err != nil {
    log.Fatal(err)
  }
  if logging { log.Println("key:" + config.ApiKey) }

  return config
}

func checkURL(urlFrag string) {
  resp, err := http.Get("https://api.stockfighter.io/ob/api/" + urlFrag)
  if err != nil || resp.StatusCode != 200 { 
    log.Fatal("Error: ", err, "\nResp: ", resp)
  }
  defer resp.Body.Close()
  log.Println(urlFrag + " UP: ", resp.Body)
}

func main() {
  config := getConfig()
  log.Println("key:" + config.ApiKey)
  //client := &http.Client{}
  
  checkURL("heartbeat")
  checkURL("venues/TESTEX/heartbeat")
  
}