package routes

import (
	_ "bufio"
	"os"
	// Echo
	echo "github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	session "github.com/labstack/echo-contrib/session"
	sessions "github.com/gorilla/sessions"
	// Handler
	handler "practice-go-echo/handlers"
	// Moduls
	util "practice-go-echo/utils"
)

func Router() *echo.Echo {
	// Write log
	outputLogs, err := os.OpenFile("./resources/logs", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		util.OccurFatalError(err.Error())
	}
	// Create echo
	app := echo.New()
	// Set session
	store := sessions.NewCookieStore([]byte("secret"))
	// Set middleware
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}", "remote_ip":"${remote_ip}", "method":"${method}", "status":"${status}", "uri":"${uri}", "user_agent":"${user_agent}"}` + "\r\n",
		Output: outputLogs,
	}))
	app.Use(middleware.Recover())
	app.Use(session.MiddlewareWithConfig(session.Config{Store: store}))

	// Set static
	app.Static("/assets", "static")
	// Set Health check api
	app.GET("/health", handler.HealthCheck)

	// Create router (/user)
	userRouter := app.Group("user")
	{
		userRouter.File("/signin", "./views/signin.html")
		userRouter.File("/signup", "./views/signup.html")

		userRouter.POST("/signup", handler.UserSignup)
		userRouter.POST("/signin", handler.UserSignin)

		userRouter.GET("/", handler.UserMain)
	}
	sessionRouter := app.Group("sess")
	{
		sessionRouter.GET("/c", handler.CreateSession)
		sessionRouter.GET("/r", handler.ReadSession)
	}

	return app
}