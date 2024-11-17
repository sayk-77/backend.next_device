package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"log"
	"os"
	"time"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"next_device/backend/db"
	"next_device/backend/di"
	"next_device/backend/tools"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/process"
)

var (
	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_server_cpu_usage_percentage",
		Help: "CPU usage percentage for the Go server process.",
	})
	memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_server_memory_usage_bytes",
		Help: "Memory usage for the Go server process in bytes.",
	})
)

func updateMetrics() {
	pid := os.Getpid()
	proc, err := process.NewProcess(int32(pid))
	if err != nil {
		log.Printf("Ошибка при получении процесса: %v", err)
		return
	}

	for {
		cpuPercent, err := proc.CPUPercent()
		if err == nil {
			cpuUsage.Set(cpuPercent)
			log.Printf("Обновление CPU: %.2f%%\n", cpuPercent)
		}

		memInfo, err := proc.MemoryInfo()
		if err == nil {
			memoryUsage.Set(float64(memInfo.RSS))
			log.Printf("Обновление памяти: %.2f GB\n", float64(memInfo.RSS)/1e9)
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)

	go updateMetrics()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	prometheusMiddleware := fiberprometheus.New("next_device_app")
	prometheusMiddleware.RegisterAt(app, "/metrics")
	app.Use(prometheusMiddleware.Middleware)

	app.Get("/metric", monitor.New(monitor.Config{Title: "backend.nextDevice (fiber)"}))

	dataBase, err := db.SetupDb()
	if err != nil {
		log.Fatal("Нет соединения с базой данных")
	}

	db.Migrate(dataBase)
	di.InitDependencies(app, dataBase)

	tools.GetImageProduct(app)
	tools.GetImageBrand(app)
	tools.GetImageBanner(app)
	tools.GetImageCategory(app)
	tools.GetReviewImage(app)

	if err := app.Listen("localhost:5000"); err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}

	fmt.Print("Сервер запущен")
}
