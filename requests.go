package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  "io/ioutil"

)

type ToDo struct {
  UserId int `json:userId`
  Id int `json:"id"`
  Title string `json:"title"`
  Completed bool `json:completed`
}

func main() {
  client := http.Client{}
  url := "https://jsonplaceholder.typicode.com/todos/1"

  req, _ := http.NewRequest("GET", url, nil)
  req.Header.Add("Accept", "application/json")
  resp, _ := client.Do(req)

  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)


  var output ToDo
  err := json.Unmarshal([]byte(body), &output)

  if err != nil {
    panic(err)
  }

  fmt.Println(output)
}
