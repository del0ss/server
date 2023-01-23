package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"smth/internal/model"
	"smth/internal/store"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *gin.Engine
	store  *store.Store
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: gin.Default(),
	}
}

func (s *Server) Start() error {

	s.configureRouter()

	err := s.configureStore()

	s.logger.Info("starting server")
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8000", s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *Server) configureRouter() {
	s.router.LoadHTMLGlob("templates/*.html")
	s.router.GET("/", s.handleHello())
	s.router.GET("/auth", s.singUpPage())
	s.router.POST("/auth/singup", s.registerNewUser())
}

func (s *Server) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *Server) handleHello() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	}

}

func (s *Server) singUpPage() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.HTML(200, "singup.html", gin.H{})
	}

}

func (s *Server) registerNewUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := model.User{Login: ctx.PostForm("login"), Email: ctx.PostForm("email"), EncPassword: ctx.PostForm("encPassword")}
		fmt.Println(user)
		result, err := s.store.User().Create(&user)
		if err != nil {
			return
		}
		ctx.Redirect(301, "/")
		fmt.Println(result)
	}

}
