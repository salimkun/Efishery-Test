package model

type Resource struct {
	UID       string `json:"uuid"`
	Commodity string `json:"komoditas"`
	AreaProv  string `json:"area_provinsi"`
	AreaCity  string `json:"area_kota"`
	Size      string `json:"size"`
	Price     string `json:"price"`
	DateParse string `json:"tgl_parsed"`
	Timestamp string `json:"timestamp"`
}

type AgregateObj struct {
	Mininum   int64     `json:"min"`
	Maximum   int64     `json:"max"`
	Median    float64   `json:"median"`
	Avg       float64   `json:"avg"`
	ArrayData []float64 `json:"array_data"`
}

type MaximumByProv struct {
	Value int64  `json:"value"`
	Prov  string `json:"prov"`
}

type AgregateResource struct {
	Price        AgregateObj `json:"price"`
	Size         AgregateObj `json:"size"`
	AreaProv     string      `json:"area_provinsi"`
	DateResource []string    `json:"dates"`
}

type AgregateResp struct {
	Result []AgregateResource `json:"result"`
}

type ConvertCurrency struct {
	Result float64 `json:"result"`
}

type GroupingResource struct {
	Year     int32
	Mount    string
	Week     int32
	Provincy string
}
