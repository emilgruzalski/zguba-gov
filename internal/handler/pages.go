package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/emilgruzalski/zguba-gov/internal/model"
	"github.com/emilgruzalski/zguba-gov/internal/municipality"
	"github.com/emilgruzalski/zguba-gov/internal/repository"
	"github.com/gin-gonic/gin"
)

type PagesHandler struct {
	tmpl   *template.Template
	repo   *repository.FoundItemRepo
	munSvc *municipality.Service
}

func NewPagesHandler(templateFS fs.FS, repo *repository.FoundItemRepo, munSvc *municipality.Service) (*PagesHandler, error) {
	funcMap := template.FuncMap{
		"json": func(v any) template.JS {
			b, _ := json.MarshalIndent(v, "", "  ")
			return template.JS(b)
		},
		"seq": func(n int) []int {
			s := make([]int, n)
			for i := range s {
				s[i] = i + 1
			}
			return s
		},
	}

	tmpl, err := template.New("").Funcs(funcMap).ParseFS(templateFS, "templates/*.html", "templates/partials/*.html")
	if err != nil {
		return nil, fmt.Errorf("parse templates: %w", err)
	}

	return &PagesHandler{tmpl: tmpl, repo: repo, munSvc: munSvc}, nil
}

type wizardData struct {
	CurrentStep      int
	TotalSteps       int
	Progress         int
	MunicipalityName string
	MunicipalityType string
	ContactEmail     string
	ItemName         string
	ItemCategory     string
	ItemDate         string
	ItemLocation     string
	ItemStatus       string
	ItemDescription  string
	StorageDeadline  string
	PickupLocation   string
	PickupHours      string
	ContactPerson    string
	Items            []model.FoundItemResponse
	Errors           []string
	Success          string
}

func (h *PagesHandler) Index(c *gin.Context) {
	items, _ := h.repo.List(model.ListParams{Limit: 50})
	var resp []model.FoundItemResponse
	for _, it := range items {
		resp = append(resp, it.ToResponse())
	}

	data := wizardData{
		CurrentStep:     1,
		TotalSteps:      4,
		Progress:        0,
		StorageDeadline: "30",
		Items:           resp,
	}

	h.render(c, "layout.html", data)
}

func (h *PagesHandler) NextStep(c *gin.Context) {
	data := h.parseForm(c)
	step := data.CurrentStep

	if errors := h.validateStep(step, data); len(errors) > 0 {
		data.Errors = errors
		c.Header("HX-Retarget", "#modals")
		c.Header("HX-Reswap", "innerHTML")
		h.renderPartial(c, "error_modal.html", data)
		return
	}

	data.CurrentStep = step + 1
	if data.CurrentStep > data.TotalSteps {
		data.CurrentStep = data.TotalSteps
	}
	data.Progress = (data.CurrentStep - 1) * 100 / data.TotalSteps

	h.renderStep(c, data)
}

func (h *PagesHandler) PrevStep(c *gin.Context) {
	data := h.parseForm(c)
	data.CurrentStep--
	if data.CurrentStep < 1 {
		data.CurrentStep = 1
	}
	data.Progress = (data.CurrentStep - 1) * 100 / data.TotalSteps

	h.renderStep(c, data)
}

func (h *PagesHandler) Autocomplete(c *gin.Context) {
	q := c.Query("q")
	unitType := c.Query("type")
	results := h.munSvc.Search(q, unitType)

	if results == nil {
		results = []municipality.TerritorialUnit{}
	}

	type acResult struct {
		Units []municipality.TerritorialUnit
	}
	h.renderPartial(c, "autocomplete.html", acResult{Units: results})
}

func (h *PagesHandler) Submit(c *gin.Context) {
	data := h.parseForm(c)

	deadline, _ := strconv.Atoi(data.StorageDeadline)

	create := model.FoundItemCreate{
		Municipality: model.MunicipalityInfo{
			Name:         data.MunicipalityName,
			Type:         data.MunicipalityType,
			ContactEmail: data.ContactEmail,
		},
		Item: model.ItemInfo{
			Name:        data.ItemName,
			Category:    data.ItemCategory,
			Date:        data.ItemDate,
			Location:    data.ItemLocation,
			Status:      data.ItemStatus,
			Description: data.ItemDescription,
		},
		Pickup: model.PickupInfo{
			Deadline: deadline,
			Location: data.PickupLocation,
			Hours:    data.PickupHours,
			Contact:  data.ContactPerson,
		},
	}

	if create.Item.Status == "" {
		create.Item.Status = "available"
	}

	_, err := h.repo.Create(create)
	if err != nil {
		data.Errors = []string{"Błąd zapisu: " + err.Error()}
		c.Header("HX-Retarget", "#modals")
		c.Header("HX-Reswap", "innerHTML")
		h.renderPartial(c, "error_modal.html", data)
		return
	}

	items, _ := h.repo.List(model.ListParams{Limit: 50})
	var resp []model.FoundItemResponse
	for _, it := range items {
		resp = append(resp, it.ToResponse())
	}

	data.Items = resp
	data.Success = "Rzecz została pomyślnie zarejestrowana w systemie."
	data.CurrentStep = 1
	data.Progress = 0
	data.MunicipalityName = ""
	data.MunicipalityType = ""
	data.ContactEmail = ""
	data.ItemName = ""
	data.ItemCategory = ""
	data.ItemDate = ""
	data.ItemLocation = ""
	data.ItemStatus = ""
	data.ItemDescription = ""
	data.StorageDeadline = "30"
	data.PickupLocation = ""
	data.PickupHours = ""
	data.ContactPerson = ""

	h.renderPartial(c, "submit_result.html", data)
}

func (h *PagesHandler) ExportJSON(c *gin.Context) {
	data := h.parseQuery(c)
	deadline, _ := strconv.Atoi(data.StorageDeadline)

	export := map[string]any{
		"municipality": map[string]string{
			"name":         data.MunicipalityName,
			"type":         data.MunicipalityType,
			"contactEmail": data.ContactEmail,
		},
		"item": map[string]string{
			"name":        data.ItemName,
			"category":    data.ItemCategory,
			"date":        data.ItemDate,
			"location":    data.ItemLocation,
			"status":      data.ItemStatus,
			"description": data.ItemDescription,
		},
		"pickup": map[string]any{
			"deadline": deadline,
			"location": data.PickupLocation,
			"hours":    data.PickupHours,
			"contact":  data.ContactPerson,
		},
	}

	b, _ := json.MarshalIndent(export, "", "  ")
	filename := fmt.Sprintf("zguba-gov-export-%s.json", time.Now().Format("2006-01-02"))
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "application/json", b)
}

func (h *PagesHandler) ExportCSV(c *gin.Context) {
	data := h.parseQuery(c)

	filename := fmt.Sprintf("zguba-gov-export-%s.csv", time.Now().Format("2006-01-02"))
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "text/csv; charset=utf-8")
	c.Writer.Write([]byte{0xEF, 0xBB, 0xBF}) // UTF-8 BOM

	w := csv.NewWriter(c.Writer)
	w.Write([]string{"Pole", "Wartość"})
	w.Write([]string{"Samorząd", data.MunicipalityName})
	w.Write([]string{"Typ samorządu", data.MunicipalityType})
	w.Write([]string{"Email kontaktowy", data.ContactEmail})
	w.Write([]string{"Nazwa przedmiotu", data.ItemName})
	w.Write([]string{"Kategoria", data.ItemCategory})
	w.Write([]string{"Data znalezienia", data.ItemDate})
	w.Write([]string{"Miejsce znalezienia", data.ItemLocation})
	w.Write([]string{"Status", data.ItemStatus})
	w.Write([]string{"Opis", data.ItemDescription})
	w.Write([]string{"Termin przechowania (dni)", data.StorageDeadline})
	w.Write([]string{"Miejsce odbioru", data.PickupLocation})
	w.Write([]string{"Godziny odbioru", data.PickupHours})
	w.Write([]string{"Osoba kontaktowa", data.ContactPerson})
	w.Flush()
}

func (h *PagesHandler) parseForm(c *gin.Context) wizardData {
	step, _ := strconv.Atoi(c.PostForm("currentStep"))
	if step == 0 {
		step = 1
	}

	items, _ := h.repo.List(model.ListParams{Limit: 50})
	var resp []model.FoundItemResponse
	for _, it := range items {
		resp = append(resp, it.ToResponse())
	}

	return wizardData{
		CurrentStep:      step,
		TotalSteps:       4,
		Progress:         (step - 1) * 100 / 4,
		MunicipalityName: c.PostForm("municipalityName"),
		MunicipalityType: c.PostForm("municipalityType"),
		ContactEmail:     c.PostForm("contactEmail"),
		ItemName:         c.PostForm("itemName"),
		ItemCategory:     c.PostForm("itemCategory"),
		ItemDate:         c.PostForm("itemDate"),
		ItemLocation:     c.PostForm("itemLocation"),
		ItemStatus:       c.PostForm("itemStatus"),
		ItemDescription:  c.PostForm("itemDescription"),
		StorageDeadline:  c.DefaultPostForm("storageDeadline", "30"),
		PickupLocation:   c.PostForm("pickupLocation"),
		PickupHours:      c.PostForm("pickupHours"),
		ContactPerson:    c.PostForm("contactPerson"),
		Items:            resp,
	}
}

func (h *PagesHandler) parseQuery(c *gin.Context) wizardData {
	deadline := c.DefaultQuery("storageDeadline", "30")
	return wizardData{
		MunicipalityName: c.Query("municipalityName"),
		MunicipalityType: c.Query("municipalityType"),
		ContactEmail:     c.Query("contactEmail"),
		ItemName:         c.Query("itemName"),
		ItemCategory:     c.Query("itemCategory"),
		ItemDate:         c.Query("itemDate"),
		ItemLocation:     c.Query("itemLocation"),
		ItemStatus:       c.Query("itemStatus"),
		ItemDescription:  c.Query("itemDescription"),
		StorageDeadline:  deadline,
		PickupLocation:   c.Query("pickupLocation"),
		PickupHours:      c.Query("pickupHours"),
		ContactPerson:    c.Query("contactPerson"),
	}
}

func (h *PagesHandler) validateStep(step int, data wizardData) []string {
	var errors []string
	switch step {
	case 1:
		if data.MunicipalityType == "" {
			errors = append(errors, "Wybierz typ samorządu")
		}
		if data.MunicipalityName == "" {
			errors = append(errors, "Podaj nazwę samorządu")
		}
		if data.ContactEmail == "" || !strings.Contains(data.ContactEmail, "@") {
			errors = append(errors, "Podaj prawidłowy adres email")
		}
	case 2:
		if data.ItemName == "" {
			errors = append(errors, "Podaj nazwę przedmiotu")
		}
		if data.ItemCategory == "" {
			errors = append(errors, "Wybierz kategorię")
		}
		if data.ItemDate == "" {
			errors = append(errors, "Podaj datę znalezienia")
		}
		if data.ItemLocation == "" {
			errors = append(errors, "Podaj miejsce znalezienia")
		}
	case 3:
		deadline, err := strconv.Atoi(data.StorageDeadline)
		if err != nil || deadline < 1 || deadline > 365 {
			errors = append(errors, "Termin przechowania musi być od 1 do 365 dni")
		}
		if data.PickupLocation == "" {
			errors = append(errors, "Podaj miejsce odbioru")
		}
	}
	return errors
}

func (h *PagesHandler) render(c *gin.Context, name string, data any) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.String(http.StatusInternalServerError, "template error: %v", err)
	}
}

func (h *PagesHandler) renderPartial(c *gin.Context, name string, data any) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.String(http.StatusInternalServerError, "template error: %v", err)
	}
}

func (h *PagesHandler) renderStep(c *gin.Context, data wizardData) {
	stepName := fmt.Sprintf("step%d.html", data.CurrentStep)
	c.Header("Content-Type", "text/html; charset=utf-8")

	if err := h.tmpl.ExecuteTemplate(c.Writer, stepName, data); err != nil {
		c.String(http.StatusInternalServerError, "template error: %v", err)
		return
	}
	h.tmpl.ExecuteTemplate(c.Writer, "progress_oob.html", data)
}
