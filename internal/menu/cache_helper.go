package menu

import "encoding/json"

func encodeMenus(menus []Menu) ([]byte, error) {
	return json.Marshal(menus)
}

func decodeMenus(data string) ([]Menu, error) {

	var menus []Menu

	err := json.Unmarshal([]byte(data), &menus)

	return menus, err
}