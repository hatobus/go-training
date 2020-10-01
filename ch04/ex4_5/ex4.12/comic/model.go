package comic

const URL = "https://xkcd.com"

type Comic struct {
	Num        int    `json:"num"`
	Title      string `json:"title"`
	SafeTitle  string `json:"safe_title"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	News       string `json:"news"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Link       string `json:"link"`
	Img        string `json:"img"`
}
