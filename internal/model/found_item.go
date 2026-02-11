package model

import (
	"database/sql"
	"encoding/json"
	"time"
)

type FoundItem struct {
	ID                string
	MunicipalityName  string
	MunicipalityType  string
	MunicipalityEmail string
	ItemName          string
	ItemCategory      string
	ItemDate          string
	ItemLocation      string
	ItemStatus        string
	ItemDescription   sql.NullString
	PickupDeadline    int
	PickupLocation    string
	PickupHours       sql.NullString
	PickupContact     sql.NullString
	Categories        sql.NullString
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type MunicipalityInfo struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	ContactEmail string `json:"contactEmail"`
}

type ItemInfo struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Date        string `json:"date"`
	Location    string `json:"location"`
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
}

type PickupInfo struct {
	Deadline int    `json:"deadline"`
	Location string `json:"location"`
	Hours    string `json:"hours,omitempty"`
	Contact  string `json:"contact,omitempty"`
}

type FoundItemCreate struct {
	Municipality MunicipalityInfo `json:"municipality"`
	Item         ItemInfo         `json:"item"`
	Pickup       PickupInfo       `json:"pickup"`
	Categories   []string         `json:"categories,omitempty"`
}

type FoundItemUpdate struct {
	Municipality *MunicipalityInfo `json:"municipality,omitempty"`
	Item         *ItemInfo         `json:"item,omitempty"`
	Pickup       *PickupInfo       `json:"pickup,omitempty"`
	Categories   *[]string         `json:"categories,omitempty"`
}

type FoundItemResponse struct {
	ID           string           `json:"id"`
	Municipality MunicipalityInfo `json:"municipality"`
	Item         ItemInfo         `json:"item"`
	Pickup       PickupInfo       `json:"pickup"`
	Categories   []string         `json:"categories"`
	CreatedAt    string           `json:"createdAt,omitempty"`
	UpdatedAt    string           `json:"updatedAt,omitempty"`
}

func (fi *FoundItem) ToResponse() FoundItemResponse {
	var cats []string
	if fi.Categories.Valid && fi.Categories.String != "" {
		_ = json.Unmarshal([]byte(fi.Categories.String), &cats)
	}
	if cats == nil {
		cats = []string{}
	}

	return FoundItemResponse{
		ID: fi.ID,
		Municipality: MunicipalityInfo{
			Name:         fi.MunicipalityName,
			Type:         fi.MunicipalityType,
			ContactEmail: fi.MunicipalityEmail,
		},
		Item: ItemInfo{
			Name:        fi.ItemName,
			Category:    fi.ItemCategory,
			Date:        fi.ItemDate,
			Location:    fi.ItemLocation,
			Status:      fi.ItemStatus,
			Description: fi.ItemDescription.String,
		},
		Pickup: PickupInfo{
			Deadline: fi.PickupDeadline,
			Location: fi.PickupLocation,
			Hours:    fi.PickupHours.String,
			Contact:  fi.PickupContact.String,
		},
		Categories: cats,
		CreatedAt:  fi.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  fi.UpdatedAt.Format(time.RFC3339),
	}
}

type StatsResponse struct {
	FoundItems        FoundItemStats      `json:"foundItems"`
	TopCategories     []CategoryCount     `json:"topCategories"`
	TopMunicipalities []MunicipalityCount `json:"topMunicipalities"`
}

type FoundItemStats struct {
	Total     int `json:"total"`
	Available int `json:"available"`
	Claimed   int `json:"claimed"`
}

type CategoryCount struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

type MunicipalityCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type ListParams struct {
	Skip         int
	Limit        int
	Category     string
	Municipality string
	Status       string
	Search       string
}
