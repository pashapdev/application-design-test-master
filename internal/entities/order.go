package entities

import "time"

type Order struct {
	Room      RoomType
	UserEmail Email
	From      time.Time
	To        time.Time
}
