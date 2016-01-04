package main

import (
	"os"

	"bitbucket.org/containeriq/wharf/cmd"
	"bitbucket.org/containeriq/wharf/setting"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func main() {
	if err := setting.SetConfig("conf/containeriq.conf"); err != nil {
		log.Error("Read config failed: %v", err.Error())
		return
	}

	app := cli.NewApp()
	app.Name = setting.AppName
	app.Usage = setting.Usage
	app.Version = setting.Version
	app.Author = setting.Author
	app.Email = setting.Email

	app.Commands = []cli.Command{
		cmd.CmdWeb,
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)

}
