package commands

import(
  "github.com/hushtache/api"
  "github.com/hushtache/utils"
  "github.com/hushtache/constants"
  "fmt"
  "strings"
  "os"
)

func Get() {

  // get the arguments
  args, _ := utils.ParseArgs()

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

  output := ""
  count := 0;

  for _,entry := range results {

    item := entry.(map[string]interface{})

    // check the type
    if fmt.Sprintf("%v", item["type"]) == "property" {

      if len(args) == 1 {

        // count property
        count = count + 1

        // output
        output = output + fmt.Sprintf("%v", item["name"])
        output = output + ": " + fmt.Sprintf("%v", item["value"]) + "\n"

      } else if len(args) > 1 {

        // the values
        propKey       := strings.ToLower(args[1]) 

        if fmt.Sprintf("%v", item["type"]) == "property" && 
            fmt.Sprintf("%v", item["name"]) == propKey {

          // count property
          count = count + 1

          output = fmt.Sprintf("%v", item["value"])

          break

        }

      }

    }

  }

  if len(args) == 2 {

    if count > 0 {
    
      fmt.Println(output)

    } else {

      fmt.Println("")

    }

  } else {

    if count > 0 {
    
      fmt.Println(strings.TrimSpace(output))

    } else {

      fmt.Println("No secrets found ...")

    }

  }

}