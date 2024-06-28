package router

import (
	"net/http"
	comm "sca/internal/dto/request"
	request "sca/internal/dto/request/target"

	"github.com/gin-gonic/gin"
)

func (r *Router) AddTarget(gin *gin.Context) {
	ctx := r.context(gin)

	var req request.AddTarget
	if bindErr := transBindJSON(gin, &req, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}

	resp, httpErr := r.targCont.AddTarget(ctx, &req)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusCreated, resp, nil)
}

func (r *Router) UpdateTarget(gin *gin.Context) {
	ctx := r.context(gin)

	var uri comm.InternalUUIDSelector
	if bindErr := transBindUri(gin, &uri, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}
	var req request.UpdateTarget
	if bindErr := transBindJSON(gin, &req, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}

	resp, httpErr := r.targCont.UpdateTarget(ctx, &uri, &req)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusOK, resp, nil)
}

func (r *Router) DeleteTarget(gin *gin.Context) {
	ctx := r.context(gin)

	var uri comm.InternalUUIDSelector
	if bindErr := transBindUri(gin, &uri, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}

	httpErr := r.targCont.DeleteTarget(ctx, &uri)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusNoContent, nil, nil)
}
