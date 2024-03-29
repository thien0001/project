package models

// TemplateData holds data sent from handlers package
type TemplateData struct {
	StringMap map[string]string
	InMap     map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
