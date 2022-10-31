package models

type HotelManage struct {
	HotelID           int    `db:"HotelID,omitempty" json:"hotel_id,omitempty"`
	Logo              string `db:"Logo,omitempty" json:"logo,omitempty"`
	HotelName         string `db:"HotelName,omitempty" json:"hotel_name,omitempty"`
	DevicesLimit      int    `db:"DevicesLimit,omitempty" json:"devices_limit,omitempty"`
	DevicesNumber     int    `db:"DevicesNumber,omitempty" json:"devices_number,omitempty"`
	TotalRoom         int    `db:"TotalRoom,omitempty" json:"total_room,omitempty"`
	Address           string `db:"Address,omitempty" json:"address,omitempty"`
	DetailDescription string `db:"DetailDescription,omitempty" json:"detail_description,omitempty"`
	IsActive          bool   `db:"IsActive,omitempty" json:"is_active,omitempty"`
}
