package models

type ViewAllTruckRequest struct {
	MemberID  string `json:"member_id" valid:"required"`
	CountPage string `json:"count_perpage" valid:"required"`
	Page      string `json:"page" valid:"required"`
	TimeStart string `json:"time_start" valid:"required"`
	TimeEnd   string `json:"time_end" valid:"required"`
	Filter    string `json:"filter"`
}
