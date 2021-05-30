package main

import (
	"os"

	config "github.com/kukkar/common-golang/pkg/config"
	appConf "github.com/kukkar/{{APP_NAME}}/conf"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.elastic.co/apm/module/apmgin"

	//_ sqlRepo "github.com/kukkar/common-golang/pkg/factory/sql"
	"strings"

	"github.com/kukkar/common-golang/globalconst"
	"github.com/kukkar/common-golang/pkg/factory/cache"
	"github.com/kukkar/common-golang/pkg/factory/sql"
	"github.com/kukkar/common-golang/pkg/healthcheck"
	"github.com/kukkar/common-golang/pkg/middleware"
	routes "github.com/kukkar/{{APP_NAME}}/src/routes"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const (
	confFile       = "conf/config.json"
	envFilePathEnv = "ENV_FILE_PATH"
)

func main() {
	//load env
	loadEnv()
	router := gin.New()
	//registering appconfig to global config
	registerConfig()
	//taking config into memory
	initConfig()
	//initiating logger
	initLogger(router)
	//registerning database mysql
	registerDBConfigMap()
	//register redis config map
	registerRedisConfigMap()
	//register setnry
	//recovery in case of panic registering
	router.Use(gin.Recovery())

	registerSentry(router)
	//register default routes
	registerDefaultRoutes(router)
	// registerning middlewares
	registerMiddleware(router)
	// registering apis
	registerApis(router)

	//register health check
	registerHealthCheck(router)
	initServer(router)
}

func loadEnv() {
	filePath := os.Getenv(envFilePathEnv)
	if filePath == "" {
		godotenv.Load()
	} else {
		godotenv.Load(filePath)
	}
}

func registerDBConfigMap() {
	conf, err := appConf.GetAppConfig()
	if err != nil {
		panic(err)
	}
	dbConfigMap := make(map[string]interface{})
	dbConfigMap[globalconst.DefaultDB] = (*conf.MySql)
	sql.InitConfigMap(dbConfigMap)
}

func registerRedisConfigMap() {
	conf, err := appConf.GetAppConfig()
	if err != nil {
		panic(err)
	}
	redisConfigMap := make(map[string]interface{})
	redisConfigMap[globalconst.DefaultRedisPoolKey] = (*conf.Cache)
	cache.InitConfigMap(redisConfigMap)
}

func registerApis(router *gin.Engine) {
	// register routing
	routes.Routes(router)
}

func initLogger(router *gin.Engine) {
	conf, err := appConf.GetGlobalConfig()
	if err != nil {
		panic(err)
	}
	err = conf.LogConfig.InitiateLogger()
	if err != nil {
		panic(err)
	}
}

//initConfig initialises the Global Application Config
func initConfig() {
	cm := new(config.ConfigManager)

	cm.InitializeGlobalConfig(confFile)
	cm.UpdateConfigFromEnv(config.GlobalAppConfig, "global")
	cm.UpdateConfigFromEnv(config.GlobalAppConfig.ApplicationConfig, "")
}

func registerConfig() {
	config.RegisterConfig(new(appConf.AppConfig))
	config.RegisterConfigEnvUpdateMap(appConf.EnvUpdateMap())
	config.RegisterGlobalEnvUpdateMap(config.GlobalEnvUpdateMap())
}

func initServer(router *gin.Engine) {
	conf, err := appConf.GetGlobalConfig()
	if err != nil {
		panic(err)
	}
	router.Run(conf.ServerHost + ":" + conf.ServerPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func registerMiddleware(router *gin.Engine) {
	router.Use(apmgin.Middleware(router))
	defaultMiddleware := middleware.DefaultMiddleware{}
	router.Use(middleware.DebugMiddleware())
	router.Use(defaultMiddleware.CORSMiddleware())
}

func registerHealthCheck(router *gin.Engine) {

	gConf, err := appConf.GetGlobalConfig()
	if err != nil {
		panic(err)
	}
	hConfig := healthcheck.Config{}
	group := router.Group(string(gConf.AppName))
	{
		group.GET("/healthcheck", healthcheck.HealthCheckHandler(healthcheck.GetHealthCheck(hConfig)))
	}
}

func registerDefaultRoutes(router *gin.Engine) {
	conf, err := appConf.GetGlobalConfig()
	if err != nil {
		panic(err)
	}
	if strings.ToLower(conf.Environment) == "dev" {
		pprof.Register(router)
		url := ginSwagger.URL("http://localhost:8086/swagger/doc.json") // The url pointing to API definition
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}

func registerSentry(router *gin.Engine) {
	conf, err := appConf.GetAppConfig()
	if err != nil {
		panic(err)
	}
	if conf.SentryDSN == "" {
		return
	}
	gConf, err := appConf.GetGlobalConfig()
	if err != nil {
		panic(err)
	}
	if gConf.Environment != "dev" {
		err = sentry.Init(sentry.ClientOptions{
			// Either set your DSN here or set the SENTRY_DSN environment variable.
			Dsn: conf.SentryDSN,
			// Enable printing of SDK debug messages.
			// Useful when getting started or trying to figure something out.
			Debug: true,
		})
	}
	if err != nil {
		panic(err)
	}
	router.Use(sentrygin.New(sentrygin.Options{}))
}
