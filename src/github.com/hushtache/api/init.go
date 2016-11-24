package api

import(
  "github.com/hushtache/pki"
  "github.com/hushtache/utils"
  "github.com/hushtache/constants"
  "github.com/nu7hatch/gouuid"
  "encoding/json"
  "fmt"
)

func Init(key string, username string) {

  // get a random password we can use
  password := utils.RandomString(200)

  // create a id for the user
  u, _ := uuid.NewV4()

  sigs := []string{}
  var storeItems []interface{}
  for i := 0; i < 1; i++ {

    // resulting item
    item              := map[string]interface{}{}
    item["uid"]       = u.String()
    item["label"]     = username
    item["key"]       = password
    enc, _            := json.Marshal(item)
    
    // get the public key
    publicKeyPEM, _ := pki.GeneratePublicKey(key)

    // get it
    signedPassword := EncryptSignature(publicKeyPEM, string(enc))

    sigs              = append(sigs, signedPassword)

    // resulting item
    storeItem              := map[string]interface{}{}
    storeItem["type"]       = "user"
    storeItem["label"]      = username
    storeItem["signature"]  = signedPassword
    storeItems = append(storeItems, storeItem)

  }

  // write it
  SaveSignatures(constants.GetSignatureFile(), sigs)

  // save to the store
  Save(constants.GetDatabaseFile(), password, storeItems)

  // output to let the user know !
  fmt.Println("hushtache initialized !")

}