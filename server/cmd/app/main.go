package main

import (
	"Remote_XML_Parser/internal/logging"
	"Remote_XML_Parser/internal/routes"
	"Remote_XML_Parser/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"sort"
)

func main() {

	err := godotenv.Load("./server/.env")
	if err != nil {
		log.Fatal(err)
	}

	var envVars map[string]string
	envVars, _ = godotenv.Read()
	keys := make([]string, 0, len(envVars))
	for key := range envVars {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var conf services.Config
	err = envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err)
	}

	logging.SetupLogger(conf.Env)
	logger := logging.GetLogger()
	conf.Logger = logger
	logger.Infof("%+v\n", conf)

	//conf.ConnectDatabase()
	//
	//defer func() {
	//	sqlDB, err := conf.PGClient.DB()
	//	if err != nil {
	//		log.Fatalf("Failed to close database connection")
	//	}
	//	err = sqlDB.Close()
	//	if err != nil {
	//		log.Fatalf("Failed to close database connection")
	//	}
	//}()

	mode := gin.ReleaseMode
	if conf.Env == "local" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	address := fmt.Sprintf(":%v", conf.ServicePort)
	h := routes.NewHandler(&conf)

	server := &http.Server{
		Addr:    address,
		Handler: h2c.NewHandler(h, &http2.Server{}),
	}
	logger.Infof("Listening on %s", address)
	logger.Fatal(server.ListenAndServe())

}
