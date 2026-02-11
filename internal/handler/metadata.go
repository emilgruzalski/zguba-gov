package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type MetadataHandler struct{}

func NewMetadataHandler() *MetadataHandler {
	return &MetadataHandler{}
}

func (h *MetadataHandler) Catalog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"@context":        "https://www.w3.org/ns/dcat",
		"@type":           "dcat:Catalog",
		"dct:title":       "Katalog Rzeczy Znalezionych - Zguba.gov",
		"dct:description": "System zarządzania znalezionymi przedmiotami w jednostkach administracji publicznej w Polsce",
		"dct:issued":      "2025-12-01T00:00:00Z",
		"dct:modified":    time.Now().UTC().Format(time.RFC3339),
		"dct:language":    "pl",
		"dcat:homepage":   "https://zguba.gov",
		"dcat:dataset": []gin.H{
			{
				"@type":                  "dcat:Dataset",
				"dct:identifier":         "found-items-dataset",
				"dct:title":              "Znalezione Rzeczy",
				"dct:description":        "Bieżąca lista znalezionych przedmiotów zgłoszonych przez jednostki administracji publicznej",
				"dct:license":            "http://creativecommons.org/licenses/by/4.0/",
				"dcat:keyword":           []string{"rzeczy znalezione", "przedmioty", "administracja publiczna", "dane otwarte"},
				"dcat:theme":             []string{"http://publications.europa.eu/resource/authority/data-theme/SOCI"},
				"dct:accrualPeriodicity": "http://purl.org/ckan/freq/daily",
				"dcat:distribution": []gin.H{
					{
						"@type":            "dcat:Distribution",
						"dct:title":        "JSON API",
						"dcat:accessURL":   "https://api.zguba.gov/api/found-items",
						"dcat:downloadURL": "https://api.zguba.gov/api/found-items?limit=1000",
						"dct:format":       "JSON",
						"dcat:mediaType":   "application/json",
					},
					{
						"@type":          "dcat:Distribution",
						"dct:title":      "OData API",
						"dcat:accessURL": "https://api.zguba.gov/odata",
						"dct:format":     "OData",
					},
				},
				"dcat:contactPoint": gin.H{
					"@type":          "vcard:Organization",
					"vcard:fn":       "Zguba.gov Support",
					"vcard:hasEmail": "mailto:support@zguba.gov",
				},
			},
		},
		"dcat:organization": gin.H{
			"@type":     "foaf:Organization",
			"foaf:name": "Zguba.gov",
		},
	})
}

func (h *MetadataHandler) Distribution(c *gin.Context) {
	id := c.Param("id")

	distributions := map[string]gin.H{
		"json-api": {
			"@type":          "dcat:Distribution",
			"dct:title":      "JSON API",
			"dcat:accessURL": "https://api.zguba.gov/api/found-items",
			"dct:format":     "JSON",
			"dcat:mediaType": "application/json",
			"dct:license":    "http://creativecommons.org/licenses/by/4.0/",
		},
		"odata": {
			"@type":          "dcat:Distribution",
			"dct:title":      "OData API",
			"dcat:accessURL": "https://api.zguba.gov/odata",
			"dct:format":     "OData",
			"dct:license":    "http://creativecommons.org/licenses/by/4.0/",
		},
	}

	dist, ok := distributions[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Distribution not found"})
		return
	}
	c.JSON(http.StatusOK, dist)
}
