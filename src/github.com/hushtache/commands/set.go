package commands

import(
  "github.com/hushtache/api"
  "github.com/hushtache/constants"
  "github.com/hushtache/utils"
  "fmt"
  "os"
  "time"
  "strings"
)

func Set() {

  // get the arguments
  args, opts := utils.ParseArgs()

  // check if name given
  if len(args[1:]) < 2 {

    fmt.Println("Name and key of new user to allow is required")

  } else {

    // the values
    propKey       := strings.ToLower(args[1])
    propValue     := strings.Join(args[2:], " ")
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

    // tag if we found it
    foundFlag := false

    // check if already defined ?
    for _,entry := range results {

      item := entry.(map[string]interface{})

      if fmt.Sprintf("%v", item["type"]) == "property" && 
            fmt.Sprintf("%v", item["name"]) == propKey {

        // done
        item["value"]     = propValue
        item["tag"]       = propTag
        item["timestamp"] = int64(time.Now().Unix())

        // yes
        foundFlag = true
      
        // results
        newResults    = append(newResults, item)

      } else {

        // results
        newResults    = append(newResults, item)

      }

    }

    // did we find it
    if foundFlag == false {

      var item            = map[string]interface{}{}
      item["type"]        = "property"
      item["tags"]        = strings.Join(propTag, ",")
      item["name"]        = propKey
      item["value"]       = propValue
      item["timestamp"]   = int64(time.Now().Unix())
      newResults          = append(newResults, item)

    }

    api.Save(constants.GetDatabaseFile(), sig["key"].(string), newResults)

  }

}