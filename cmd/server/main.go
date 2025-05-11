package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/emaldie/secret-api/internal/server/config"
	"github.com/emaldie/secret-api/internal/server/db"
	"github.com/go-playground/validator"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var programLevel = new(slog.LevelVar)

func main() {
	programLevel.Set(slog.LevelInfo)

	configPath := getConfigPath()
	setupLogger(configPath)

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		slog.Error("Error loading config", "error", err)
		os.Exit(1)
	}

	setLogLevel((*cfg).Log.Level)
	slog.Info("Set logging level", "config_path", configPath)

	app, err := initApp(cfg)
	if err != nil {
		slog.Error("Failed to initialize application", "error", err)
		os.Exit(1)
	}

	serverErrCh := startServer(app, (*cfg).Server.Port, (*cfg).Server.ReadTimeout, (*cfg).Server.WriteTimeout)

	shutdownApp(app, serverErrCh)
}

type App struct {
	DB        *mongo.Client
	Redis     *redis.Client
	Router    *http.ServeMux
	Validator *validator.Validate
	Server    *http.Server
}

func initApp(cfg *config.AppConfig) (*App, error) {
	slog.Info("Starting app initialization...")

	slog.Info("Connecting to database...")
	database, err := db.InitMongo(&cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("error initializing MongoDB: %w", err)
	}
	slog.Info("Successfully connected to MongoDB")

	slog.Info("Connecting to Redis...")
	redisClient, err := db.InitRedis(&cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("error initializing Redis: %w", err)
	}
	slog.Info("Successfully connected to Redis")

	validate := validator.New()

	// slog.Info("Initializing DI system...")
	// deps := injection.NewDependencies(
	// 	database,
	// 	redisClient,
	// 	validate,
	// 	cfg,
	// 	slog.Default(),
	// )
	// slog.Info("Successfully initialized DI system")

	router := http.NewServeMux()

	// slog.Info("Setting up API...")
	// api.Setup(router, api.RouterConfig{
	// })
	// slog.Info("Successfully set up API")

	return &App{
		DB:     database,
		Redis:  redisClient,
		Router: router,
		//Cache:     cacheInstance,
		Validator: validate,
		//Deps:      deps,
	}, nil
}

func startServer(app *App, port int, readTimeout, writeTimeout time.Duration) <-chan error {
	errCh := make(chan error, 1)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      app.Router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	app.Server = server

	go func() {
		slog.Info("Server start up", "port", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("server err: %w", err)
		}
	}()

	return errCh
}

func setupLogger(configPath string) {
	cfg, err := config.LoadConfig(configPath)
	var output io.Writer = os.Stdout
	handlerOptions := &slog.HandlerOptions{
		AddSource: true,
		Level:     programLevel,
	}

	if err != nil {
		slog.Warn("Failed to load log config using default console text output", "error", err)
	} else {
		if cfg.Log.File != "" {
			currentDate := time.Now().Format("2006-01-02")

			logDir := filepath.Dir(cfg.Log.File)
			logFileName := filepath.Base(cfg.Log.File)
			ext := filepath.Ext(logFileName)
			nameWithoutExt := strings.TrimSuffix(logFileName, ext)
			newLogFileName := fmt.Sprintf("%s-%s%s", nameWithoutExt, currentDate, ext)
			logFilePath := filepath.Join(logDir, newLogFileName)

			if mkDirErr := os.MkdirAll(logDir, 0755); mkDirErr != nil {
				slog.Error("Unable to create log directory", "dir", logDir, "error", mkDirErr)
			} else {
				logFile, openErr := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
				if openErr != nil {
					slog.Error("Unable to create log file", "path", logFilePath, "error", openErr)
				} else {
					if cfg.Log.Console {
						output = io.MultiWriter(os.Stdout, logFile)
						slog.Info("Logs will be output to both the console and the file", "file", logFilePath)
					} else {
						output = logFile
						slog.Info("Logs will be output to a file", "file", logFilePath)
					}
				}
			}
		} else {
			slog.Info("Logs will be output to console")
		}
	}

	logger := slog.New(slog.NewTextHandler(output, handlerOptions))
	slog.SetDefault(logger)
}

func getConfigPath() string {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/config.yaml"
	}
	return configPath
}

func setLogLevel(level string) {
	var l slog.Level
	switch level {
	case "debug":
		l = slog.LevelDebug
	case "info":
		l = slog.LevelInfo
	case "warn":
		l = slog.LevelWarn
	case "error":
		l = slog.LevelError
	default:
		slog.Warn("Unknown log level, Info level will be used instead", "configured_level", level)
		l = slog.LevelInfo
	}
	programLevel.Set(l)
	slog.Info("The log level is set to", "level", l.String())
}

func shutdownApp(app *App, serverErrCh <-chan error) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case err := <-serverErrCh:
		slog.Error("Server error", "error", err)
	case sig := <-signalCh:
		slog.Info("Received sys signal to begin graceful shutdown", "signal", sig.String())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if app.Server != nil {
		slog.Info("Shutting down the server...")
		if err := app.Server.Shutdown(ctx); err != nil {
			slog.Error("Failed to shut down the server", "error", err)
		}
	}

	if app.DB != nil {
		slog.Info("Disconnecting from MongoDB...")
		err := app.DB.Disconnect(ctx)
		if err != nil {
			slog.Error("Error disconnecting from MongoDB", "error", err)
		}
	}

	if app.Redis != nil {
		slog.Info("Disconnecting from Redis...")
		if err := app.Redis.Close(); err != nil {
			slog.Error("Error disconnecting from Redis", "error", err)
		}
	}

	slog.Info("Application shutdown complete")
}
