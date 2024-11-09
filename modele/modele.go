package modele

type Endpoint struct {
	Tag         string
	Root        string
	Method      string
	Body        string
	Response    string
	Status_code string
	Description string
	Summary     string
}

type Lazy_doc struct {
	Name        string
	Version     string
	Description string
	Endpoints   []Endpoint
}
