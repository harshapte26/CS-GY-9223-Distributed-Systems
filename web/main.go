package main

type Server struct{
	router *gin.Engine
	dbClient *db.Client
	tables GlobalTables
}

