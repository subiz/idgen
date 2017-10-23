package ID_test

import (
	"testing"
  _ "fmt"
	. "bitbucket.org/subiz/id"
	"bitbucket.org/subiz/gocommon"
)

func TestId(t *testing.T) {
	ids := []struct{
		name, id string
	}{
		{ "tag", NewTagID() },
		{ "account", NewAccountID() },
		{ "group", NewAgentGroupID() },
		{ "agent", NewAgentID() },
		{ "event", NewEventID() },
		{ "user", NewUserID() },
		{ "file", NewFileID() },
		{ "button", NewButtonID() },
		{ "conversation", NewConversationID() },
		{ "invitation", NewInvitationID() },
		{ "scheduler item", NewScheduleItemID() },
		{ "web send", NewWebsendID() },
		{ "auth token", NewAuthToken() },
		{ "refresh token", NewRefreshToken() },
		{ "rule", NewRuleID() },
		{ "client", NewClientID() },
		{ "error", NewErrorID() },
		{ "canned response", NewCannedResponseID() },
		{ "ws conn", NewWsConnID(10) },
	}

	for i := range ids {
		common.Logf("%15s: %s", ids[i].name, ids[i].id)
	}
}
