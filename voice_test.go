package disgord

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/andersfylling/disgord/endpoint"
	"github.com/andersfylling/disgord/httd"
	"github.com/andersfylling/disgord/ratelimit"
)

func TestStateMarshalling(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/voice/state1.json")
	check(err, t)

	state := VoiceState{}
	err = httd.Unmarshal(data, &state)
	check(err, t)
}

func TestVoice_InterfaceImplementations(t *testing.T) {
	t.Run("VoiceState", func(t *testing.T) {
		var u interface{} = &VoiceState{}
		t.Run("DeepCopier", func(t *testing.T) {
			if _, ok := u.(DeepCopier); !ok {
				t.Error("does not implement DeepCopier")
			}
		})

		t.Run("Copier", func(t *testing.T) {
			if _, ok := u.(Copier); !ok {
				t.Error("does not implement Copier")
			}
		})
	})
	t.Run("VoiceRegion", func(t *testing.T) {
		var u interface{} = &VoiceRegion{}
		t.Run("DeepCopier", func(t *testing.T) {
			if _, ok := u.(DeepCopier); !ok {
				t.Error("does not implement DeepCopier")
			}
		})

		t.Run("Copier", func(t *testing.T) {
			if _, ok := u.(Copier); !ok {
				t.Error("does not implement Copier")
			}
		})
	})
}

func TestListVoiceRegions(t *testing.T) {
	client, _, err := createTestRequester()
	if err != nil {
		t.Skip()
		return
	}

	builder := &listVoiceRegionsBuilder{}
	builder.IgnoreCache().setup(nil, client, &httd.Request{
		Method:      http.MethodGet,
		Ratelimiter: ratelimit.VoiceRegions(),
		Endpoint:    endpoint.VoiceRegions(),
	}, nil)

	list, err := builder.Execute()
	if err != nil {
		t.Error(err)
	}

	if len(list) == 0 {
		t.Error("expected at least one voice region")
	}
}

func TestVoiceState_CopyOverTo(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/voice/state1.json")
	check(err, t)

	tei := VoiceState{}
	err = httd.Unmarshal(data, &tei)
	check(err, t)

	t.Run("CopyOver no error", func(t *testing.T) {
		actual := VoiceState{}

		err = tei.CopyOverTo(&actual)
		check(err, t)

		if actual.ChannelID != tei.ChannelID {
			t.Errorf("Expected ChannelId: %v, got: %v", tei.ChannelID, actual.ChannelID)
		}

		if actual.GuildID != tei.GuildID {
			t.Errorf("Expected GuildId: %v, got: %v", tei.GuildID, actual.GuildID)
		}

		if actual.UserID != tei.UserID {
			t.Errorf("Expected UserID: %v, got: %v", tei.UserID, actual.UserID)
		}

		if actual.SessionID != tei.SessionID {
			t.Errorf("Expected SessionID: %v, got: %v", tei.ChannelID, actual.ChannelID)
		}
	})

	t.Run("Should handle wrong interface input", func(t *testing.T) {
		twi := VoiceRegion{}

		err = tei.CopyOverTo(twi)

		if err == nil {
			t.Error("Expected error, got non")
		}

		if _, ok := err.(*ErrorUnsupportedType); !ok {
			t.Error("Wrong error type returned")
		}
	})
}

func TestVoiceState_DeepCopy(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/voice/state1.json")
	check(err, t)

	tei := VoiceState{}
	err = httd.Unmarshal(data, &tei)
	check(err, t)

	actual := tei.DeepCopy()

	if _, ok := actual.(*VoiceState); !ok {
		t.Error("Wrong type of interface returned")
	}

	if actual.(*VoiceState).ChannelID != tei.ChannelID {
		t.Errorf("Expected ChannelID: %v, got: %v", tei.ChannelID, actual.(*VoiceState).ChannelID)
	}

	if actual.(*VoiceState).GuildID != tei.GuildID {
		t.Errorf("Expected GuildID: %v, got: %v", tei.GuildID, actual.(*VoiceState).GuildID)
	}

	if actual.(*VoiceState).UserID != tei.UserID {
		t.Errorf("Expected UserID: %v, got: %v", tei.UserID, actual.(*VoiceState).UserID)
	}

	if actual.(*VoiceState).SessionID != tei.SessionID {
		t.Errorf("Expected SessionID: %v, got: %v", tei.ChannelID, actual.(*VoiceState).ChannelID)
	}
}

func TestVoiceRegion_CopyOverTo(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/voice/region1.json")
	check(err, t)

	tei := VoiceRegion{}
	err = httd.Unmarshal(data, &tei)
	check(err, t)

	t.Run("CopyOver no error", func(t *testing.T) {
		actual := VoiceRegion{}

		err = tei.CopyOverTo(&actual)
		check(err, t)

		if actual.ID != tei.ID {
			t.Errorf("Expected ID: %v, got: %v", tei.ID, actual.ID)
		}

		if actual.Name != tei.Name {
			t.Errorf("Expected Name: %v, got: %v", tei.SampleHostname, actual.SampleHostname)
		}

		if actual.SamplePort != tei.SamplePort {
			t.Errorf("Expected SamplePort: %v, got: %v", tei.SamplePort, actual.SamplePort)
		}
	})

	t.Run("Should handle wrong interface input", func(t *testing.T) {
		twi := VoiceState{}

		err = tei.CopyOverTo(twi)

		if err == nil {
			t.Error("Expected error, got non")
		}

		if _, ok := err.(*ErrorUnsupportedType); !ok {
			t.Error("Wrong error type returned")
		}
	})
}

func TestVoiceRegion_DeepCopy(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/voice/region1.json")
	check(err, t)

	tei := VoiceRegion{}
	err = httd.Unmarshal(data, &tei)
	check(err, t)

	actual := tei.DeepCopy()

	if _, ok := actual.(*VoiceRegion); !ok {
		t.Error("Wrong type of interface returned")
	}

	if actual.(*VoiceRegion).ID != tei.ID {
		t.Errorf("Expected ID: %v, got: %v", tei.ID, actual.(*VoiceRegion).ID)
	}

	if actual.(*VoiceRegion).Name != tei.Name {
		t.Errorf("Expected Name: %v, got: %v", tei.Name, actual.(*VoiceRegion).Name)
	}

	if actual.(*VoiceRegion).SampleHostname != tei.SampleHostname {
		t.Errorf("Expected SampleHostName: %v, got: %v", tei.SampleHostname, actual.(*VoiceRegion).SampleHostname)
	}
}
