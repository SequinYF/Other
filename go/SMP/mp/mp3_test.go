package mp

import "testing"

func TestMP3Player_Play(t *testing.T) {
	m := new(MP3Player)
	m0 := "haha"
	if  m == nil {
		t.Error("Mp3Player failed")
	}

	m.Play(m0)
	if m.progress == 0 {
		t.Error("Mp3Palyer.Play() Mp3 failed")
	}
}
