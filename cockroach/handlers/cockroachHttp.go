package handlers

import (
  "net/http"

  "go_project_structure/cockroach/models"
  "go_project_structure/cockroach/services"
  "github.com/labstack/echo/v4"

  "github.com/labstack/gommon/log"
)

type cockroachHttpHandler struct {
 cockroachService services.CockroachService
}

func NewCockroachHttpHandler(cockroachService services.CockroachService) CockroachHandler {
  return &cockroachHttpHandler{
    cockroachService: cockroachService,
  }
}

func (h *cockroachHttpHandler) DetectCockroach(c echo.Context) error {
  reqBody := new(models.AddCockroachData)

  if err := c.Bind(reqBody); err != nil {
    log.Errorf("Error binding request body: %v", err)
    return response(c, http.StatusBadRequest, "Bad request")
  }

  if err := h.cockroachService.CockroachDataProcessing(reqBody); err != nil {
    return response(c, http.StatusInternalServerError, "Processing data failed")
  }

  return response(c, http.StatusOK, "Success 🪳🪳🪳")
}