package main

type Server struct{
	router *gin.Engine
	dbClient *db.Client
	tables GlobalTables
}

func (s *Server) NewRouter() *gin.Engine{
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", middleware.TokenHeader},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
}