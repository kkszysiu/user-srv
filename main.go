package main

import (
	"github.com/codegangsta/cli"
	log "github.com/golang/glog"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	"github.com/micro/user-srv/db"
	"github.com/micro/user-srv/handler"
	proto "github.com/micro/user-srv/proto/account"
)

func main() {
	app := cli.NewApp()
	app.HideVersion = true
	app.Flags = cmd.Flags
	app.Flags = append(app.Flags,
		cli.StringFlag{
			Name:   "database_url",
			EnvVar: "DATABASE_URL",
			Usage:  "The database URL e.g root@tcp(127.0.0.1:3306)/user",
		},
	)

	app.Before = cmd.Setup
	app.Action = func(c *cli.Context) {
		db.Url = c.String("database_url")
	}
	app.RunAndExitOnError()

	db.Init()

	server.Init(
		server.Name("go.micro.srv.user"),
	)

	proto.RegisterAccountHandler(server.DefaultServer, new(handler.Account))

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
