package api

import(
  "github.com/hushtache/utils"
  "github.com/hushtache/gzip"
  "encoding/json"
  "fmt"
  "os"
)

/**
* Wrapper to check if the user is allowed in this store
**/
func Allowed(signatures []string, key string) (map[string]interface {}, error) {

  // loop the items
  for _,line := range signatures {

    // try to decrypt
    decrypted := DecryptSignature(key, line)

    // check if we got something and no error :)
    if decrypted != "" {

      // parse out the params sent in if any
      var item map[string]interface{}
      json.Unmarshal([]byte(decrypted), &item)

      // found it !
      return item, nil

    }

  }

  // nope out of here
  fmt.Println("The user is not in the list of trusted source :(")

  // done
  os.Exit(1)

  // done
  return nil, nil

}

/**
* Tries to decrypt the signatures using the current private key
* if any of them works, the user is allowed to decrypt the secrets :)
**/
func FindSignature(signatures []string, key string) (map[string]interface {}, error) {

  // loop the items
  for _,line := range signatures {

    // try to decrypt
    decrypted := DecryptSignature(key, line)

    // check if we got something and no error :)
    if decrypted != "" {

      // parse out the params sent in if any
      var item map[string]interface{}
      json.Unmarshal([]byte(decrypted), &item)

      // found it !
      return item, nil

    }

  }

  // return null and error
  return nil, nil

}

func FetchSignatures(filename string) ([]string, error) {

  content, _  := utils.FileRead(filename)

  // inflate
  inflatedText, _ := gzip.Inflate(content)

  // parse out the params sent in if any
  results := []string{}
  json.Unmarshal([]byte(inflatedText), &results)

  // return it
  return results, nil

}

func SaveSignatures(filename string, items []string) (bool) {

  // encrypt the item
  utils.ShuffleStrings(items)
  enc, _            := json.Marshal(items)

  // done
  deflatedText, _   := gzip.Deflate(string(enc))

  // done
  utils.FileWrite(filename, deflatedText)

  // return the item
  return true

}