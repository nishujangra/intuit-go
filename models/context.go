package models

type Context struct {
	Mobile      string `json:"mobile,omitempty" type:"boolean"`
	IsEcommerce string `json:"isEcommerce,omitempty" type:"boolean"`

	Tax        string      `json:"tax,omitempty"`
	DeviceInfo *DeviceInfo `json:"deviceInfo,omitempty"`
	Recurring  bool        `json:"recurring,omitempty" type:"boolean"`
	Restaurent *Restaurent `json:"restaurent,omitempty"`
	Loding     *Loding     `json:"lodgin,omitempty"`
}

type DeviceInfo struct {
	DeviceID    string  `json:"deviceId,omitempty"`
	MACAddress  string  `json:"macAddress,omitempty"`
	Encrypted   bool    `json:"encrypted,omitempty"`
	IPAddress   string  `json:"ipAddress,omitempty"`
	Long        float64 `json:"longitude,omitempty"`
	PhoneNumber string  `json:"phoneNumber,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Type        string  `json:"type,omitempty"`
}

type Restaurent struct {
	BeverageAmount string `json:"beverageAmount,omitempty"`
	ServerID       string `json:"serverId,omitempty"`
	TaxAmount      string `json:"taxAmount,omitempty"`
	FoodAmount     string `json:"foodAmount,omitempty"`
	TipAmount      string `json:"tipAmount,omitempty"`
}

type Loding struct {
	LengthOfStay   string `json:"lengthOfStay,omitempty"`
	CheckInDate    string `json:"checkInDate,omitempty"` // CCCC-MM-DD
	RoomRate       string `json:"roomRate,omitempty"`
	SpecialProgram string `json:"specialProgram,omitempty" enum:"AdvanceDeposit, AssuredReservation, DelayedCharge, ExpressService, NormalCharge, NoShowCharge"`
}
