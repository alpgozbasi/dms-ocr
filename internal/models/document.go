package models

import "time"

type Document struct {
	ID        int64     `db:"id"`         // primary key
	FileName  string    `db:"file_name"`  // original uploaded file name
	FilePath  string    `db:"file_path"`  // path
	OCRText   string    `db:"ocr_text"`   // output text
	CreatedAt time.Time `db:"created_at"` // creation time
	UpdatedAt time.Time `db:"updated_at"` // last update time
}
