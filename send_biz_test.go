package whatsmeow

import (
	"testing"

	"go.mau.fi/whatsmeow/proto/waE2E"
)

// Real clients attach no <biz> node to a button or list RESPONSE; the server
// rejects one with 479 smax-invalid (tulir/whatsmeow#1221).
func TestButtonTypeOmittedForResponses(t *testing.T) {
	cases := map[string]struct {
		msg  *waE2E.Message
		want string
	}{
		"list_response":      {&waE2E.Message{ListResponseMessage: &waE2E.ListResponseMessage{}}, ""},
		"buttons_response":   {&waE2E.Message{ButtonsResponseMessage: &waE2E.ButtonsResponseMessage{}}, ""},
		"list":               {&waE2E.Message{ListMessage: &waE2E.ListMessage{}}, "list"},
		"buttons":            {&waE2E.Message{ButtonsMessage: &waE2E.ButtonsMessage{}}, "buttons"},
		"ephemeral_list_rsp": {&waE2E.Message{EphemeralMessage: &waE2E.FutureProofMessage{Message: &waE2E.Message{ListResponseMessage: &waE2E.ListResponseMessage{}}}}, ""},
	}
	for name, c := range cases {
		if got := getButtonTypeFromMessage(c.msg); got != c.want {
			t.Errorf("%s: getButtonTypeFromMessage = %q, want %q", name, got, c.want)
		}
	}
}

// The mediatype attribute stays: the PR found it correct as-is.
func TestMediaTypeKeptForListResponse(t *testing.T) {
	if got := getMediaTypeFromMessage(&waE2E.Message{ListResponseMessage: &waE2E.ListResponseMessage{}}); got != "list_response" {
		t.Errorf("mediatype = %q, want list_response", got)
	}
}
