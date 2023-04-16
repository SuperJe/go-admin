package dto

type AllDashboardReq struct {
}

type Progression struct {
	Done       uint32 `json:"done"`
	Unfinished uint32 `json:"unfinished"`
	Total      uint32 `json:"total"`
}

type CampaignProgression struct {
	Dungeon  *Progression `json:"dungeon"`
	Forrest  *Progression `json:"forrest"`
	Desert   *Progression `json:"desert"`
	Mountain *Progression `json:"mountain"`
	Glacier  *Progression `json:"glacier"`
}

type AllDashboardRsp struct {
	CampProgressions *CampaignProgression `json:"camp_progressions"`
}
