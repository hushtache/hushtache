package hushtache

import (
  "fmt"
  "strings"
  "github.com/hushtache/commands"
  "github.com/hushtache/utils"
)

func Boot() {

  // get the arguments
  args, _ := utils.ParseArgs()

  // boot up the client
  commands.Boot()

  // check if any given ...
  if len(args) == 0 {

    // init dir here
    commands.Get()
  
  } else {

    // lower the command
    cmd     := strings.ToLower(args[0])

    if strings.Index(cmd, "init") == 0 {

      commands.Init()

    } else if strings.Index(cmd, "users") == 0 {

      commands.Users()

    } else if strings.Index(cmd, "allowed") == 0 {

      commands.Allowed()

    } else if strings.Index(cmd, "allow") == 0 {

      commands.Allow()

    } else if strings.Index(cmd, "deny") == 0 {

      commands.Deny()

    } else if strings.Index(cmd, "key") == 0 {

      commands.PrivateKey()

    } else if strings.Index(cmd, "render") == 0 {

      commands.Render()

    } else if strings.Index(cmd, "get") == 0 {

      commands.Get()

    } else if strings.Index(cmd, "remove") == 0 {

      commands.Remove()

    } else if strings.Index(cmd, "set") == 0 {

      commands.Set()

    }  else {

      fmt.Println("Unknown command ... Check --help for more details")

    }

  }

}