package model

type User struct {
	Uid      int64   `gorm:"AUTO_INCREMENT;primary_key"`
	Username string  `gorm:"not null;unique;unique_index"`
	Money    float32 `sql:"type:decimal(10,2);"`
}
