package fruit

// Fruit 🍉
type Fruit struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
	View int64  `json:"view"`
}
