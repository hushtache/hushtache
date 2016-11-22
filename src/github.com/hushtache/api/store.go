package api

import(
  "github.com/hushtache/gzip"
  "github.com/hushtache/utils"
  "github.com/hushtache/aes"
  "encoding/json"
  "time"
)

func AlreadyConfigured(filename string) bool {

  if utils.FileExists(filename) {
    return true
  } else {
    return false
  }

}

func Create() (interface{}) {

  var result = map[string]interface{}{}
  result["timestamp"] = int32(time.Now().Unix())
  return result

}

func Fetch(filename string, password string) ([]interface{}, error) {

  content, _  := utils.FileRead(filename)

  // decrypt the line
  decryptedLine, _ := aes.Decrypt(aes.SHA256(password), content)

  // inflate
  inflatedText, _ := gzip.Inflate(decryptedLine)

  // parse out the params sent in if any
  var results []interface{}
  json.Unmarshal([]byte(inflatedText), &results)

  // return it
  return results, nil

}

func Save(filename string, password string, items []interface{}) (bool) {

  // encrypt the item
  utils.ShuffleObjects(items)
  enc, _            := json.Marshal(items)
  deflatedText, _   := gzip.Deflate(string(enc))
  result, _         := aes.Encrypt(aes.SHA256(password), deflatedText)

  // done
  utils.FileWrite(filename, result)

  // return the item
  return true

}

func Get(key string) {}