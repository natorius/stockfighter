package main

import (
  //"net/http"
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
  
  // c  := &Config{ApiKey: "test"}
  // encoder := json.NewEncoder(os.Stdout)
  // encoder.Encode(c)
  
  // fmt.Println("enc test")
  //     enc := json.NewEncoder(os.Stdout)
  //   d := map[string]int{"apple": 5, "lettuce": 7}
  //   enc.Encode(d)
  
}


func main() {
  config := getConfig()
  log.Println("key:" + config.ApiKey)
}