package bridge

import "testing"

func Test_CommonSMS(t *testing.T) {
	m := NewCommonMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via SMS
}

func Test_CommonEmail(t *testing.T) {
	m := NewCommonMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via Email
}

func Test_UrgencySMS(t *testing.T) {
	m := NewUrgencyMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via SMS
}

func Test_UrgencyEmail(t *testing.T) {
	m := NewUrgencyMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via Email
}
