package aes

import (
  "crypto/rand"
  "encoding/base64"
  "crypto/sha256"
  "crypto/aes"
  "crypto/cipher"
  "errors"
  "io"
)

/**
* Util function to return the SHA256 of the given text
*
* @author:  github.com/johanndutoit
* @date:    20 November 2016
**/
func SHA256(text string) string {

  h := sha256.New()
  h.Write([]byte(text))
  return string(h.Sum(nil))

}

/**
* Encrypts the given text using AES with the given key, 
* the key is hashed with SHA256 to enforce bit length and
* allow longer keys to be given
*
* @author:  github.com/johanndutoit
* @date:    20 November 2016
**/
func Encrypt(key string, text string) (string, error) {

  salt := SHA256(key)

  block, err := aes.NewCipher([]byte(string(salt)))
  if err != nil {
      return "", err
  }
  b := []byte(text)
  ciphertext := make([]byte, aes.BlockSize+len(b))
  iv := ciphertext[:aes.BlockSize]
  if _, err := io.ReadFull(rand.Reader, iv); err != nil {
      return "", err
  }
  cfb := cipher.NewCFBEncrypter(block, iv)
  cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
  return string(base64.StdEncoding.EncodeToString(ciphertext)), nil
}

/**
* Tries to decrypt the given text with the given key, assuming it was
* encrypted with padded AES. The key is made a SHA256 hash to enforce
* bit length vales.
*
* @author:  github.com/johanndutoit
* @date:    20 November 2016
**/
func Decrypt(key string, text string) (string, error) {
  
  salt := SHA256(key)

  block, err := aes.NewCipher([]byte(string(salt)))
  if err != nil {
      return "", err
  }
  if len(text) < aes.BlockSize {
      return "", errors.New("ciphertext too short")
  }

  dataf, err := base64.StdEncoding.DecodeString(string(text))
  if err != nil {
      return "", err
  }

  formatting := []byte(dataf)
  iv := formatting[:aes.BlockSize]
  formatting = formatting[aes.BlockSize:]
  cfb := cipher.NewCFBDecrypter(block, iv)
  cfb.XORKeyStream(formatting, formatting)
  return string(formatting), nil

}