package controller

import (
	"context"
	"net/http"

	"github.com/anggunpermata/custom-webhook/config"
	"github.com/anggunpermata/custom-webhook/models"
	"github.com/labstack/echo"
)

func AssignAgentWebhook(c echo.Context) error {
	var customerReq models.InitiateReq
	if err := c.Bind(&customerReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Not valid",
		})
	}

	resp, err := AgentAllocationWebhook(c.Request().Context(), customerReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "request not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": resp,
	})
}

func AgentAllocationWebhook(ctx context.Context, req models.InitiateReq) (interface{}, error) {
	//agent used: anggun@qiscus.cx
	agentId := config.GoDotEnvVariable("AgentId")
	r, err := AssignAgent(req, agentId)
	if err != nil {
		return nil, err
	}
	return r, nil
}
