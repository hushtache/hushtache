package utils

import(
  "io/ioutil"
  "os"
)

/**
* Util function to write a string to file
* @author github.com/johanndutoit
* @date   21 November 2016
**/
func FileWrite(filename string, content string) error {

  // write to file
  return ioutil.WriteFile(filename, []byte(content), 0644)

}

/**
* Util function to read a string from file
* @author github.com/johanndutoit
* @date   21 November 2016
**/
func FileRead(filename string) (string, error) {

  // write to file
  dat, err := ioutil.ReadFile(filename)
  if err != nil {
    return "", err
  } else {
    return string(dat), nil
  }

}

/**
* Util function to check if the file exists
* @author github.com/johanndutoit
* @date   21 November 2016
**/
func FileExists(filename string) (bool) {

  if _, err := os.Stat(filename); err == nil {
    return true
  } else {
    return false
  }

}