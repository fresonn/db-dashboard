package main

import (
	"dashboard/api/internal/config"
	"dashboard/api/pkg/logger"
	"flag"
	"fmt"
	"log"
	"runtime"
	"runtime/debug"

	"github.com/joho/godotenv"
)

func init() {

	info, _ := debug.ReadBuildInfo()
	fmt.Println("Go version:", info.GoVersion, runtime.GOARCH)

	runtimeFlag := flag.String("runtime", "", "specify runtime flag value: -runtime=<value>")
	flag.Parse()

	if *runtimeFlag == config.DevRuntime {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
		fmt.Println("⚙️ Dev runtime, loaded .env")
	}
}

func main() {

	lgr := logger.New(logger.LiveLoggerVariant)

	cfg := config.New()

	lgr.Info("hello logger")

	fmt.Printf("Result config: %+v\n", cfg)
}
