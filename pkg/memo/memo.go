package memo

type Memo struct {
	Id   int    `gorm:"primaryKey" ` // id
	Name string `validate:"min=1"`   //name
	Date string `validate:"len=10"`
	Shop string
	Page string
	Play string
	Talk string
}
