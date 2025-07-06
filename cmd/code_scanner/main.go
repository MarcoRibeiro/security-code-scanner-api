// @title Security Code Scanner API
// @version 1.0
// @description API for running security code scans.
// @host localhost:8080
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/marrcoribeiro/security-scanner-api/docs"
	"github.com/marrcoribeiro/security-scanner-api/internal/analyzers"
	"github.com/marrcoribeiro/security-scanner-api/internal/api/v1/scans"
	"github.com/marrcoribeiro/security-scanner-api/internal/domain"
	"github.com/marrcoribeiro/security-scanner-api/internal/useCases"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	fmt.Println("Starting the Security Code Scanner...")

	router := gin.Default()

	analyzersList := []domain.Analyzer{
		analyzers.NewSQLInjectionAnalyzer(),
		analyzers.NewCrossSiteScriptingAnalyzer(),
	}

	scanRunner := useCases.NewScanRunner()

	scansController := scans.NewController(analyzersList, scanRunner)

	scansController.MountRoutes(&router.RouterGroup)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")

	fmt.Println("Security Code Scanner is running on port 8080...")
}