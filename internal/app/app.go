package app

import (
	"net/http"
	"sca/internal/config"
	"sca/internal/controller"
	"sca/internal/log"
	"sca/internal/provider/db/sql/postgres"
	catPg "sca/internal/repository/cat/postgres"
	misPg "sca/internal/repository/mission/postgres"
	targPg "sca/internal/repository/target/postgres"
	"sca/internal/router"
	"sca/internal/service"
	"time"
)

func Run() error {
	conf, err := config.NewConfig()
	if err != nil {
		return err
	}
	log := log.NewLogger(conf.Server.LogMode, conf.Server.LogPath)
	cli := &http.Client{Timeout: 5 * time.Second}

	transServ := service.NewTranslationService()

	db, err := postgres.NewPostgresProvider(conf)
	if err != nil {
		return err
	}
	conn := db.Conn()
	catPg := catPg.NewCatPostgresRepo(conn)
	misPg := misPg.NewMissionPostgresRepo(conn)
	targPg := targPg.NewTargetPostgresRepo(conn)

	catCont := controller.NewCatController(catPg)
	misCont := controller.NewMissionController(misPg)
	targCont := controller.NewTargetController(targPg, misPg)

	router := router.NewRouter(log, conf, cli, transServ, catCont, misCont, targCont)

	return router.SetupHandlers()
}
