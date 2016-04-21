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

type Stock struct {
  Venue  string `json:"venue"`
  Symbol string `json:"symbol"`
}

type Order struct {
  Price int  `json:"price"`
  Qty   int  `json:"qty"`
  IsBuy bool `json:"isBuy"`
}
type OrderBook struct {
  Stock
  Bids []Order `json:"bids"`
  Asks []Order `json:"asks"`
  Time string  `json:"ts"`
}
type OrderBookResponse struct {
  Response
  OrderBook
}
type NoInterface interface {} // TODO: nointerface is probly bad

// TODO: url and body expectation are inseperable, should be turned
//       into a struct
func (c Client) getURL(urlFrag string, body NoInterface) {
  url      := "https://api.stockfighter.io/ob/api/" + urlFrag
  req, err := http.NewRequest("GET", url, nil)
  if err != nil { log.Fatal(err) }
  req.Header.Add("X-Starfighter-Authorization", c.ApiKey)

  resp, err := c.Client.Do(req)
  if err != nil || resp.StatusCode != 200 { 
    log.Fatal("Error: ", err, "\nResp: ", resp)
  }
  defer resp.Body.Close()

  decoder := json.NewDecoder(resp.Body)
  err      = decoder.Decode(body)
  if err != nil { log.Fatal(err) }

  log.Printf("%v UP: %+v", urlFrag, body)
}

func main() {
  config := getConfig()
  log.Println("key:" + config.ApiKey)
  // TODO: rename Client, AuthClient?
  client := Client{&http.Client{}, config.ApiKey}
  
  var resp Response
  client.getURL("heartbeat", &resp)
  client.getURL("venues/TESTEX/heartbeat", &resp)

  var orderBookResponse OrderBookResponse
  client.getURL("venues/TESTEX/stocks/FOOBAR", &orderBookResponse)
  
  
}