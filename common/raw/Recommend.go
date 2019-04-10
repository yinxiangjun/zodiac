package raw

type RecommendList struct {
	Id           int
	Name         string `orm:"size(20)"`
	Title        string `orm:"size(20)"`
	Desc         string `orm:"size(100)"`
	Status       int
	CategoryId   int
	CategoryName string `orm:"size(20)"`
}

type RecommendVideo struct {
	Id           int
	ListId       int
	VideoId      int64
	StatusInList int
	OrderInList  int
	CustomImg    string
}