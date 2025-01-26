package main

import (
	"asme8/linker/src/config"
	"asme8/linker/src/core"
	"asme8/linker/src/object"
	"flag"
	"fmt"
	"log"
)

func main() {
	// Define flags
	configPath := flag.String("config", "", "Path to the linker config file")
	flag.Parse()

	// Get remaining arguments as .o files
	objectFilePaths := flag.Args()

	// Check if config path is provided
	if *configPath == "" {
		fmt.Println("Config path is required")
		return
	}

	// Load config
	conf := config.NewConfig()
	err := conf.ParseConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	objectFiles, err := object.ReadObjectFiles(objectFilePaths...)
	if err != nil {
		fmt.Println(err)
		return
	}

	l := core.NewLinker(conf, objectFiles...)
	err = l.Link()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OK")
}
