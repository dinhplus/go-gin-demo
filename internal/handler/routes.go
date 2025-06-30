package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	// Essential middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(requestid.New())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(otelgin.Middleware("my-gin-app"))

	prom := ginprometheus.NewPrometheus("gin")
	prom.Use(r)

	// JWT Auth setup
	jwtMw := JwtMiddleware()
	if jwtMw == nil {
		panic("JWT middleware failed to initialize")
	}
	// Public endpoints
	r.POST("/login", LoginHandler)
	r.GET("/refresh_token", RefreshTokenHandler)
	r.POST("/register", RegisterHandler)
	r.POST("/forgot_password", ForgotPasswordHandler)

	// Protected routes
	auth := r.Group("/api")
	auth.Use(jwtMw.MiddlewareFunc())
	{
		auth.GET("/ping", Ping)
		auth.GET("/users", GetUsers)
		auth.POST("/users", CreateUser)
		auth.GET("/departments", GetDepartments)
		auth.POST("/departments", CreateDepartment)
		auth.GET("/stacks", GetStacks)
		auth.POST("/stacks", CreateStack)
		auth.GET("/positions", GetPositions)
		auth.POST("/positions", CreatePosition)
		// Example: admin only
		auth.GET("/admin", AdminOnly(), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "admin access granted"})
		})
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func SwaggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ...existing code for swagger handler..
		c.Next()

	}
}
