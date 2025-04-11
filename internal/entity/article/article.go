package article

import "time"

type Posts struct {
	ID          uint      `gorm:"column:Id;primaryKey;autoIncrement" json:"id,omitempty"`
	Title       string    `gorm:"column:Title;type:varchar(200);not null"`
	Content     string    `gorm:"column:Content;type:text;not null"`
	Category    string    `gorm:"column:Category;type:varchar(100);not null"`
	CreatedDate time.Time `gorm:"column:Created_date;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_date,omitempty"`
	UpdatedDate time.Time `gorm:"column:Updated_date;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_date,omitempty"`
	Status      string    `gorm:"column:Status;type:varchar(100);not null;check:status IN ('Publish', 'Draft', 'Trash')"`
}

type PostsResponse struct {
	ID       uint   `gorm:"column:Id" json:"id,omitempty"`
	Title    string `gorm:"column:Title" json:"title,omitempty"`
	Content  string `gorm:"column:Content" json:"content,omitempty"`
	Category string `gorm:"column:Category" json:"category,omitempty"`
	Status   string `gorm:"column:Status" json:"status,omitempty"`
}
