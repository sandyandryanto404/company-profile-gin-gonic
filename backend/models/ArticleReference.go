package models

type ArticleReference struct {
	ArticleId   uint64 `gorm:"primary_key;auto_increment:false"`
	ReferenceId uint64 `gorm:"primary_key;auto_increment:false"`
}

func (ArticleReference) TableName() string {
	return "articles_references"
}
