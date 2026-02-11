package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kacperfilipiuk/zguba-gov/internal/odata"
	"github.com/kacperfilipiuk/zguba-gov/internal/repository"
)

type ODataHandler struct {
	repo *repository.FoundItemRepo
}

func NewODataHandler(repo *repository.FoundItemRepo) *ODataHandler {
	return &ODataHandler{repo: repo}
}

func (h *ODataHandler) Query(c *gin.Context) {
	skip, _ := strconv.Atoi(c.DefaultQuery("$skip", "0"))
	top, _ := strconv.Atoi(c.DefaultQuery("$top", "50"))
	if top > 100 {
		top = 100
	}

	filter := c.Query("$filter")
	orderby := c.Query("$orderby")
	countParam, _ := strconv.ParseBool(c.DefaultQuery("$count", "false"))

	filterClause, err := odata.ParseFilter(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderClause := odata.ParseOrderBy(orderby)

	query := "SELECT id, municipality_name, municipality_type, municipality_email, item_name, item_category, item_date, item_location, item_status, item_description, pickup_deadline, pickup_location, pickup_hours, pickup_contact, categories, created_at, updated_at FROM found_items"
	var args []any

	if filterClause.Where != "" {
		query += " WHERE " + filterClause.Where
		args = append(args, filterClause.Args...)
	}

	query += " ORDER BY " + orderClause
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", top, skip)

	items, err := h.repo.QueryRaw(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	value := make([]gin.H, 0, len(items))
	for _, item := range items {
		resp := item.ToResponse()
		value = append(value, gin.H{
			"id":                 resp.ID,
			"municipality_name":  resp.Municipality.Name,
			"municipality_type":  resp.Municipality.Type,
			"municipality_email": resp.Municipality.ContactEmail,
			"item_name":          resp.Item.Name,
			"item_category":      resp.Item.Category,
			"item_date":          resp.Item.Date,
			"item_location":      resp.Item.Location,
			"item_status":        resp.Item.Status,
			"item_description":   resp.Item.Description,
			"pickup_deadline":    resp.Pickup.Deadline,
			"pickup_location":    resp.Pickup.Location,
			"pickup_hours":       resp.Pickup.Hours,
			"pickup_contact":     resp.Pickup.Contact,
			"categories":         resp.Categories,
			"created_at":         resp.CreatedAt,
			"updated_at":         resp.UpdatedAt,
		})
	}

	result := gin.H{
		"odata.context": "https://api.zguba.gov/odata/$metadata",
		"value":         value,
	}

	if countParam {
		countWhere := ""
		var countArgs []any
		if filterClause.Where != "" {
			countWhere = filterClause.Where
			countArgs = filterClause.Args
		}
		count, _ := h.repo.Count(countWhere, countArgs...)
		result["odata.count"] = count
	}

	c.JSON(http.StatusOK, result)
}

func (h *ODataHandler) Metadata(c *gin.Context) {
	c.Header("Content-Type", "application/xml")
	c.String(http.StatusOK, odataMetadataXML)
}

const odataMetadataXML = `<?xml version="1.0" encoding="utf-8"?>
<edmx:Edmx Version="4.0" xmlns:edmx="http://docs.oasis-open.org/odata/ns/edmx">
  <edmx:DataServices>
    <Schema Namespace="ZgubaGov" xmlns="http://docs.oasis-open.org/odata/ns/edm">
      <EntityType Name="FoundItem">
        <Key><PropertyRef Name="id"/></Key>
        <Property Name="id" Type="Edm.String" Nullable="false"/>
        <Property Name="municipality_name" Type="Edm.String"/>
        <Property Name="municipality_type" Type="Edm.String"/>
        <Property Name="municipality_email" Type="Edm.String"/>
        <Property Name="item_name" Type="Edm.String"/>
        <Property Name="item_category" Type="Edm.String"/>
        <Property Name="item_date" Type="Edm.String"/>
        <Property Name="item_location" Type="Edm.String"/>
        <Property Name="item_status" Type="Edm.String"/>
        <Property Name="item_description" Type="Edm.String"/>
        <Property Name="pickup_deadline" Type="Edm.Int32"/>
        <Property Name="pickup_location" Type="Edm.String"/>
        <Property Name="pickup_hours" Type="Edm.String"/>
        <Property Name="pickup_contact" Type="Edm.String"/>
        <Property Name="categories" Type="Collection(Edm.String)"/>
        <Property Name="created_at" Type="Edm.DateTimeOffset"/>
        <Property Name="updated_at" Type="Edm.DateTimeOffset"/>
      </EntityType>
      <EntityContainer Name="Default">
        <EntitySet Name="FoundItems" EntityType="ZgubaGov.FoundItem"/>
      </EntityContainer>
    </Schema>
  </edmx:DataServices>
</edmx:Edmx>`
