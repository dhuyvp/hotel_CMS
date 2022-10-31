package models

type ChannelListManage struct {
	ChannelListID     int    `db:"ChannelListID,omitempty" json:"channel_list_id,omitempty"`
	ChannelPackID     int    `db:"ChannelPackID,omitempty" json:"channel_pack_id,omitempty"`
	ChannelListName   string `db:"ChannelListName,omitempty" json:"channel_list_name,omitempty"`
	DetailDescription string `db:"DetailDescription,omitempty" json:"detail_description,omitempty"`
	SortOrder         int    `db:"SortOrder,omitempty" json:"sort_order,omitempty"`
	CreatedAt         string `db:"CreatedAt,omitempty" json:"created_at,omitempty"`
	UpdatedAt         string `db:"UpdatedAt,omitempty" json:"updated_at,omitempty"`
}
