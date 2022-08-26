package translator

import "encoding/json"
import "io/ioutil"
import "net/http"
import "log"
import "os"

func Translate(text string) Result {
  key := os.Getenv("PAPAGO_API_KEY")
  response, wrong := http.Get(apiUrl(key, text))

  if wrong != nil {
    log.Fatal(wrong)
  }

  var result Result
  data, wrong := ioutil.ReadAll(response.Body)

  if wrong != nil {
    log.Fatal(wrong)
  }

  json.Unmarshal(data, &result)
  return result
}

func NaverTranslate(text string) NaverResult {
  client := &http.Client{}
  key := os.Getenv("PAPAGO_API_KEY")
  response, wrong := client.Do(apiRequest(key, text))

  if wrong != nil {
    log.Fatal(wrong)
  }

  var result NaverResult
  data, wrong := ioutil.ReadAll(response.Body)

  if wrong != nil {
    log.Fatal(wrong)
  }

  json.Unmarshal(data, &result)
  return result
}
