package gzip

import(
  
  "bytes"
  "compress/gzip"
  "bufio"
  "encoding/base64"
  "io/ioutil"
  
)

func Deflate(text string) (string, error) {

  var buf bytes.Buffer
  bufzip := gzip.NewWriter(&buf)
  bufWrite := bufio.NewWriter(bufzip)
  bufWrite.WriteString(text)
  bufWrite.Flush()
  bufzip.Close()
  return string(base64.StdEncoding.EncodeToString(buf.Bytes())), nil

}

func Inflate(text string) (string, error) {

  data, err := base64.StdEncoding.DecodeString(string(text))
  if err != nil {
      return "", err
  }

  r, _ := gzip.NewReader(bytes.NewBuffer(data))
  defer r.Close()

  content, err := ioutil.ReadAll(r)
  return string(content), err

}