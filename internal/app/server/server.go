package server

import (
	"fmt"
	"hack2023/internal/app/config"
	"hack2023/internal/app/model"
	"hack2023/internal/app/store"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	router *echo.Echo
	store  *store.Store
	config config.Config
}

// NewServer инициализируем сервер
func NewServer(store *store.Store, config config.Config) *server {
	s := &server{
		router: echo.New(),
		store:  store,
		config: config,
	}

	// Конфигурируем роутинг
	s.configureRouter()
	return s
}

// Start Включаем прослушивание HTTP протокола
func (s *server) Start(config config.Config) error {
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	err := s.router.Start(address)
	if err != nil {
		return err
	}
	return nil
}

// configureRouter Объявляем список доступных роутов
func (s *server) configureRouter() {
	s.router.Use(middleware.RequestID())
	s.router.Use(middleware.Logger())
	s.router.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	s.router.POST("/token", s.handleToken)
	s.router.POST("/login", s.login)
	s.router.GET("/", s.handleVersion)
	api := s.router.Group("/api", s.ErrorHandler)
	{
		api.GET("/status", s.handleStatus)
		v1 := api.Group("/v1", s.ErrorHandler)
		{
			v1.Use(middleware.Logger())
			v1.GET("/typelist", s.typelist)
			v1.POST("/typelist", s.typelistFilter)
			authGroup := v1.Group("/", s.ErrorHandler)
			{
				authGroup.Use(echojwt.WithConfig(echojwt.Config{
					ParseTokenFunc: s.ParseTokenFunc,
				}))
			}
		}
	}
}

func (s *server) ParseTokenFunc(c echo.Context, tokenString string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SigningKey), nil
	})

	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
