package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PutItemRequest struct {
	Name string `json:"name"`
}

type API struct {
	mu    sync.RWMutex
	items map[int]Item
}

func newAPI() *API {
	return &API{
		items: make(map[int]Item),
	}
}

func (a *API) registerRoutes(e *echo.Echo) {
	e.GET("/item/:id", a.getItem)
	e.PUT("/item/:id", a.putItem)
}

func (a *API) putItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var req PutItemRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "name is required"})
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	item := Item{ID: id, Name: req.Name}
	a.items[item.ID] = item

	return c.JSON(http.StatusOK, item)
}

func (a *API) getItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	item, found := a.items[id]
	if !found {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "item not found"})
	}

	return c.JSON(http.StatusOK, item)
}

func zapRequestLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			latency := time.Since(start)

			logger.Info("http_request",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Path()),
				zap.Int("status", c.Response().Status),
				zap.Duration("latency", latency),
			)

			return err
		}
	}
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	e := echo.New()
	e.Use(zapRequestLogger(logger))
	e.Use(middleware.Recover())

	api := newAPI()
	api.items[1] = Item{ID: 1, Name: "sample"}
	api.registerRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
