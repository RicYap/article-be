package article

import "time"

type Posts struct {
	ID          uint      `gorm:"column:Id;primaryKey;autoIncrement"`
	Title       string    `gorm:"column:Title;type:varchar(200);not null"`
	Content     string    `gorm:"column:Content;type:text;not null"`
	Category    string    `gorm:"column:Category;type:varchar(100);not null"`
	CreatedDate time.Time `gorm:"column:Created_date;type:timestamp"`
	UpdatedDate time.Time `gorm:"column:Updated_date;type:datetime"`
	Status      string    `gorm:"column:Status;type:varchar(100);not null"`
}

type PostsSlim struct {
	ID       int   `gorm:"column:Id" json:"id,omitempty"`
	Title    string `gorm:"column:Title" json:"title,omitempty"`
	Content  string `gorm:"column:Content" json:"content,omitempty"`
	Category string `gorm:"column:Category" json:"category,omitempty"`
	Status   string `gorm:"column:Status" json:"status,omitempty"`
}
