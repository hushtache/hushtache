package commands

import(
  "github.com/hushtache/pki"
  "github.com/hushtache/utils"
  "github.com/hushtache/constants"
)

func Boot() {

  if utils.FileExists(constants.GetUserKey()) == false {

    // create a key we can use
    privateKeyPEM, _ := pki.GeneratePrivateKey(4 * 1024)

    // write the cert to file
    utils.FileWrite(constants.GetUserKey(), privateKeyPEM)

  }

}