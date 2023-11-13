package model

type Profile struct {
	Url          string   `bson:"url"`
	Id           string   `bson:"_id"`
	Title        string   `bson:"title"`
	ReleaseTime  string   `bson:"releaseTime"`
	FetchTime    string   `bson:"fetchTime"`
	Province     string   `bson:"province"`
	City         string   `bson:"city"`
	Age          string   `bson:"age"`
	FaceValue    string   `bson:"faceValue"`
	Expenses     string   `bson:"expenses"`
	Services     string   `bson:"services"`
	QQ           string   `bson:"qq"`
	Wechat       string   `bson:"wechat"`
	Telegram     string   `bson:"telegram"`
	Tel          string   `bson:"tel"`
	Address      string   `bson:"address"`
	Pics         []string `bson:"pics"`
	Introduction string   `bson:"introduction"`
}
