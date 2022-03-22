package models

// Image is a docker image to run a workspace container.
type Image struct {
	Id           int64    `db:"id" json:"id"`
	Name         string   `db:"name" json:"name"`                 // required
	Title        string   `db:"title" json:"title"`               // optional
	Description  string   `db:"description" json:"description"`   // optional
	Environments []EnvVar `db:"environments" json:"environments"` // optional
}
