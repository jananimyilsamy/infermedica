package infermedica

import (
	"encoding/json"
	"net/http"
)

type SymptomRes struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	CommonName     string         `json:"common_name"`
	Category       string         `json:"category"`
	Seriousness    string         `json:"seriousness"`
	Children       []SymptomChild `json:"children"`
	ImageURL       string         `json:"image_url"`
	ImageSource    string         `json:"image_source"`
	ParentID       string         `json:"parent_id"`
	ParentRelation string         `json:"parent_relation"`
}

type SymptomChild struct {
	ID             string `json:"id"`
	ParentRelation string `json:"parent_relation"`
}

func (a *app) Symptoms() (*[]SymptomRes, error) {
	req, err := a.prepareRequest("GET", "symptoms", nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	r := []SymptomRes{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *app) SymptomByID(id string) (*SymptomRes, error) {
	req, err := a.prepareRequest("GET", "symptoms/"+id, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	r := SymptomRes{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}