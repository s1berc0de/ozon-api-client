package v1

type GEORestrictionsCatalogByFilterRequestFilter struct {
	Names       []string `json:"names,omitempty"`
	OnlyVisible bool     `json:"only_visible"`
}

type GEORestrictionsCatalogByFilterRequest struct {
	Filter          GEORestrictionsCatalogByFilterRequestFilter `json:"filter"`
	LastOrderNumber int64                                       `json:"last_order_number"`
	Limit           int64                                       `json:"limit"`
}

type GEORestrictionsCatalogByFilterResponseRestriction struct {
	ID          string `json:"id"`
	IsVisible   bool   `json:"is_visible"`
	Name        string `json:"name"`
	OrderNumber int64  `json:"order_number"`
}

type GEORestrictionsCatalogByFilterResponse struct {
	Restrictions []GEORestrictionsCatalogByFilterResponseRestriction
}
