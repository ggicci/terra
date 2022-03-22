package models

import "time"

type (
	WorkspaceState      string
	WorkspaceVisibility string
)

const (
	WorkspaceStateCreated    WorkspaceState = "created"
	WorkspaceStateUpgrading  WorkspaceState = "upgrading"
	WorkspaceStateStarting   WorkspaceState = "starting"
	WorkspaceStateError      WorkspaceState = "error" // any failure
	WorkspaceStateRunning    WorkspaceState = "running"
	WorkspaceStateStopping   WorkspaceState = "stopping"
	WorkspaceStateStopped    WorkspaceState = "stopped"
	WorkspaceStateDestroying WorkspaceState = "destroying"
	WorkspaceStateDestroyed  WorkspaceState = "destroyed"
)

const (
	WorkspaceVisibilityPublic  WorkspaceVisibility = "public"
	WorkspaceVisibilityPrivate WorkspaceVisibility = "private"
)

// Workspace
type Workspace struct {
	Id         int64               `db:"id" json:"id"`
	CreatedAt  time.Time           `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time           `db:"updated_at" json:"updated_at"`
	Owner      string              `db:"owner" json:"owner"` // User.Login
	Repo       string              `db:"repo" json:"repo"`   // immutable
	Slug       string              `db:"slug" json:"slug"`
	State      WorkspaceState      `db:"state" json:"state"`
	Visibility WorkspaceVisibility `db:"visibility" json:"visibility"`
	ConfigPath string              `db:"config_path" json:"config_path"`
	SnapshotId int64               `db:"snapshot_id" json:"snapshot_id"`

	ImageDetails *Image             `db:"-" json:"image_details"` // lazy, by Image
	Snapshot     *WorkspaceSnapshot `db:"-" json:"snapshot"`      // lazy, by SnapshotId
}

type WorkspaceSnapshot struct {
	Id          int64     `db:"id" json:"id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	UpgradedAt  time.Time `db:"upgraded_at" json:"upgraded_at"`
	WorkspaceId int64     `db:"workspace_id" json:"workspace_id"`

	// Version is an abbrev formatted representation of Raw.Sha.
	Version string `db:"version" json:"version"`

	// Raw is an object populated by consulting GitHub's file API. IMMUTABLE
	Raw WorkspaceConfigFile `db:"raw" json:"raw"`

	// Config is populated by Raw.Content. aka. "Preset Settings".
	Config WorkspaceConfig `db:"config" json:"config"`

	// Runtime can be configured on our WebUI, or through API. aka. "User Settings".
	Runtime WorkspaceConfig `db:"runtime" json:"runtime"`
}

type WorkspaceConfigFile struct {
	Name        string `db:"name" json:"name"`
	Path        string `db:"path" json:"path"`
	Sha         string `db:"sha" json:"sha"`
	Url         string `db:"url" json:"url"`
	HtmlUrl     string `db:"html_url" json:"html_url"`
	GitUrl      string `db:"git_url" json:"git_url"`
	DownloadUrl string `db:"download_url" json:"download_url"`
	Content     string `db:"content" json:"content"`
}

type WorkspaceConfig struct {
	Spec         string   `db:"spec" json:"spec"` // version of the spec, e.g. "1.0"
	Image        string   `db:"image" json:"image"`
	Title        string   `db:"title" json:"title"`
	Description  string   `db:"description" json:"description"`
	Environments []EnvVar `db:"environments" json:"environments"`
}
