package entities

type RoomType string

const (
	RoomTypeEconom   = "econom"
	RoomTypeStandart = "standart"
	RoomTypeLux      = "lux"
)

var ValidRoomTypes = map[string]struct{}{
	"econom":   {},
	"standart": {},
	"lux":      {},
}
