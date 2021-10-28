package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type Server struct {
	app    *fiber.App
	wg     sync.WaitGroup
	reader Reader
}

func NewServer(rdr Reader) *Server {
	app := fiber.New(fiber.Config{
		ReadTimeout: 5 * time.Second,
	})
	app.Use(cors.New())
	v1 := app.Group("/v1")
	v1.Get("/dashboard", monitor.New())
	v1.Get("/credentials/:key", func(c *fiber.Ctx) error {
		k := c.Params("key")
		cred, err := rdr.Read(k)
		if err != nil {
			fmt.Printf("error %v", err)
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"success":    true,
			"credential": &cred,
		})
	})

	return &Server{
		app:    app,
		wg:     sync.WaitGroup{},
		reader: rdr,
	}
}

func (s *Server) Start(port string) error {
	s.wg.Add(1)
	var err error
	go func() {
		defer s.wg.Done()
		err := s.app.Listen(fmt.Sprintf(":%v", port))

		if err != nil {
			fmt.Printf("API Server stopped listening due to %v", err)
		}
	}()

	return err
}

func (s *Server) Shutdown() error {
	err := s.app.Shutdown()

	if err != nil {
		s.wg.Wait()
	}

	return err
}
