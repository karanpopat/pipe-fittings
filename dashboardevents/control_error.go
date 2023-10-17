package dashboardevents

import (
	"github.com/turbot/pipe-fittings/controlstatus"
	"time"
)

type ControlError struct {
	Control     controlstatus.ControlRunStatusProvider
	Progress    *controlstatus.ControlProgress
	Name        string
	Session     string
	ExecutionId string
	Timestamp   time.Time
}

// IsDashboardEvent implements DashboardEvent interface
func (*ControlError) IsDashboardEvent() {}
