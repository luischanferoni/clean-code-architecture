package domain

// Pagination uses an offset-limit pattern where the offset is used
// to determine the number of items to skip unti the beginning of the page
// Consider:
//  * Some rows might be lost between page switchs if in the middle rows were added
//  * Is more performant with the first rows than the last ones
type Pagination struct {
	Skip  *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
}
