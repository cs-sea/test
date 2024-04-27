package controllers

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Api struct {
	esClient *elasticsearch.Client
}

func NewApi(esClient *elasticsearch.Client) *Api {
	return &Api{esClient: esClient}
}

func (a *Api) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (a *Api) Health(c echo.Context) error {
	return c.String(http.StatusOK, "green!")
}

func (a *Api) EsQuery(c echo.Context) error {
	return nil
}
