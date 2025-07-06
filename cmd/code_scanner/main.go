package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/marrcoribeiro/security-scanner-api/internal/analyzers"
	"github.com/marrcoribeiro/security-scanner-api/internal/api/v1/scans"
	"github.com/marrcoribeiro/security-scanner-api/internal/domain"
	"github.com/marrcoribeiro/security-scanner-api/internal/useCases"
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

	router.Run(":8080")

	fmt.Println("Security Code Scanner is running on port 8080...")
}