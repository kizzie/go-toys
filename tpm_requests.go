package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "github.com/simplereach/timeutils"
  "strconv"
)

type Project struct {
  Id    int      `json:"id"`
  Name  string   `json:"name"`
}

type Password struct {
  Id                int             `json:"id"`
  Name              string          `json:"name"`
  Project           Project         `json:"project"`
  Notes_snippet     string          `json:"notes_snippet"`
  Tags              string          `json:"tags"`
  Username          string          `json:"username"`
  Email             string          `json:"email"`
  Expiry_date       timeutils.Time  `json:"expiry_date"`
  Expiry_status     int             `json:"expiry_status"`
  Archived          bool            `json:"archived"`
  Favourite         bool            `json:"favourite§"`
  Num_files         int             `json:"num_files"`
  Locked            bool            `json:"locked"`
  External_sharing  bool            `json:"external_sharing"`
  Updated_on        timeutils.Time  `json:"updated_on"`
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

func GetPassword(id int) Password {
  var output Password
  err := json.Unmarshal(tpm_request("passwords/" + strconv.Itoa(id) + ".json"), &output)

  if err != nil {
    panic(err)
  }

  return output
}
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
  fmt.Println(GetPassword(1))
}
