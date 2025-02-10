package entities

type Entity struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (e Entity) TableName() string {
	return "entities"
}
