package menu

type MenuResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
	Available   bool   `json:"available"`
	ImageURL    string `json:"image_url"`
}

func ToResponse(menu *Menu) MenuResponse {
	return MenuResponse{
		ID:          menu.ID.String(),
		Name:        menu.Name,
		Description: menu.Description,
		Price:       menu.Price,
		Stock:       menu.Stock,
		Available:   menu.Available,
		ImageURL:    menu.ImageURL,
	}
}

func ToResponses(menus []Menu) []MenuResponse {

	responses := make([]MenuResponse, 0, len(menus))

	for _, menu := range menus {

		responses = append(responses, MenuResponse{
			ID:          menu.ID.String(),
			Name:        menu.Name,
			Description: menu.Description,
			Price:       menu.Price,
			Stock:       menu.Stock,
			Available:   menu.Available,
			ImageURL:    menu.ImageURL,
		})

	}

	return responses
}