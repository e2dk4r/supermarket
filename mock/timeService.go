package mock

import "time"

type TimeService struct {
	NowFn      func() time.Time
	nowInvoked bool
}

func (ts *TimeService) Now() time.Time {
	ts.nowInvoked = true
	return ts.NowFn()
}

func (ts *TimeService) NowInvoked() bool {
	return ts.nowInvoked
}
