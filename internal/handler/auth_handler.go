package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"
	"time"

	"my-gin-app/internal/model"
	"my-gin-app/internal/repository"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var identityKey = "id"

func getPasswordSecret() string {
	secret := os.Getenv("PASSWORD_SECRET")
	if secret == "" {
		secret = "your-strong-secret-key" // fallback, should be replaced with a strong secret
	}
	return secret
}

func hashPasswordWithSecret(password string) string {
	h := hmac.New(sha256.New, []byte(getPasswordSecret()))
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

func JwtMiddleware() *jwt.GinJWTMiddleware {
	mw, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "my-gin-app",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBindJSON(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			// Find user in DB by username (email)
			user, err := findUserByEmail(username)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			if hashPasswordWithSecret(password) != user.Password {
				return nil, jwt.ErrFailedAuthentication
			}
			return map[string]interface{}{
				identityKey: user.Email,
				"role":      user.Role.Name,
				"role_id":   user.RoleID,
				"user_id":   user.ID,
			}, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(map[string]interface{}); ok {
				return jwt.MapClaims{
					identityKey: v[identityKey],
					"role":      v["role"],
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return map[string]interface{}{
				identityKey: claims[identityKey],
				"role":      claims["role"],
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(map[string]interface{}); ok {
				// Allow all authenticated users
				return v[identityKey] != nil
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	return mw
}

// findUserByEmail find user by email (username)
func findUserByEmail(email string) (model.User, error) {
	return repository.FindUserByEmail(email)
}

// Example of a protected endpoint with role-based access
type Role string

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		if claims["role"] != RoleAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin only"})
			return
		}
		c.Next()
	}
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login  body      Login  true  "Login credentials"
// @Success      200   {object}  map[string]interface{}
// @Failure      401   {object}  map[string]string
// @Router       /login [post]
func LoginHandler(c *gin.Context) {
	JwtMiddleware().LoginHandler(c)
}

// RefreshToken godoc
// @Summary      Refresh JWT token
// @Description  Refresh an existing JWT token
// @Tags         auth
// @Produce      json
// @Success      200   {object}  map[string]interface{}
// @Failure      401   {object}  map[string]string
// @Router       /refresh_token [get]
func RefreshTokenHandler(c *gin.Context) {
	JwtMiddleware().RefreshHandler(c)
}

// AdminOnly godoc
// @Summary      Admin only endpoint
// @Description  Only accessible by admin role
// @Tags         admin
// @Produce      json
// @Success      200   {object}  map[string]string
// @Failure      403   {object}  map[string]string
// @Router       /api/admin [get]

// Register godoc
// @Summary      Register user
// @Description  Register a new user and store securely hashed password in the database. Username must be unique and password is hashed with a secret key before saving. Returns user_id if successful.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      Login  true  "User registration info"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /register [post]
func RegisterHandler(c *gin.Context) {
	var req Login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if len(req.Username) < 3 || len(req.Password) < 6 {
		c.JSON(400, gin.H{"error": "Username must be at least 3 chars, password at least 6 chars"})
		return
	}
	// Check for duplicate email
	_, err := repository.FindUserByEmail(req.Username)
	if err == nil {
		c.JSON(400, gin.H{"error": "Email already exists"})
		return
	}
	// Hash password with secret
	hash := hashPasswordWithSecret(req.Password)
	user := model.User{
		Email:    req.Username,
		Password: hash,
		Name:     req.Username,
		RoleID:   2,
		// Assuming default values for demo purposes
		DepartmentID: 1,
		PositionID:   1,
		StackID:      1,
	}
	err = repository.CreateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "registered", "user_id": user.ID})
}

// ForgotPassword godoc
// @Summary      Forgot password
// @Description  Send password reset instructions
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        email  body      map[string]string  true  "User email"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Router       /forgot_password [post]
func ForgotPasswordHandler(c *gin.Context) {
	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "reset instructions sent (demo only)", "email": req["email"]})
}
