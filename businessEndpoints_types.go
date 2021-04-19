package yelp

type SpecialHours struct {
	Date        string `json:"date"`
	IsClosed    bool   `json:"is_closed"`
	Start       string `json:"start"`
	End         string `json:"end"`
	IsOvernight bool   `json:"is_overnight"`
}

type Open struct {
	Day         int    `json:"day"`
	Start       string `json:"start"`
	End         string `json:"end"`
	IsOvernight bool   `json:"is_overnight"`
}

type Hours struct {
	IsOpenNow bool   `json:"is_open_now"`
	HoursType string `json:"hours_type"`
	Open      []Open `json:"open"`
}

type Messaging struct {
	Url         string `json:"url"`
	UseCaseText string `json:"use_case_text"`
}

type Coordinates struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type CategoryId struct {
	Title string `json:"title"`
	Alias string `json:"alias"`
}

type Business struct {
	Rating       float32      `json:"rating"`
	Price        string       `json:"price"`
	DisplayPhone string       `json:"display_phone"`
	Phone        string       `json:"phone"`
	Id           string       `json:"id"`
	Alias        string       `json:"alias"`
	IsClosed     bool         `json:"is_closed"`
	IsClaimed    bool         `json:"is_claimed"`
	ReviewCount  int          `json:"review_count"`
	Name         string       `json:"name"`
	Url          string       `json:"url"`
	ImageUrl     string       `json:"image_url"`
	Distance     float32      `json:"distance"`
	Transactions []string     `json:"transactions"`
	Location     Location     `json:"location"`
	Coordinates  Coordinates  `json:"coordinates"`
	Categories   []CategoryId `json:"categories"`
	Messaging    Messaging    `json:"messaging"`
	Photos       []string     `json:"photos"`
	Hours        []Hours      `json:"hours"`
	SpecialHours SpecialHours `json:"special_hours"`
}

type Region struct {
	Center Coordinates `json:"center"`
}

type Businesses struct {
	Total      int        `json:"total"`
	Businesses []Business `json:"businesses"`
	Region     Region     `json:"region"`
}

type User struct {
	Id         string `json:"id"`
	ProfileUrl string `json:"profile_url"`
	Name       string `json:"name"`
	ImageUrl   string `json:"image_url"`
}

type Review struct {
	Id          string `json:"id"`
	Text        string `json:"text"`
	Url         string `json:"url"`
	Rating      int    `json:"rating"`
	TimeCreated string `json:"time_created"`
	User        User   `json:""`
}

type Reviews struct {
	Total             int      `json:"total"`
	PossibleLanguages []string `json:"possible_languages"`
	Reviews           []Review `json:"reviews"`
}

type Term struct {
	Text string `json:"text"`
}

type BusinessId struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Autocomplete struct {
	Terms      []Term       `json:"terms"`
	Businesses []BusinessId `json:"businesses"`
	Categories []CategoryId `json:"categories"`
}
