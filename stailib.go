package stailib

import (
	"github.com/bigobject-inc/stailib/cbe"
	"github.com/bigobject-inc/stailib/tt"
	"github.com/bigobject-inc/stailib/cs"
)

// STAILibrary STAI library
type STAILibrary interface {
	GetServTrackTrace() tt.TrackTrace
	GetServCBE() cbe.CBEServ
	GetServCS() cs.CSServ
	GetVersion() string
}
