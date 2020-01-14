package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "github.com/simplereach/timeutils"
)

type Project struct {
  Id    int      `json:"id"`
  Name  string   `json:"name"`
}

type Password struct {
  Id                int
  Name              string
  Project           Project
  Notes_snippet     string
  Tags              string
  Username          string
  Email             string
  Expiry_date       timeutils.Time
  Expiry_status     int
  Archived          bool
  Favourite         bool
  Num_files         int
  Locked            bool
  External_sharing  bool
  Updated_on        timeutils.Time
}

type PasswordList []Password

func GetPasswordList() PasswordList {
  var output PasswordList
  err := json.Unmarshal(tpm_request("passwords.json"), &output)

  if err != nil {
    panic(err)
  }

  return output
}
//
// func GetPassword(id int) Password {
//
// }
//
// func SavePassword(password Password) {
//
// }

func tpm_request(endpoint string) []byte {
  base_url :="http://localhost/teampasswordmanager/index.php/api/v4/"
  client := http.Client{}

  req, _ := http.NewRequest("GET", base_url + "/" + endpoint, nil)
  req.Header.Add("Content-Type", "application/json; charset=UTF-8")
  req.Header.Add("Authorization", "Basic a2F0OnBhc3N3b3Jk")
  resp, _ := client.Do(req)

  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)

  return body
}

func main() {
  fmt.Println(GetPasswordList())
}
