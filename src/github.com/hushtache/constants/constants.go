package constants

import(
  "os/user"
  "github.com/hushtache/utils"
  "fmt"
)

func GetSignatureFile() string {

  // get the arguments
  _, opts := utils.ParseArgs()

  // check if test was defined
  if fmt.Sprintf("%v", opts["Keys"]) != "" {

    // return that
    return "" + fmt.Sprintf("%v", opts["Keys"])

  } else {

    // return that file !
    return ".hushtache.keys"

  }

}

func GetDatabaseFile() string {

  // get the arguments
  _, opts := utils.ParseArgs()

  // check if test was defined
  if fmt.Sprintf("%v", opts["File"]) != "" {

    // return that
    return "" + fmt.Sprintf("%v", opts["File"])

  } else {

    // return that file !
    return ".hushtache"

  }

}

func GetHomeDir() string {

  usr, _ := user.Current()
  return usr.HomeDir

}

func GetUserKey() string {

  // get the arguments
  _, opts := utils.ParseArgs()

  // check if test was defined
  if fmt.Sprintf("%v", opts["Key"]) != "" {

    // return that
    return "" + fmt.Sprintf("%v", opts["Key"])

  } else {

    // return that file !
    return GetHomeDir() + "/.hushtache.key"

  }
  
} 

func GetRounds() int {

  return 20

} 