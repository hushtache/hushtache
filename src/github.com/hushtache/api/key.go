package api

import(
  "github.com/hushtache/pki"
  "github.com/hushtache/utils"
  "github.com/hushtache/constants"
)

/**
*
**/
func DecryptSignature(key string, content string) string {

  // encrypt the generated password, append it
  output, _ := pki.Decrypt(key, content)

  // encrypt
  // output, _ = aes.Decrypt(aes.SHA256(key), output)

  // return
  return output

}

/**
*
**/
func EncryptSignature(key string, content string) string {

  // encrypt
  // output, _ := aes.Encrypt(aes.SHA256(key), content)

  // encrypt the generated password, append it
  finalEncryption, _ := pki.Encrypt(key, content)

  // done
  return finalEncryption

}

/**
* @author github.com/johanndutoit
* @date   21 November 2016
**/
func GetCurrentKey() string {

  if utils.FileExists(constants.GetUserKey()) == false {

    // create a key we can use
    privateKeyPEM, _ := pki.GeneratePrivateKey(4 * 1024)

    // write the cert to file
    utils.FileWrite(constants.GetUserKey(), privateKeyPEM)

    // return our key
    return privateKeyPEM

  } else {

    // write the cert to file
    privateKeyPEM, _ := utils.FileRead(constants.GetUserKey())

    // check it
    return privateKeyPEM

  }

}