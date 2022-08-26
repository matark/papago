package translator

import "net/http"
import "net/url"
import "strings"
import "bytes"
import "fmt"

func apiUrl(key string, text string) string {
  keys := strings.Split(key, ":")
  baseURL, _ := url.Parse("https://fanyi.youdao.com/openapi.do")

  query := baseURL.Query()
  query.Add("keyfrom", keys[0])
  query.Add("key", keys[1])
  query.Add("doctype", "json")
  query.Add("version", "1.1")
  query.Add("type", "data")
  query.Add("q", text)

  baseURL.RawQuery = query.Encode()
  return baseURL.String()
}

func apiRequest(key string, text string) *http.Request {
  keys := strings.Split(key, ":")
  data := fmt.Sprintf(`{"source":"%s","target":"%s","text":"%s"}`, "en", "ko", text)
  request, _ := http.NewRequest(http.MethodPost, "https://openapi.naver.com/v1/papago/n2mt", bytes.NewBuffer([]byte(data)))

  request.Header.Add("X-Naver-Client-Id", keys[0])
  request.Header.Add("X-Naver-Client-Secret", keys[1])
  request.Header.Add("Content-Type", "application/json")

  return request
}
