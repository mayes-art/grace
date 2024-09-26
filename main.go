package main

import (
	"grace/database"
	"grace/routers"
	"os"
)

func main() {
	// 連接數據庫
	database.Connect()

	// 自動遷移數據庫
	// if err := database.GetDB().AutoMigrate(&model.User{}); err != nil {
	// 	log.Fatal("failed to migrate database:", err)
	// }
	// database.Migrate()

	// 設置路由
	router := routers.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 運行服務器
	router.Run(":" + port)
}
