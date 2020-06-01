/*
 * Copyright (c) 2020.  This software is made for the Black Rock City Aid group and is provided AS IS with no support or liability under the Apache 2 license.
 */

package apiimpl

import (
	"fmt"
	"github.com/bmason42/opencrisisline2/pkg/errors"
	"github.com/bmason42/opencrisisline2/pkg/generated/v1"
	"github.com/bmason42/opencrisisline2/pkg/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func sendError(c *gin.Context, err error) {
	httpCode, errs := handleErrors(c, err)
	c.JSON(httpCode, errs)
}

func aboutGetUnversioned(c *gin.Context) {
	var resp v1.AboutResponse
	resp.AppVersion = "1.0"
	resp.ApiVersions = make([]string, 0)
	resp.ApiVersions = append(resp.ApiVersions, "v1")

	c.JSON(http.StatusOK, resp)
	log.Println("In about")
}

type Message struct {
	Body string
}
type Response struct {
	Message Message
}

func callbackHandler(c *gin.Context) {
	c.Request.ParseForm()
	for key, value := range c.Request.Form {
		fmt.Printf("%s = %s \n", key, value)
	}
	var r Response
	r.Message.Body = "hello back " + c.Request.Form.Get("Body")
	c.XML(200, r)
}
func postHandler(c *gin.Context) {
	var help v1.HelpRequest
	err := c.ShouldBindJSON(&help)
	if err != nil {
		weberr := MkErrorResponse(errors.OCERROR_ERROR, errors.ERROR_CODE_INVALID_USER_INPUT, c, map[string]string{"data": err.Error()})
		c.JSON(400, weberr)
	}
	ms := model.NewMessageSystem()
	err = ms.SendText(help.PhoneNumber, help.CallerName, help.Message)
	if err != nil {
		weberr := MkErrorResponse(errors.OCERROR_ERROR, errors.ERROR_CODE_UNKNOWN, c, map[string]string{"data": err.Error()})
		c.JSON(500, weberr)
	}
	c.JSON(201, "")

}
func healthCheckGetUnversioned(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func swaggerUIGetHandler(c *gin.Context) {
	c.Redirect(302, "/opencrisisline2/swaggerui/index_v1.html")
}
