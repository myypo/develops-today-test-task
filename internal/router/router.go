package router

import (
	stdCtx "context"
	"net/http"
	"sca/internal/config"
	"sca/internal/context"
	"sca/internal/controller"
	"sca/internal/dto/response"
	herr "sca/internal/error/http"
	"sca/internal/service"

	middleware "sca/internal/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	log  *zap.Logger
	conf *config.Config

	cli *http.Client

	transServ service.TranslationService

	catCont  controller.CatController
	misCont  controller.MissionController
	targCont controller.TargetController
}

func NewRouter(
	log *zap.Logger,
	conf *config.Config,

	cli *http.Client,

	transServ service.TranslationService,

	catCont controller.CatController,
	misCont controller.MissionController,
	targCont controller.TargetController,
) Router {
	return Router{
		log, conf,

		cli,

		transServ,

		catCont, misCont, targCont,
	}
}

const basePath = "/api/v1"

func (r *Router) SetupHandlers() error {
	h := gin.Default()
	gin.SetMode(string(r.conf.Server.LogMode))

	h.Use(middleware.NewLoggerMiddleware(r.log))

	a := h.Group(basePath)

	catsGroup := a.Group("/cats")
	{
		catsGroup.POST("", r.CreateCat)
		catsGroup.PATCH("/:id", r.UpdateCat)
		catsGroup.DELETE("/:id", r.DeleteCat)
		catsGroup.GET("/:id", r.GetCat)
		catsGroup.GET("", r.ListCats)
	}

	missionsGroup := a.Group("/missions")
	{
		missionsGroup.POST("", r.CreateMission)
		missionsGroup.PATCH("/:id", r.UpdateMission)
		missionsGroup.DELETE("/:id", r.DeleteMission)
		missionsGroup.GET("/:id", r.GetMission)
		missionsGroup.GET("", r.ListMissions)
	}

	targetsGroup := a.Group("/targets")
	{
		targetsGroup.POST("", r.AddTarget)
		targetsGroup.PATCH("/:id", r.UpdateTarget)
		targetsGroup.DELETE("/:id", r.DeleteTarget)
	}

	return h.Run(r.conf.Server.Address())
}

func transBindUri[T any](
	gin *gin.Context,
	req *T,
	tsServ service.TranslationService,
) herr.ErrorHttp {
	if err := gin.BindUri(req); err != nil {
		return tsServ.TranslateEN(err)
	}
	return nil
}

func transBindJSON[T any](
	gin *gin.Context,
	req *T,
	tsServ service.TranslationService,
) herr.ErrorHttp {
	if err := gin.BindJSON(req); err != nil {
		return tsServ.TranslateEN(err)
	}
	return nil
}

func sendResponse(gin *gin.Context, statusCode int, data any, err error) {
	var errString string
	if err != nil {
		errString = err.Error()
	}
	resp := response.NewBaseResponse(data, errString)
	gin.JSON(statusCode, resp)
}

func (r *Router) context(stdCtx stdCtx.Context) *context.Context {
	return context.NewContext(stdCtx, r.log, r.cli)
}
