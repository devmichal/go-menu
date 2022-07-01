package domain

type MenuRepository interface {
	SaveMenu(*Menu)
	UpdateMenu(*Menu)
	GetMenu(idMenu string) Menu
}
