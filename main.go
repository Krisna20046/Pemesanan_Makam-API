package main

import (
	// "embed"
	"fmt"
	// "net/http"
	"sync"
	"time"

	_ "embed"

	"github.com/Krisna20046/db"
	"github.com/Krisna20046/handler/api"
	"github.com/Krisna20046/middleware"
	"github.com/Krisna20046/model"
	repo "github.com/Krisna20046/repository"
	service "github.com/Krisna20046/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPIHandler      api.UserAPI
	MakamAPIHandler     api.MakamAPI
	JenazahAPIHandler   api.JenazahAPI
	PemesananAPIHandler api.PemesananAPI
}

// type ClientHandler struct {
// 	AuthWeb      web.AuthWeb
// 	HomeWeb      web.HomeWeb
// 	DashboardWeb web.DashboardWeb
// 	ModalWeb     web.ModalWeb
// }

//go..:embed views/*
// var Resources embed.FS

func main() {
	gin.SetMode(gin.ReleaseMode) //release

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		router := gin.New()
		db := db.NewDB()
		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] \"%s %s %s\"\n",
				param.TimeStamp.Format(time.RFC822),
				param.Method,
				param.Path,
				param.ErrorMessage,
			)
		}))
		router.Use(gin.Recovery())

		dbCredential := model.Credential{
			Host:         "localhost",
			Username:     "postgres",
			Password:     "admin",
			DatabaseName: "grave",
			Port:         5432,
			Schema:       "public",
		}

		conn, err := db.Connect(&dbCredential)
		if err != nil {
			panic(err)
		}

		conn.AutoMigrate(&model.User{}, &model.Session{}, &model.DataJenazah{}, &model.DataMakam{}, &model.Pemesanan{})

		router = RunServer(conn, router)
		// router = RunClient(conn, router, Resources)

		fmt.Println("Server is running on localhost:8080")
		err = router.Run(":8080")
		if err != nil {
			panic(err)
		}

	}()

	wg.Wait()
}

func RunServer(db *gorm.DB, gin *gin.Engine) *gin.Engine {
	userRepo := repo.NewUserRepo(db)
	sessionRepo := repo.NewSessionsRepo(db)
	makamRepo := repo.NewMakamRepo(db)
	jenazahRepo := repo.NewJenazahRepo(db)
	pemesananRepo := repo.NewPemesananRepo(db)

	userService := service.NewUserService(userRepo, sessionRepo)
	makamService := service.NewMakamService(makamRepo)
	jenazahService := service.NewJenazahService(jenazahRepo)
	pemesananService := service.NewPemesananService(pemesananRepo)

	userAPIHandler := api.NewUserAPI(userService)
	makamAPIHandler := api.NewMakamAPI(makamService)
	jenazahAPIHandler := api.NewJenazahAPI(jenazahService)
	pemesananAPIHandler := api.NewPemesananAPI(pemesananService)

	apiHandler := APIHandler{
		UserAPIHandler:      userAPIHandler,
		MakamAPIHandler:     makamAPIHandler,
		JenazahAPIHandler:   jenazahAPIHandler,
		PemesananAPIHandler: pemesananAPIHandler,
	}

	version := gin.Group("/api/v1")
	{
		user := version.Group("/user")
		{
			user.POST("/login", apiHandler.UserAPIHandler.Login)
			user.POST("/register", apiHandler.UserAPIHandler.Register)
			user.Use(middleware.Auth())

			pemesanan := user.Group("/pemesanan")
			{
				
				pemesanan.POST("/add", apiHandler.PemesananAPIHandler.AddPemesanan)
				pemesanan.GET("/get/:id", apiHandler.PemesananAPIHandler.GetPemesananByID)
				pemesanan.PUT("/update/:id", apiHandler.PemesananAPIHandler.UpdatePemesanan)
				pemesanan.DELETE("/delete/:id", apiHandler.PemesananAPIHandler.DeletePemesanan)
				pemesanan.GET("/list", apiHandler.PemesananAPIHandler.GetPemesananList)
			}
		}

		admin := version.Group("/admin")
		{
			admin.Use(middleware.Auth())
			admin.Use(middleware.AuthAdmin())
			admin.POST("/register", apiHandler.UserAPIHandler.Register)

			jenazah := admin.Group("/data-jenazah")
			{
				jenazah.POST("/add", apiHandler.JenazahAPIHandler.AddJenazah)
				jenazah.GET("/get/:id", apiHandler.JenazahAPIHandler.GetJenazahByID)
				jenazah.PUT("/update/:id", apiHandler.JenazahAPIHandler.UpdateJenazah)
				jenazah.DELETE("/delete/:id", apiHandler.JenazahAPIHandler.DeleteJenazah)
				jenazah.GET("/list", apiHandler.JenazahAPIHandler.GetJenazahList)
			}
			makam := admin.Group("/data-makam")
			{
				makam.POST("/add", apiHandler.MakamAPIHandler.AddMakam)
				makam.GET("/get/:id", apiHandler.MakamAPIHandler.GetMakamByID)
				makam.PUT("/update/:id", apiHandler.MakamAPIHandler.UpdateMakam)
				makam.DELETE("/delete/:id", apiHandler.MakamAPIHandler.DeleteMakam)
				makam.GET("/list", apiHandler.MakamAPIHandler.GetMakamList)
			}
		}
	}

	return gin
}

// func RunClient(db *gorm.DB, gin *gin.Engine, embed embed.FS) *gin.Engine {
// 	sessionRepo := repo.NewSessionsRepo(db)
// 	sessionService := service.NewSessionService(sessionRepo)

// 	userClient := client.NewUserClient()

// 	authWeb := web.NewAuthWeb(userClient, sessionService, embed)
// 	modalWeb := web.NewModalWeb(embed)
// 	homeWeb := web.NewHomeWeb(embed)
// 	dashboardWeb := web.NewDashboardWeb(userClient, sessionService, embed)

// 	client := ClientHandler{
// 		authWeb, homeWeb, dashboardWeb, modalWeb,
// 	}

// 	gin.StaticFS("/static", http.Dir("frontend/public"))

// 	gin.GET("/", client.HomeWeb.Index)

// 	user := gin.Group("/client")
// 	{
// 		user.GET("/login", client.AuthWeb.Login)
// 		user.POST("/login/process", client.AuthWeb.LoginProcess)
// 		user.GET("/register", client.AuthWeb.Register)
// 		user.POST("/register/process", client.AuthWeb.RegisterProcess)

// 		user.Use(middleware.Auth())
// 		user.GET("/logout", client.AuthWeb.Logout)
// 	}

// 	main := gin.Group("/client")
// 	{
// 		main.Use(middleware.Auth())
// 		main.GET("/dashboard", client.DashboardWeb.Dashboard)
// 		main.GET("/task", client.TaskWeb.TaskPage)
// 		user.POST("/task/add/process", client.TaskWeb.TaskAddProcess)
// 		user.POST("/category/add/process", client.CategoryWeb.CategoryAddProcess)
// 		main.GET("/category", client.CategoryWeb.CategoryPage)
// 	}

// 	modal := gin.Group("/client")
// 	{
// 		modal.GET("/modal", client.ModalWeb.Modal)
// 	}

// 	return gin
// }
