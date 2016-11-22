package commands

import(
  "github.com/hushtache/pki"
  "github.com/hushtache/gzip"
  "github.com/hushtache/api"
  "fmt"
)

func PrivateKey() {

  // get the key
  key           := api.GetCurrentKey()

  // get our pub key :)
  pubKey, _     := pki.GeneratePublicKey(key)
  decrypted, _  := gzip.Deflate(pubKey)

  // output
  fmt.Println( decrypted )

}