package entity

import "time"

type File struct {
	FileID    int64     `json:"file_id"`
	Name      string    `json:"file_name"`
	Path      string    `json:"file_path"`
	Format    string    `json:"format"`
	FileData  byte      `json:"file_data"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
