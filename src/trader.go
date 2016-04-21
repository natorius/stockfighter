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

// TODO: extract json decoding to function  
  decoder := json.NewDecoder(file)
  config  := Config{}
  
  err     = decoder.Decode(&config)
  if err != nil {
    log.Fatal(err)
  }
  if logging { log.Println("key:" + config.ApiKey) }

  return config
}

type Client struct {
  Client *http.Client
  ApiKey string
}

type Response struct {
  Ok bool `json:"ok"`
}

func (c Client) checkURL(urlFrag string) {
  url      := "https://api.stockfighter.io/ob/api/" + urlFrag
  req, err := http.NewRequest("GET", url, nil)
  if err != nil { log.Fatal(err) }
  req.Header.Add("X-Starfighter-Authorization", c.ApiKey)

  resp, err := c.Client.Do(req)
  if err != nil || resp.StatusCode != 200 { 
    log.Fatal("Error: ", err, "\nResp: ", resp)
  }
  defer resp.Body.Close()

  var sfresp Response
  decoder := json.NewDecoder(resp.Body)

  err      = decoder.Decode(&sfresp)
  if err != nil { log.Fatal(err) }

  log.Printf("%v UP: %+v", urlFrag, sfresp)
}

func main() {
  config := getConfig()
  log.Println("key:" + config.ApiKey)
  client := Client{&http.Client{}, config.ApiKey}
  
  client.checkURL("heartbeat")
  client.checkURL("venues/TESTEX/heartbeat")
//  client.checkURL("venues/UKHNEX/heartbeat")
  
  
}