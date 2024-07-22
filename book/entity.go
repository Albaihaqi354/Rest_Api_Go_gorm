package book

type Book struct {
	Id          int    `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(100)"`
	Price       int    `gorm:"type:int"`
	Rating      int    `gorm:"type:int"`
}
