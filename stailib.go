package stailib

import (
	"github.com/bigobject-inc/stailib/cbe"
	"github.com/bigobject-inc/stailib/cs"
	"github.com/bigobject-inc/stailib/tt"
)

// STAILibrary STAI library
type STAILibrary interface {
	GetServTrackTrace() tt.TrackTrace
	GetServCBE() cbe.CBEServ
	GetServCS() cs.CSServ
	GetVersion() string
}
