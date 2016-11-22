package pki

import (
  "crypto/rand"
  "encoding/base64"
  "encoding/pem"
  "crypto"
  "crypto/x509"
  "crypto/sha256"
  "crypto/rsa"
  "golang.org/x/crypto/ssh"
  "fmt"
)

/**
* Encrypts the given text using AES with the given key, 
* the key is hashed with SHA256 to enforce bit length and
* allow longer keys to be given
*
* @author:  github.com/johanndutoit
* @date:    20 November 2016
**/
func Encrypt(key string, text string) (string, error) {

  // decode PEM encoding to ANS.1 PKCS1 DER
  block, _ := pem.Decode([]byte(key))
  pub, err := x509.ParsePKIXPublicKey(block.Bytes)
  pubkey, _ := pub.(*rsa.PublicKey)   

  // create the propertiess
  message := []byte(text)
  label := []byte("")
  hash := sha256.New()

  ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pubkey, message, label)
  return string(base64.StdEncoding.EncodeToString(ciphertext)), err

}

func Decrypt(key string, text string) (string, error) {

  // decode PEM encoding to ANS.1 PKCS1 DER
  block, _ := pem.Decode([]byte(key))
  priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)

  message, err := base64.StdEncoding.DecodeString(string(text))
  if err != nil {
    return "", err
  }
  
  label := []byte("")
  hash := sha256.New()

  // Decrypt Message
  plainText, err := rsa.DecryptOAEP(hash, rand.Reader, priv, message, label)  

  return string(plainText), err

}

func Verify(key string, text string) (bool, error) {

  // decode PEM encoding to ANS.1 PKCS1 DER
  block, _ := pem.Decode([]byte(key))
  pub, err := x509.ParsePKIXPublicKey(block.Bytes)
  fmt.Println(pub)

  return true, err

}

func Sign(key string, text string) (string, error) {

  // decode PEM encoding to ANS.1 PKCS1 DER
  block, _ := pem.Decode([]byte(key))
  privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

  message := []byte(text)

  hash := sha256.New()
  hash.Write(message)
  hashed := hash.Sum(nil)

  // Message - Signature
  var opts rsa.PSSOptions
  opts.SaltLength = rsa.PSSSaltLengthAuto

  signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed, &opts)

  return string(base64.StdEncoding.EncodeToString(signature)), err

}

/**
* Returns the public key from the private key
**/
func GeneratePublicKey(key string) (string, error) {

  // decode PEM encoding to ANS.1 PKCS1 DER
  block, _ := pem.Decode([]byte(key))
  // if block == nil { fmt.Printf("No Block found in keyfile\n"); os.Exit(1) }
  // if block.Type != "RSA PRIVATE KEY" { fmt.Printf("Unsupported key type"); os.Exit(1) }

  // parse DER format to a native type
  privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

  bb, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)

  pemdata := pem.EncodeToMemory(
    &pem.Block{
      Type: "PUBLIC KEY",
      Bytes: bb,
    },
  )

  return string(pemdata), nil

}

/**
* Returns the public key from the private key
**/
func GeneratePublicAuthKey(key string) (string, error) {

  // decode PEM encoding to ANS.1 PKCS1 DER
  block, _ := pem.Decode([]byte(key))
  // if block == nil { fmt.Printf("No Block found in keyfile\n"); os.Exit(1) }
  // if block.Type != "RSA PRIVATE KEY" { fmt.Printf("Unsupported key type"); os.Exit(1) }

  // parse DER format to a native type
  privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

  pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
  if err != nil {
    return "", err
  }
  return string(ssh.MarshalAuthorizedKey(pub)), nil

}

/**
*
**/
func GeneratePrivateKey(bitlevel int) (string, error) {

  key, err := rsa.GenerateKey(rand.Reader, bitlevel)
  if err != nil {
    return "", err
  }

  pemdata := pem.EncodeToMemory(
    &pem.Block{
      Type: "RSA PRIVATE KEY",
      Bytes: x509.MarshalPKCS1PrivateKey(key),
    },
  )

  return string(pemdata), nil

}