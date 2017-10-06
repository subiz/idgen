package ID_test

import (
	"testing"
	. "bitbucket.org/subiz/id"
	"bitbucket.org/subiz/gocommon"
)

func TestId(t *testing.T) {
	ids := []struct{
		name, id string
	}{
		{ "tag", NewTagID() },
		{ "account", NewAccountID() },
		{ "agent", NewAgentID() },
		{ "event", NewEventID() },
		{ "user", NewUserID() },
	}

	for i := range ids {
		common.Logf("%8s: %s", ids[i].name, ids[i].id)
	}
}
