package commands

import(
  "github.com/hushtache/api"
  "github.com/hushtache/constants"
  "github.com/hushtache/utils"
  "fmt"
  "os"
  "strings"
)

func Remove() {

  // get the arguments
  args, opts := utils.ParseArgs()

  // check if name given
  if len(args[1:]) == 0 {

    fmt.Println("Keys to delete are required")

  } else {

    // the values
    propKey       := strings.ToLower(args[1])
    propTag       := []string{}

    // add the tags
    propTags      := strings.Split(fmt.Sprintf("%v", opts["Tag"]), ",")

    // check if already defined ?
    for _,tag := range propTags {

      // trim it
      if tag != "" {

        propTag = append(propTag, strings.TrimSpace(tag))

      }

    }

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

    // get the database list
    results, _       := api.Fetch(constants.GetDatabaseFile(), sig["key"].(string))

    // new results
    var newResults []interface{}

    // check if already defined ?
    for _,entry := range results {

      item := entry.(map[string]interface{})

      if fmt.Sprintf("%v", item["type"]) == "property" && 
            strings.ToLower(fmt.Sprintf("%v", item["name"])) == propKey {

      } else if fmt.Sprintf("%v", item["name"]) == "" || 
                  fmt.Sprintf("%v", item["name"]) == "nil" {

      } else {

        // results
        newResults    = append(newResults, item)

      }

    }

    api.Save(constants.GetDatabaseFile(), sig["key"].(string), newResults)

  }

}