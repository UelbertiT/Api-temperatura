package models

type Teste struct {
	ID    uint `gorm: "primaryKey; autoInccrement:true"`
	Texto string
}
