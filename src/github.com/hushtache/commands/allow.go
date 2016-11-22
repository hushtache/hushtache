package commands

import(
  "github.com/hushtache/api"
  "github.com/hushtache/utils"
  "github.com/hushtache/constants"
  "github.com/hushtache/gzip"
  "github.com/nu7hatch/gouuid"
  "encoding/json"
  "fmt"
  "os"
)

func Allow() {

  // get the arguments
  args, _ := utils.ParseArgs()

  // check if name given
  if len(args[1:]) != 2 {

    fmt.Println("Name and key of new user to allow is required")

  } else {

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

    // generate uuid
    u, _ := uuid.NewV4()

    // resulting item
    sigItem               := map[string]interface{}{}
    sigItem["uid"]        = u.String()
    sigItem["label"]      = args[1]
    sigItem["key"]        = sig["key"].(string)
    enc, _                := json.Marshal(sigItem)
  
    inflated, _           := gzip.Inflate(args[2])

    // get it
    signedPassword        := api.EncryptSignature(inflated, string(enc))

    // fetch the details of the data store
    results, _       := api.Fetch(constants.GetDatabaseFile(), sig["key"].(string))

    for _,entry := range results {

      item := entry.(map[string]interface{})

      // check if defined
      if(fmt.Sprintf("%v", item["type"]) == "user") {

        // check if defined
        if(fmt.Sprintf("%v", item["label"]) == args[1]) {

          // check if not already defined
          fmt.Println("User with that name already registered ...")

          // exit
          os.Exit(1)

        }

      }

    }

    // append to the list of keys
    sigs                  = append(sigs, signedPassword)

    // add to the database
    // resulting item
    storeItem               := map[string]interface{}{}
    storeItem["type"]       = "user"
    storeItem["signature"]  = signedPassword
    storeItem["label"]      = args[1]
    results                 = append(results, storeItem)

    // save the signatures
    api.SaveSignatures(constants.GetSignatureFile(), sigs)

    // save to "database"
    api.Save(constants.GetDatabaseFile(), sig["key"].(string), results)

  }

}