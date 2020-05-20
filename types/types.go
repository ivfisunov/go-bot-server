package types

// MobileJSONRequest type
type MobileJSONRequest struct {
	Mobile string `json:"mobile" binding:"required"`
}

// UserLastNameJSONRequest type
type UserLastNameJSONRequest struct {
	LastName string `json:"lastName" binding:"required"`
}

// FindUserByMobileResponse response json
type FindUserByMobileResponse struct {
	DisplayName  string `json:"displayName"`
	MobileNumber string `json:"mobileNumber"`
	AccountName  string `json:"accountName"`
	PhoneNumber  string `json:"phoneNumber"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Position     string `json:"position"`
	Unit         string `json:"unit"`
	Message      string `json:"message"`
	Status       string `json:"status"`
}

// FindUserByLastNameResponse response json
type FindUserByLastNameResponse struct {
	Employees []User `json:"employees"`
	Message   string `json:"message"`
	Status    string `json:"status"`
}

// User all user's fields
type User struct {
	PersonalID     string `json:"personalId"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	MiddleName     string `json:"middleName"`
	DisplayName    string `json:"displayName"`
	Birthday       string `json:"birthday"`
	Position       string `json:"position"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	AccountName    string `json:"accountName"`
	EmploymentType string `json:"employmentType"`
	Hired          string `json:"hired"`
	Fired          string `json:"fired"`
	State          string `json:"state"`
	ExpDate        string `json:"expDate"`
	Unit           string `json:"unit"`
	RoomID         string `json:"roomId"`
	WorkplaceID    string `json:"workplaceId"`
	MobileNumber   string `json:"mobileNumber"`
	HideBirthday   bool   `json:"hideBirthday"`
	Telegram       string `json:"telegram"`
	Expertise      string `json:"expertise"`
	TeachSkills    string `json:"techSkills"`
	About          string `json:"about"`
	PlaceInfo      struct {
		WorkplaceID string `json:"workplaceId"`
		Place       string `json:"place"`
		Description string `json:"description"`
	} `json:"placeInfo"`
}
