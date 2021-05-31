package router

import (
	"log"

	"github.com/LotteWong/giotto-gateway-admin/admin_middleware"
	"github.com/LotteWong/giotto-gateway-admin/controller"
	"github.com/LotteWong/giotto-gateway-admin/docs"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)

	// health check route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	redisAddr := lib.GetStringSliceConf("redis_map.list.default.proxy_list")[0]
	store, err := sessions.NewRedisStore(10, "tcp", redisAddr, "", []byte("secret"))
	if err != nil {
		log.Fatalf("sessions.NewRedisStore err: %v", err)
	}
	adminMiddlewares := []gin.HandlerFunc{
		sessions.Sessions("gateway_session", store),
		admin_middleware.IpAuthMiddleware(),
	}
	enableRateLimiter := lib.GetBoolConf("base.rate_limiter.enable")
	if enableRateLimiter {
		adminMiddlewares = append(adminMiddlewares, admin_middleware.RateLimitMiddleware())
	}
	enableCircuitBreaker := lib.GetBoolConf("base.circuit_breaker.enable")
	if enableCircuitBreaker {
		adminMiddlewares = append(adminMiddlewares, admin_middleware.CircuitBreakMiddleware())
	}

	// swagger api routes
	docs.SwaggerInfo.Title = lib.GetStringConf("base.swagger.title")
	docs.SwaggerInfo.Description = lib.GetStringConf("base.swagger.desc")
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = lib.GetStringConf("base.swagger.host")
	docs.SwaggerInfo.BasePath = lib.GetStringConf("base.swagger.base_path")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// login api routes
	loginGroup := router.Group("")
	loginGroup.Use(adminMiddlewares...)
	{
		// POST   /login
		// POST   /logout
		controller.RegistLoginRoutes(loginGroup)
	}

	// user api routes
	userGroup := router.Group("/users")
	userGroup.Use(adminMiddlewares...)
	userGroup.Use(admin_middleware.SessionAuthMiddleware())
	{
		// GET    /users/admin
		// PATCH  /users/admin
		controller.RegistUserRoutes(userGroup)
	}

	// service api routes
	serviceGroup := router.Group("/services")
	serviceGroup.Use(adminMiddlewares...)
	serviceGroup.Use(admin_middleware.SessionAuthMiddleware())
	{
		// GET    /services
		// GET    /services/:service_id
		// POST   /services/http
		// PUT    /services/http/:service_id
		// POST   /services/tcp
		// PUT    /services/tcp/:service_id
		// POST   /services/grpc
		// PUT    /services/grpc/:service_id
		// DELETE /services/:service_id
		controller.RegistServiceRoutes(serviceGroup)
	}

	// app api routes
	appGroup := router.Group("/apps")
	appGroup.Use(adminMiddlewares...)
	appGroup.Use(admin_middleware.SessionAuthMiddleware())
	{
		// GET    /apps
		// GET    /apps/:app_id
		// POST   /apps
		// PUT    /apps/:app_id
		// DELETE /apps/:app_id
		controller.RegistAppRoutes(appGroup)
	}

	// dashboard api routes
	dashboardGroup := router.Group("/dashboard")
	dashboardGroup.Use(adminMiddlewares...)
	dashboardGroup.Use(admin_middleware.SessionAuthMiddleware())
	{
		// GET    /dashboard/statistics
		// GET    /dashboard/flow
		// GET    /dashboard/flow/services/:service_id
		// GET    /dashboard/flow/apps/:app_id
		// GET    /dashboard/percentage/services
		controller.RegistDashboardRoutes(dashboardGroup)
	}

	// static resource route
	router.Static("/dist", "./dist")

	return router
}
