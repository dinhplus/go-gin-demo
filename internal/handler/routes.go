package handler

import (
	docs "my-gin-app/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type HandlerList struct {
	LoginHandler          gin.HandlerFunc
	RefreshTokenHandler   gin.HandlerFunc
	RegisterHandler       gin.HandlerFunc
	ForgotPasswordHandler gin.HandlerFunc
	Ping                  gin.HandlerFunc
	GetUsers              gin.HandlerFunc
	CreateUser            gin.HandlerFunc
	GetDepartments        gin.HandlerFunc
	CreateDepartment      gin.HandlerFunc
	GetStacks             gin.HandlerFunc
	CreateStack           gin.HandlerFunc
	GetPositions          gin.HandlerFunc
	CreatePosition        gin.HandlerFunc
	AdminOnly             gin.HandlerFunc
}

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
	r.GET("/ping", Ping)

	// GraphQL endpoints
	r.POST("/graphql", GraphQLHandler())
	r.GET("/playground", PlaygroundHandler())

	// Protected routes
	// @BasePath /
	// @securityDefinitions.apikey BearerAuth
	// @in header
	// @type http
	// @name Authorization
	// @description JWT Authorization header using the Bearer scheme. Example: "Authorization: Bearer {token}"
	// @Security BearerAuth
	auth := r.Group("")
	auth.Use(jwtMw.MiddlewareFunc())
	{

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

		// RBAC/Feature flag endpoints
		auth.GET("/permissions", GetPermissions)
		auth.POST("/permissions", CreatePermission)
		auth.GET("/resources", GetResources)
		auth.POST("/resources", CreateResource)
		auth.GET("/feature-flags", GetFeatureFlags)
		auth.POST("/feature-flags", CreateFeatureFlag)
	}

	docs.SwaggerInfo.Title = "Dinh DZ app"
	docs.SwaggerInfo.Description = "This is a sample server for a Gin application."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// Swagger setup
	r.GET("/swagger/*any", SwaggerHandler())

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func SwaggerHandler() gin.HandlerFunc {
	return ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DeepLinking(true),
		ginSwagger.DocExpansion("none"),
	)
}
