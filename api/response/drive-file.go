package response

type DriveFileProperties struct {
	Width       *int    `json:"width,omitempty"`
	Height      *int    `json:"height,omitempty"`
	Orientation *int    `json:"orientation,omitempty"`
	AvgColor    *string `json:"avgColor,omitempty"`
}
