package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"

  "github.com/mh-cbon/gen-version-file/GenVersionFile"
  "github.com/urfave/cli"
)

func main() {

  templates := map[string]string{
    "go": `package GenVersionFile

func Version () string {
  return "$VERSION"
}

`,
  }
  type LangWriter func (string) error

  writers := map[string]LangWriter{
    "go": func (content string) error {
      err := os.MkdirAll("GenVersionFile", 0766)
      if err!=nil {
        return err
      }
      return ioutil.WriteFile("GenVersionFile/index.go", []byte(content), 0766)
    },
  }

  app := cli.NewApp()
  app.Name = "gen-version-file"
  app.Version = GenVersionFile.Version()
  app.Usage = "Generate version file"
  app.UsageText = "gen-version-file --ver=0.0.1 --lang=go"
  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "ver",
      Value: "",
      Usage: "Version number to write",
    },
    cli.StringFlag{
      Name: "lang, l",
      Value: "go",
      Usage: "Language to generate for.",
    },
  }

  app.Action = func(c *cli.Context) error {
    version := c.String("ver")
    lang := c.String("lang")

    if version=="" {
      cli.ShowAppHelp(c)
      return cli.NewExitError("Version argument is required", 1)
    }

    template, ok := templates[lang]
    if ok==false {
      return cli.NewExitError("Sorry, there is no template for this language: " + lang, 1)
    }
    template = strings.Replace(template, "$VERSION", version, -1)

    writer, ok := writers[lang]
    if ok==false {
      return cli.NewExitError("Sorry, there is no writer for this language: " + lang, 1)
    }

    err := writer(template)
    if err==nil {
      fmt.Println("Ok, version file wrote.")
    }

    return err
  }

  app.Run(os.Args)
}
