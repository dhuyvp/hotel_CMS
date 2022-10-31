package models

type DeviceManage struct {
	DeviceID          int    `db:"DeviceID,omitempty" json:"device_id,omitempty"`
	HotelID           int    `db:"HotelID,omitempty" json:"hotel_id,omitempty"`
	MacWired          string `db:"MacWired,omitempty" json:"mac_wired,omitempty"`
	MacWireless       string `db:"MacWireless,omitempty" json:"mac_wireless,omitempty"`
	DeviceName        string `db:"DeviceName,omitempty" json:"device_name,omitempty"`
	DetailDescription string `db:"DetailDescription,omitempty" json:"detail_description,omitempty"`
	IsActive          bool   `db:"IsActive,omitempty" json:"is_active,omitempty"`
	CreatedAt         string `db:"CreatedAt,omitempty" json:"created_at,omitempty"`
	UpdatedAt         string `db:"UpdatedAt,omitempty" json:"updated_at,omitempty"`
}
