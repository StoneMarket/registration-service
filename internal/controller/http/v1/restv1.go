package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/StoneMarket/registration-service/internal/models"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type V1 struct {
	c *Controller
}

func (v1 *V1) RestV1(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "api/v1/user/create":
		if fmt.Sprint(ctx.Method()) == "POST" {
			v1.createHandler(ctx)
		} else {
			undefinedHandler(ctx)
		}
	case "/api/v1/user/login":
		if fmt.Sprint(ctx.Method()) == "POST" {
			v1.loginHandler(ctx)
		} else {
			undefinedHandler(ctx)
		}
	default:
		undefinedHandler(ctx)
	}
}

func undefinedHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetBodyRaw([]byte("wrong url path or method for this service"))
	ctx.Response.SetStatusCode(http.StatusBadRequest)
}

func (v1 *V1) loginHandler(ctx *fasthttp.RequestCtx) {
	rawJson := ctx.Request.Body()
	if rawJson == nil {
		v1.c.logger.Error("login, empty body", zap.String("url", string(ctx.Path())), zap.String("IP", string(ctx.RemoteIP())))
		ctx.Response.SetBodyRaw([]byte("empty body"))
		ctx.Response.SetStatusCode(http.StatusBadRequest)

		return
	}

	var user models.User
	if err := json.Unmarshal(rawJson, &user); err != nil {
		v1.c.logger.Error("parse", zap.Error(err), zap.String("json", string(rawJson)))
		ctx.Response.SetBodyRaw([]byte("parse error"))
		ctx.Response.SetStatusCode(http.StatusBadRequest)

		return
	}

	token, err := v1.c.Login(context.Background(), user.Login, user.Password)
	if err != nil {
		ctx.Response.SetBodyRaw([]byte("error with login"))
		ctx.Response.SetStatusCode(http.StatusInternalServerError)

		return
	}

	rawJson, err = json.Marshal(&models.Response{AccessLogin: string(token), Type: "Bearer"})
	if err != nil {
		v1.c.logger.Error("parse", zap.Error(err), zap.Any("response", string(rawJson)))
		ctx.Response.SetBodyRaw([]byte("parse error"))
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
	}

	ctx.Response.SetBodyRaw(rawJson)
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (v1 *V1) createHandler(ctx *fasthttp.RequestCtx) {
	rawJson := ctx.Request.Body()
	if rawJson == nil {
		v1.c.logger.Error("login, empty body", zap.String("url", string(ctx.Path())), zap.String("IP", string(ctx.RemoteIP())))
		ctx.Response.SetBodyRaw([]byte("empty body"))
		ctx.Response.SetStatusCode(http.StatusBadRequest)

		return
	}

	var user models.User
	if err := json.Unmarshal(rawJson, &user); err != nil {
		v1.c.logger.Error("parse", zap.Error(err), zap.String("json", string(rawJson)))
		ctx.Response.SetBodyRaw([]byte("parse error"))
		ctx.Response.SetStatusCode(http.StatusBadRequest)

		return
	}

	token, err := v1.c.RegisterNewUser(context.Background(), &user)
	if err != nil {
		ctx.Response.SetBodyRaw([]byte("error with registration"))
		ctx.Response.SetStatusCode(http.StatusInternalServerError)

		return
	}

	rawJson, err = json.Marshal(&models.Response{AccessLogin: string(token), Type: "Bearer"})
	if err != nil {
		v1.c.logger.Error("parse", zap.Error(err), zap.Any("response", string(rawJson)))
		ctx.Response.SetBodyRaw([]byte("parse error"))
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
	}

	ctx.Response.SetBodyRaw(rawJson)
	ctx.Response.SetStatusCode(http.StatusOK)
}
