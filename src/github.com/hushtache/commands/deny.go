package commands

import(
  "github.com/hushtache/api"
  "github.com/hushtache/utils"
  "github.com/hushtache/constants"
  "fmt"
  "os"
)

func Deny() {

  // get the arguments
  args, _ := utils.ParseArgs()

  // check if name given
  if len(args[1:]) != 1 {

    fmt.Println("Name of user to remove is required")

  } else {

    // get the key
    key           := api.GetCurrentKey()

    // get the sigs
    sigs, _       := api.FetchSignatures(constants.GetSignatureFile())

    // check if we want to allow this ...
    if len(sigs) == 1 {

      // output
      fmt.Println("Cannot remove the only user able to access the secrets ...")

      // nope
      os.Exit(1)

    } 

    // find the certain one
    sig, _        := api.FindSignature(sigs, key)

    // did we find it ?
    if sig == nil {

      fmt.Println("NOT ALLOWED !")
      os.Exit(1)
      
    }

    // signature to remove
    userSignature := ""

    // fetch the details of the data store
    results, _       := api.Fetch(constants.GetDatabaseFile(), sig["key"].(string))

    // new results to save
    var newResults []interface{}

    for _,entry := range results {

      item := entry.(map[string]interface{})

      // check if defined
      if(fmt.Sprintf("%v", item["type"]) != "user") {

        // add to list
        newResults = append(newResults, item)

      } else if(fmt.Sprintf("%v", item["label"]) != args[1]) {

        // add to list
        newResults = append(newResults, item)

      } else {

        // set sig
        userSignature = fmt.Sprintf("%v", item["signature"])

      }

    }

    // append to the list of keys
    newSigs                  := []string{}

    // loop the sigs and create new list
    for _,sigLine := range sigs {

      // check if defined
      if(sigLine != userSignature) {

        // add to list
        newSigs = append(newSigs, sigLine)

      }

    }

    // save the signatures
    api.SaveSignatures(constants.GetSignatureFile(), newSigs)

    // save to "database"
    api.Save(constants.GetDatabaseFile(), sig["key"].(string), newResults)

  }

}