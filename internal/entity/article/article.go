package article

import "time"

type Posts struct {
	ID          uint      `gorm:"column:Id;primaryKey;autoIncrement"`
	Title       string    `gorm:"column:Title;type:varchar(200);not null"`
	Content     string    `gorm:"column:Content;type:text;not null"`
	Category    string    `gorm:"column:Category;type:varchar(100);not null"`
	CreatedDate time.Time `gorm:"column:Created_date;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedDate time.Time `gorm:"column:Updated_date;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Status      string    `gorm:"column:Status;type:varchar(100);not null;check:status IN ('Publish', 'Draft', 'Trash')"`
}

type PostsResponse struct {
	Title       string    `gorm:"column:Title;type:varchar(200);not null"`
	Content     string    `gorm:"column:Content;type:text;not null"`
	Category    string    `gorm:"column:Category;type:varchar(100);not null"`
	Status      string    `gorm:"column:Status;type:varchar(100);not null;check:status IN ('Publish', 'Draft', 'Trash')"`
}
