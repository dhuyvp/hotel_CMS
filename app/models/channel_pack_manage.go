package models

type ChannelPackManage struct {
	ChannelPackID     int    `db:"ChannelPackID,omitempty" json:"channel_pack_id,omitempty"`
	HotelID           int    `db:"HotelID,omitempty" json:"hotel_id,omitempty"`
	Logo              string `db:"Logo,omitempty" json:"logo,omitempty"`
	ChannelPackName   string `db:"ChannelPackName,omitempty" json:"channel_pack_name,omitempty"`
	DetailDescription string `db:"DetailDescription,omitempty" json:"detail_description,omitempty"`
	Note              string `db:"Note,omitempty" json:"note,omitempty"`
	IsActive          bool   `db:"IsActive,omitempty" json:"is_active,omitempty"`
}
