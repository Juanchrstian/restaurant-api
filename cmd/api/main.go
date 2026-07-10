package main

import (
	"log"
	"net/http"
	"time"

	"github.com/juanchrstian/restaurant-api/seed"

	"github.com/juanchrstian/restaurant-api/internal/health"
	"github.com/juanchrstian/restaurant-api/internal/menu"
	"github.com/juanchrstian/restaurant-api/internal/order"
	"github.com/juanchrstian/restaurant-api/internal/router"
	"github.com/juanchrstian/restaurant-api/internal/session"

	"github.com/juanchrstian/restaurant-api/internal/shared/cache"
	"github.com/juanchrstian/restaurant-api/internal/shared/config"
	"github.com/juanchrstian/restaurant-api/internal/shared/database"
	redisdb "github.com/juanchrstian/restaurant-api/internal/shared/redis"
)

func main() {

	// =========================================
	// LOAD CONFIGURATION
	// =========================================

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting %s...", cfg.App.Name)

	// =========================================
	// CONNECT DATABASE
	// =========================================

	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✓ PostgreSQL Connected")

	// =========================================
	// CONNECT REDIS
	// =========================================

	redisClient, err := redisdb.New(cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✓ Redis Connected")

	cacheClient := cache.NewRedis(
		redisClient,
	)

	// =========================================
	// HEALTH MODULE
	// =========================================

	// Sementara service belum menggunakan DB dan Redis.
	// Nanti constructor akan berubah menjadi:
	// health.NewService(db, redisClient)

	healthService := health.NewService()

	healthHandler := health.NewHandler(healthService)

	menuRepository := menu.NewRepository(db)

	menuService := menu.NewService(

		menuRepository,

		cacheClient,
	)

	menuHandler := menu.NewHandler(menuService)

	sessionRepository := session.NewRepository(db)

	sessionService := session.NewService(

		sessionRepository,
	)

	sessionHandler := session.NewHandler(sessionService)

	orderRepository := order.NewRepository(db)

	orderService := order.NewService(
		db,

		orderRepository,

		menuRepository,

		sessionService,
	)

	orderHandler := order.NewHandler(orderService)

	// =========================================
	// ROUTER
	// =========================================

	engine := router.New(
		healthHandler,

		menuHandler,

		sessionHandler,

		orderHandler,
	)

	// =========================================
	// DATABASE SEED
	// =========================================

	if err := seed.Run(db); err != nil {
		log.Fatal(err)
	}

	log.Println("✓ Database Seeded")

	// =========================================
	// HTTP SERVER
	// =========================================

	server := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Server listening on :%s", cfg.App.Port)

	if err := server.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
