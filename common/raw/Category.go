package raw

type Category struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Type       int    `json:"type"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type CategoryFiltertype struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Sort       int    `json:"sort"`
	CreateTime string `json:"create_time"`
	Categoryid int64  `json:"categoryid"`
}

type CategoryFilteritem struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Sort      int    `json:"sort"`
	Desc      string `json:"desc"`
	Condition string `json:"condition"`
	Filterid  int64  `json:"filterid"`
}
