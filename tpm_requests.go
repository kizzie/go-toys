package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "github.com/simplereach/timeutils"
  "strconv"
  "log"
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
  body, _ := tpm_request("passwords.json")
  err := json.Unmarshal(body, &output)

  if err != nil {
    panic(err)
  }

  return output
}

func GetPassword(id int) Password {
  var output Password
  body, err := tpm_request("passwords/" + strconv.Itoa(id) + ".json")
  if err != 200 {
    if err == 404 {
      log.Panic("Requested password not found " + strconv.Itoa(id))
    }
    panic("Error getting password: " + strconv.Itoa(err))
  }

  marshal_err := json.Unmarshal(body, &output)

  if marshal_err != nil {
    panic(marshal_err)
  }

  return output
}

// return the id of the saved password
// func SavePassword(password Password) int {
//   return 1
// }

func tpm_request(endpoint string) ([]byte, int) {
  base_url :="http://localhost/teampasswordmanager/index.php/api/v4/"
  client := http.Client{}

  req, err := http.NewRequest("GET", base_url + "/" + endpoint, nil)
  if err != nil {
    log.Print("Failed to create request " + base_url + "/" + endpoint)
    panic(err)
  }

  req.Header.Add("Content-Type", "application/json; charset=UTF-8")
  req.Header.Add("Authorization", "Basic a2F0OnBhc3N3b3Jk")
  resp, err := client.Do(req)

  if err != nil {
    log.Print("Failed to get " + base_url + "/" + endpoint)
    panic(err)
  }

  if resp.StatusCode != 200 {
    return nil, resp.StatusCode
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Panic("Error when converting to json")
    panic(err)
  }

  return body, 200
}

func main() {
  fmt.Println(GetPasswordList())
  fmt.Println(GetPassword(1))
  fmt.Println(GetPassword(2))
}
