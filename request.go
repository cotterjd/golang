package main

import (
  "net/http"
  "io/ioutil"
)

func main () {
  response, getErr := http.Get("https://google.com")
  if getErr != nil {
    print(getErr)
  }
  defer response.Body.Close()

  result, readErr := ioutil.ReadAll(response.Body)
  if readErr != nil {
    print(readErr)
  }
  print(string(result))
}
