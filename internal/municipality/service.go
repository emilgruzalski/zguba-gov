package municipality

import (
	"embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

//go:embed territorial-units.json
var dataFS embed.FS

type FlexString string

func (f *FlexString) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		*f = FlexString(s)
		return nil
	}
	var n json.Number
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	*f = FlexString(fmt.Sprintf("%07s", n.String()))
	return nil
}

type TerritorialUnit struct {
	ID          FlexString `json:"id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Email       string     `json:"email,omitempty"`
	OfficeName  string     `json:"officeName,omitempty"`
	Voivodeship string     `json:"voivodeship,omitempty"`
	County      string     `json:"county,omitempty"`
}

type searchEntry struct {
	unit       TerritorialUnit
	normalized string
}

type Service struct {
	units   []TerritorialUnit
	indexed []searchEntry
}

func NewService() (*Service, error) {
	data, err := dataFS.ReadFile("territorial-units.json")
	if err != nil {
		return nil, err
	}

	var units []TerritorialUnit
	if err := json.Unmarshal(data, &units); err != nil {
		return nil, err
	}

	indexed := make([]searchEntry, len(units))
	for i, u := range units {
		indexed[i] = searchEntry{
			unit:       u,
			normalized: normalize(u.Name),
		}
	}

	return &Service{units: units, indexed: indexed}, nil
}

func (s *Service) Search(query, unitType string) []TerritorialUnit {
	if len(query) < 2 {
		return nil
	}

	q := normalize(query)

	type scored struct {
		unit  TerritorialUnit
		score int
	}

	var results []scored
	for _, e := range s.indexed {
		if unitType != "" && e.unit.Type != unitType {
			continue
		}
		idx := strings.Index(e.normalized, q)
		if idx < 0 {
			continue
		}
		// Score: exact prefix match = 0, contains = position
		results = append(results, scored{unit: e.unit, score: idx})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].score < results[j].score
	})

	limit := 20
	if len(results) < limit {
		limit = len(results)
	}

	out := make([]TerritorialUnit, limit)
	for i := 0; i < limit; i++ {
		out[i] = results[i].unit
	}
	return out
}

func (s *Service) GenerateEmail(unit TerritorialUnit) string {
	if unit.Email != "" {
		return unit.Email
	}

	name := strings.ToLower(normalize(unit.Name))
	// Remove type prefix
	for _, prefix := range []string{"powiat ", "gmina ", "miasto "} {
		name = strings.TrimPrefix(name, prefix)
	}
	name = strings.ReplaceAll(name, " ", "")

	switch unit.Type {
	case "wojewodztwo":
		return "kontakt@" + name + ".uw.gov.pl"
	case "powiat":
		return "starostwo@" + name + ".pl"
	case "miasto":
		return "urzad@um." + name + ".pl"
	default:
		return "ug@" + name + ".pl"
	}
}

var polishReplacements = map[rune]rune{
	'ą': 'a', 'ć': 'c', 'ę': 'e', 'ł': 'l', 'ń': 'n',
	'ó': 'o', 'ś': 's', 'ź': 'z', 'ż': 'z',
	'Ą': 'a', 'Ć': 'c', 'Ę': 'e', 'Ł': 'l', 'Ń': 'n',
	'Ó': 'o', 'Ś': 's', 'Ź': 'z', 'Ż': 'z',
}

func normalize(s string) string {
	s = norm.NFC.String(s)
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if repl, ok := polishReplacements[r]; ok {
			b.WriteRune(repl)
		} else {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}
