package alerter

import (
	"context"

	"github.com/rs/zerolog/log"
)

type defaultAlerter struct {
	logAlerts bool
}

// DefaultAlerter useful when an external source is disabled, but you
// want to offer the interface as a noop (logAlerts=false) or log to
// the logger (logAlerts=true)
func NewDefaultAlerter(logAlerts bool) Alerter {
	return &defaultAlerter{logAlerts: logAlerts}
}

func (a *defaultAlerter) Alert(ctx context.Context, format string, v ...interface{}) {
	if a.logAlerts {
		log.Error().Str("alert", "alert").Msgf(format, v...)
	}
}
