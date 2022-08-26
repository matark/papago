package translator

type Result struct {
  Query       string   `json:"query"`
  Basic       Basic    `json:"basic"`
  Code        int      `json:"errorCode"`
  Translation []string `json:"translation"`
  Web         []Web    `json:"web"`
}

type Basic struct {
  Phonetic   string   `json:"phonetic"`
  USPhonetic string   `json:"us-phonetic"`
  UKPhonetic string   `json:"uk-phonetic"`
  Explains   []string `json:"explains"`
}

type Web struct {
  Key    string   `json:"key"`
  Values []string `json:"value"`
}

type NaverResult struct {
  Message NaverMessage `json:"message"`
}

type NaverMessage struct {
  Content NaverContent `json:"result"`
}

type NaverContent struct {
  Text string `json:"translatedText"`
}
