package main

import (
  "os"
  "path/filepath"
  "fmt"
  "log"
  "strings"
  "regexp"

  "github.com/codegangsta/cli"
  "github.com/mitchellh/go-homedir"
)

func main() {
  app := cli.NewApp()
  app.Name = "ghq-pwd"
  app.Usage = "shorten pwd path under the ghq."
  app.ActionFunc = func(c *cli.Context) {
    fmt.Println(ghqpwd())
  }

  app.Run(os.Args)
}

func ghqpwd() (string) {
    // get pwd
    fullPwd, pwdErr := filepath.Abs(filepath.Dir(os.Args[0]))
    if pwdErr != nil {
      log.Fatal(pwdErr)
    }

    // get $HOME
    homeDir, homeDirErr := homedir.Dir()
    if homeDirErr != nil {
      log.Fatal(homeDirErr)
    }

    rep := regexp.MustCompile(`^~/(.go/src|.ghq)/.+?/`)

    return rep.ReplaceAllString(strings.Replace(fullPwd, homeDir, "~", 1), "ghq:")
}
