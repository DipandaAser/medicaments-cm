package handlers

import (
	"github.com/gin-gonic/gin"
	"medicaments-cm/pkg/scrapper"
	"net/http"
)

func ListMedicaments(ctx *gin.Context) {
	keywords := ctx.Query("search")

	medicaments, err := scrapper.SearchMedicaments(keywords)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, medicaments)
}
