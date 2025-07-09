package scans

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marrcoribeiro/security-scanner-api/internal/api/v1/scans/dto"
	"github.com/marrcoribeiro/security-scanner-api/internal/domain"
)

type Controller struct {
	analyzers  []domain.Analyzer
	scanRunner domain.ScanRunner
}

func NewController(analyzers []domain.Analyzer, scanRunner domain.ScanRunner) *Controller {
	return &Controller{
		analyzers:  analyzers,
		scanRunner: scanRunner,
	}
}

func (c *Controller) MountRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("v1/scans", c.ScanHandler)
}

// ScanHandler handles POST /v1/scans
// @Summary      Start a security scan
// @Description  Starts a security scan on the given path using the configured analyzers.
// @Tags         scans
// @Accept       json
// @Produce      json
// @Param        scan  body  dto.CreateScanRequest  true  "Scan configuration"
// @Success      200   {object}  dto.ScanResponse
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /v1/scans [post]
func (c *Controller) ScanHandler(ctx *gin.Context) {
	var req dto.CreateScanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
			"details": err.Error(),
		})
		return
	}

	scan := req.ToDomain()

	c.scanRunner.RunScan(scan, c.analyzers)
	if scan.Err != "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": scan.Err})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToResponse(scan))
}
