package entity

import "time"

type File struct {
	FileID    string    `json:"file_id"`
	Name      string    `json:"file_name"`
	Size      int       `json:"file_size"`
	MimeType  string    `json:"file_mime_type"`
	Content   []byte    `json:"file_content"`
	Extension string    `json:"file_extension"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
