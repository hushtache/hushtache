package commands

import(
  "github.com/aymerick/raymond"
  "github.com/hushtache/api"
  "github.com/hushtache/utils"
  "github.com/hushtache/constants"
  "fmt"
  "strings"
  "os"
)

func Render() {

  // get the key
  key           := api.GetCurrentKey()

  // get it
  args, _       := utils.ParseArgs()

  // get the sigs
  sigs, _       := api.FetchSignatures(constants.GetSignatureFile())

  // find the certain one
  sig, _        := api.FindSignature(sigs, key)

  // did we find it ?
  if sig == nil {

    fmt.Println("NOT ALLOWED !")

    os.Exit(1)
    
  }

  // get all the results
  results, _       := api.Fetch(constants.GetDatabaseFile(), sig["key"].(string))

  // create the context
  ctx := map[string]string{}

  // loop the results
  for _,entry := range results {

    item := entry.(map[string]interface{})

    // check if defined
    if(item["type"].(string) == "property") {

      // get the label
      label := fmt.Sprintf("%v", item["name"])
      value := fmt.Sprintf("%v", item["value"])

      // add to the list
      ctx[strings.ToUpper(label)] = value
      ctx[strings.ToLower(label)] = value

    }

  }

  // check if any were given
  if len(args) == 1 {

    // done
    fmt.Println("Template files should be given ...")

    // exit
    os.Exit(1)

  }

  for _,file := range args[1:] {

    // read from file
    contents, err := utils.FileRead(args[1])

    // check for error
    if err != nil {

      // output error
      fmt.Println("Problem reading file: " + file)

      // done
      os.Exit(1)

    }

    // render the template
    result, err := raymond.Render(contents, ctx)

    // check for a error
    if err != nil {
      panic("Please fill a bug :)")
    }

    fmt.Print(result)

  }

}