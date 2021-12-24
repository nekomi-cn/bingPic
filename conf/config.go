package conf

// 配置文件定义
type BingImagesDoc struct {
	Images   []ImageDoc  `json:"images" yaml:"images"`
	Tooltips interface{} `json:"tooltips" yaml:"tooltips"`
}

type ImageDoc struct {
	Startdate     string        `json:"startdate" yaml:"startdate"`
	Fullstartdate string        `json:"fullstartdate" yaml:"fullstartdate"`
	Enddate       string        `json:"enddate" yaml:"enddate"`
	Url           string        `json:"url" yaml:"url"`
	Urlbase       string        `json:"urlbase" yaml:"urlbase"`
	Copyright     string        `json:"copyright" yaml:"copyright"`
	Copyrightlink string        `json:"copyrightlink" yaml:"copyrightlink"`
	Title         string        `json:"title" yaml:"title"`
	Quiz          string        `json:"quiz" yaml:"quiz"`
	Wp            bool          `json:"wp" yaml:"wp"`
	Hsh           string        `json:"hsh" yaml:"hsh"`
	Drk           int           `json:"drk" yaml:"drk"`
	Top           int           `json:"top" yaml:"top"`
	Bot           int           `json:"bot" yaml:"bot"`
	Hs            []interface{} `json:"hs" yaml:"hs"`
}
