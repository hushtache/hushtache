package utils

import(
  "github.com/jessevdk/go-flags"
  "github.com/fatih/structs"
)

var (
  args []string
  opts map[string]interface {}
)
var loaded = false

func ParseArgs() ([]string, map[string]interface {}) {

  // check if defined
  if loaded == false {

    // set opts
    var optss struct {
      
      Tag string `short:"t" long:"tag" description:"Tags (delimited by comma) of the property to either set or filter with" value-name:"FILE"`
      File string `short:"f" long:"file" description:"Store to instead of default .hushtache in current folder" value-name:"FILE"`
      Keys string `short:"s" long:"signature" description:"Signature file to instead of default .hushtache.sig in current folder" value-name:"FILE"`
      Key string `short:"k" long:"key" description:"Key file to instead of default .hushtache.sig in current folder" value-name:"FILE"`

    }

    // set args
    args, _ = flags.Parse(&optss)

    // convert and set
    opts = structs.Map(optss)

    // mark as loaded
    loaded = true

    // return them
    return args, opts

  } else {

    // return them
    return args, opts

  }

}