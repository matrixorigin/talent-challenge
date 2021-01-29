package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matrixorigin/talent-challenge/matrixbase/distributed/pkg/cfg"
	"github.com/matrixorigin/talent-challenge/matrixbase/distributed/pkg/model"
	"github.com/matrixorigin/talent-challenge/matrixbase/distributed/pkg/store"
)

// Server http api server
type Server struct {
	cfg    cfg.Cfg
	store  store.Store
	engine *gin.Engine
}

// NewServer create the server
func NewServer(cfg cfg.Cfg) (*Server, error) {
	s, err := store.NewStore(cfg.Store)
	if err != nil {
		return nil, err
	}

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	return &Server{
		cfg:    cfg,
		store:  s,
		engine: engine,
	}, nil
}

// Start start the server
func (s *Server) Start() error {
	s.engine.GET("/kv", s.doGet)
	s.engine.POST("/kv", s.doSet)
	s.engine.DELETE("/kv", s.doDelete)
	return s.engine.Run(s.cfg.API.Addr)
}

// Stop stop the server
func (s *Server) Stop() error {
	return nil
}

func (s *Server) doGet(c *gin.Context) {
	key := []byte(c.Query("key"))
	value, err := s.store.Get(key)
	if err != nil {
		c.JSON(http.StatusOK, returnError(err))
		return
	}

	c.JSON(http.StatusOK, returnData(value))
}

func (s *Server) doSet(c *gin.Context) {
	req := &model.Request{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, returnError(err))
		return
	}

	err = s.store.Set([]byte(req.Key), []byte(req.Value))
	if err != nil {
		c.JSON(http.StatusOK, returnError(err))
		return
	}

	c.JSON(http.StatusOK, returnData("OK"))
}

func (s *Server) doDelete(c *gin.Context) {
	key := []byte(c.Query("key"))
	err := s.store.Delete(key)
	if err != nil {
		c.JSON(http.StatusOK, returnError(err))
		return
	}

	c.JSON(http.StatusOK, returnData("OK"))
}
