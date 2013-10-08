package addons

import "net/url"
import "strings"

// PgBackups Addon

type PgBackup struct {
	Id          int64
	UserId      int64  `json:"user_id"`
	FromUrl     string `json:"from_url"`
	FromName    string `json:"from_name"`
	ToUrl       string `json:"to_url"`
	ToName      string `json:"to_name"`
	CreatedAt   string `json:"created_at"`
	StartedAt   string `json:"started_at"`
	UpdatedAt   string `json:"updated_at"`
	FinishedAt  string `json:"finished_at"`
	ErrorAt     string `json:"error_at"`
	DestroyedAt string `json:"destroyed_at"`
	Progress    string
	Size        string
	Expire      bool
	CleanedAt   string `json:"cleaned_at"`
	User        string
	Type        string
	Duration    float64
	Log         string `json:-`
	PublicUrl   string `json:"public_url"`
}

const pgBackupNameExt = ".dump"

func (backup PgBackup) Name() string {
	toUrl, err := url.Parse(backup.ToUrl)
	if err != nil {
		panic(err)
	}
	components := strings.Split(toUrl.Path, "/")
	if len(components) <= 0 {
		return ""
	}
	lastPart := components[len(components)-1]
	return strings.TrimSuffix(lastPart, pgBackupNameExt)
}
