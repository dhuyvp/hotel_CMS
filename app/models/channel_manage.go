package models

type ChannelManage struct {
	ChannelID         int    `db:"ChannelID,omitempty" json:"channel_id,omitempty"`
	ChannelListID     int    `db:"ChannelListID,omitempty" json:"channel_list_id,omitempty"`
	Logo              string `db:"Logo,omitempty" json:"logo,omitempty"`
	ChannelName       string `db:"ChannelName,omitempty" json:"channel_name,omitempty"`
	Link              string `db:"Link,omitempty" json:"link,omitempty"`
	DetailDescription string `db:"DetailDescription,omitempty" json:"detail_description,omitempty"`
	IsActive          bool   `db:"IsActive,omitempty" json:"is_active,omitempty"`
	CreatedAt         string `db:"CreatedAt,omitempty" json:"created_at,omitempty"`
}
