package commands

import(
  "github.com/hushtache/api"
  "github.com/hushtache/constants"
  "fmt"
  "os"
)

func Users() {

  // get the key
  key           := api.GetCurrentKey()

  // get the sigs
  sigs, _       := api.FetchSignatures(constants.GetSignatureFile())

  // find the certain one
  sig, _        := api.FindSignature(sigs, key)

  // did we find it ?
  if sig == nil {

    fmt.Println("NOT ALLOWED !")

    os.Exit(1)
    
  }

  results, _       := api.Fetch(constants.GetDatabaseFile(), sig["key"].(string))

  for _,entry := range results {

    item := entry.(map[string]interface{})

    // check if defined
    if(item["type"].(string) == "user") {

      fmt.Println("-> " + item["label"].(string))

    }

  }

}