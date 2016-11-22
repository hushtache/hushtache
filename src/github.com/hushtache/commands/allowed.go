package commands

import(
  "github.com/hushtache/api"
  "github.com/hushtache/constants"
  "fmt"
)

func Allowed() {

  // get the key
  key           := api.GetCurrentKey()

  // get the sigs
  sigs, _     := api.FetchSignatures(constants.GetSignatureFile())

  // find the certain one
  sig, _      := api.FindSignature(sigs, key)

  // did we find it ?
  if sig != nil {

    fmt.Println("Allowed ! Registered as " + sig["label"].(string))
    
  } else {

    fmt.Println("NOT ALLOWED !")

  }

}