package commands

import(
  "github.com/hushtache/api"
  "github.com/hushtache/constants"
  "fmt"
  "os/user"
)

func Init() {

  // current user
  usr, _ := user.Current()

  // check if we are not already configured
  if api.AlreadyConfigured( constants.GetDatabaseFile() ) == true {

    // stop right here
    fmt.Println("already contains an Hushtache store ... Run `hushtache allowed` to check if you are able to decrypt")

    // done 
    return

  }

  // get the key
  privateKeyPEM := api.GetCurrentKey()

  // generate the local datastore
  api.Init(privateKeyPEM, usr.Username)

}