package cbe

// CBEObjStruct : CBEObjStruct
type CBEObjStruct interface {
	// get info from frame
	GetFrameID() string
	GetSensorID() string
	GetFloorPlanID() string
	GetAreaID() string
	GetFieldID() string

	GetTimestamp() int

	// get object info
	GetGlobalID() string
	GetBackLink() string
	GetSequenceID() string
	GetTrackingID() int
	GetName() string
	GetProbability() float64
	GetConfidence() float64
	GetBox() BoundingBox
	GetLocation() Location
	GetTag(label, name string) (v interface{})
	GetFrameInfo() CBE

	// set object info
	SetGlobalID(globalID string)
	SetBackLink(backLink string)
	SetTrackingID(trackingID int)
	SetName(name string)
	SetProbability(probability float64)
	SetConfidence(confidence float64)
	SetBox(box BoundingBox)
	SetLocation(geo Location)
	SetTag(label, name string, value interface{})

	SetFrameInfo(CBE)
	IsName(name string) bool
}
