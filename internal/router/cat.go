package router

import (
	"net/http"
	comm "sca/internal/dto/request"
	request "sca/internal/dto/request/cat"

	"github.com/gin-gonic/gin"
)

func (r *Router) CreateCat(gin *gin.Context) {
	ctx := r.context(gin)

	var req request.CreateCat
	if bindErr := transBindJSON(gin, &req, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}

	resp, httpErr := r.catCont.CreateCat(ctx, &req)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusCreated, resp, nil)
}

func (r *Router) UpdateCat(gin *gin.Context) {
	ctx := r.context(gin)

	var uri comm.InternalUUIDSelector
	if bindErr := transBindUri(gin, &uri, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}
	var req request.UpdateCat
	if bindErr := transBindJSON(gin, &req, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}

	resp, httpErr := r.catCont.UpdateCat(ctx, &uri, &req)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusOK, resp, nil)
}

func (r *Router) DeleteCat(gin *gin.Context) {
	ctx := r.context(gin)

	var uri comm.InternalUUIDSelector
	if bindErr := transBindUri(gin, &uri, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}

	httpErr := r.catCont.DeleteCat(ctx, &uri)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusNoContent, nil, nil)
}

func (r *Router) GetCat(gin *gin.Context) {
	ctx := r.context(gin)

	var uri comm.InternalUUIDSelector
	if bindErr := transBindUri(gin, &uri, r.transServ); bindErr != nil {
		sendResponse(gin, http.StatusBadRequest, nil, bindErr)
		return
	}

	resp, httpErr := r.catCont.GetCat(ctx, &uri)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusOK, resp, nil)
}

func (r *Router) ListCats(gin *gin.Context) {
	ctx := r.context(gin)

	resp, httpErr := r.catCont.ListCats(ctx)
	if httpErr != nil {
		sendResponse(gin, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(gin, http.StatusOK, resp, nil)
}
