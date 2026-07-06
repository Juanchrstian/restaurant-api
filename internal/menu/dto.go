package menu

type MenuResponse struct {

	ID          string  `json:"id"`

	Name        string  `json:"name"`
	
	Description string  `json:"description"`
	
	Price       int64 `json:"price"`
	
	Stock       int     `json:"stock"`
	
	Available   bool    `json:"available"`
	
	ImageURL    string  `json:"image_url"`

}
