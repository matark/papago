package main

import "flag"
import "fmt"
import "os"
import "strings"
import "neweos.de/treely/tree"
import "neweos.de/papago/lolcat"
import "neweos.de/papago/translator"

const Version = "1.0.1"

func main() {
  flag.Parse()
  explains := make([]tree.Input, 0)
  translations := make([]tree.Input, 0)

  if len(flag.Args()) == 0 {
    fmt.Println("No query")
    os.Exit(0)
  }

  result := translator.Translate(strings.Join(flag.Args(), " "))

  for _, explain := range result.Basic.Explains {
    explains = append(explains, tree.Input{
      Name: explain,
    })
  }

  translation := tree.Input{
    Name: result.Translation[0],
    Contents: explains,
  }

  translations = append(translations, translation)

  for _, web := range result.Web {
    translations = append(translations, tree.Input{
      Name: fmt.Sprintf("%s: %s", web.Key, strings.Join(web.Values, " ")),
    })
  }

  input := tree.Input{
    Name: fmt.Sprintf("%s [%s]", result.Query, result.Basic.Phonetic),
    Contents: translations,
  }

  lolcat.Rainbow([]rune(tree.Tree(append(make([]tree.Input, 0), input))))
}
