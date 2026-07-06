package menu

func ToResponse(menu Menu) MenuResponse {

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

	result := make([]MenuResponse, 0, len(menus))

	for _, menu := range menus {

		result = append(result, ToResponse(menu))

	}

	return result
}