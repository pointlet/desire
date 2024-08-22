package models

type Path struct {
	ID             uint
	Title          string
	Description    string
	ImageLookupId  uint
	SignedImageUrl string
	Coordinate     Coordinate
}
