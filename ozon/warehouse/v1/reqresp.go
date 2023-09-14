package v1

type ListRequest struct {
}

type ListResponseResultFirstMileType struct {
	DropoffPointID      string                                       `json:"dropoff_point_id"`
	DropoffTimeslotID   int64                                        `json:"dropoff_timeslot_id"`
	FirstMileIsChanging bool                                         `json:"first_mile_is_changing"`
	FirstMileType       ListResponseResultFirstMileTypeFirstMileType `json:"first_mile_type"`
}

type ListResponseResult struct {
	HasEntrustedAcceptance bool                            `json:"has_entrusted_acceptance"`
	IsRFBS                 bool                            `json:"is_rfbs"`
	Name                   string                          `json:"name"`
	WarehouseID            int64                           `json:"warehouse_id"`
	CanPrintActInAdvance   bool                            `json:"can_print_act_in_advance"`
	FirstMileType          ListResponseResultFirstMileType `json:"first_mile_type"`
	HasPostingsLimit       bool                            `json:"has_postings_limit"`
	IsKarantin             bool                            `json:"is_karantin"`
	IsKGT                  bool                            `json:"is_kgt"`
	IsTimetableEditable    bool                            `json:"is_timetable_editable"`
	MinPostingsLimit       int32                           `json:"min_postings_limit"`
	PostingsLimit          int32                           `json:"postings_limit"`
	MinWorkingDays         int64                           `json:"min_working_days"`
	Status                 ListResponseResultStatus        `json:"status"`
	WorkingDays            []ListResponseResultWorkingDay  `json:"working_days"`
}

type ListResponse struct {
	Result []ListResponseResult `json:"result"`
}
