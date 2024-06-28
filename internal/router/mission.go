package router

import (
	"net/http"
	comm "sca/internal/dto/request"
	request "sca/internal/dto/request/mission"

	"github.com/gin-gonic/gin"
)

func (r *Router) CreateMission(gin *gin.Context) {
	ctx := r.context(gin)

	var req request.CreateMission
	if bindErr := transBindJSON(gin, &req, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}

	resp, httpErr := r.misCont.CreateMission(ctx, &req)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusCreated, resp, nil)
}

func (r *Router) UpdateMission(gin *gin.Context) {
	ctx := r.context(gin)

	var uri comm.InternalUUIDSelector
	if bindErr := transBindUri(gin, &uri, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}
	var req request.UpdateMission
	if bindErr := transBindJSON(gin, &req, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}

	resp, httpErr := r.misCont.UpdateMission(ctx, &uri, &req)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusOK, resp, nil)
}

func (r *Router) DeleteMission(gin *gin.Context) {
	ctx := r.context(gin)

	var uri comm.InternalUUIDSelector
	if bindErr := transBindUri(gin, &uri, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}

	httpErr := r.misCont.DeleteMission(ctx, &uri)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusNoContent, nil, nil)
}

func (r *Router) GetMission(gin *gin.Context) {
	ctx := r.context(gin)

	var uri comm.InternalUUIDSelector
	if bindErr := transBindUri(gin, &uri, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}

	resp, httpErr := r.misCont.GetMission(ctx, &uri)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusOK, resp, nil)
}

func (r *Router) ListMissions(gin *gin.Context) {
	ctx := r.context(gin)

	resp, httpErr := r.misCont.ListMissions(ctx)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusOK, resp, nil)
}
