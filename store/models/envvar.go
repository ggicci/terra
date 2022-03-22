package models

type EnvVar struct {
	Key      string `db:"key" json:"key"`
	Value    string `db:"value" json:"value"`
	Required bool   `db:"required" json:"required"`
	Default  string `db:"default" json:"default"`
	Document string `db:"document" json:"document"` // support markdown
}
