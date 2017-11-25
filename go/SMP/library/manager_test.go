package library

import (
	"testing"
	"fmt"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()

	if mm == nil {
		t.Error("NerMusicManager failed.")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager faild, not empty.")
	}

	m0 := &MusicEntry {"0", "My Heart whill go on", "Celion Dion","http:/...", "MP3"}

	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed.")
	}

	m := mm.Find(m0.Name)

	if m == nil {
		t.Error("MusicManager.Find() failed")
	}

	if m.Id != m0.Id || m.Artist != m0.Artist || m.Name != m0.Name || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager.Find() failed. Found item mismath")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager Get() failed,", err)
	} else {
		fmt.Println(m.Id, m.Name, m.Type, m.Source, m.Artist)
	}

	ret := mm.RemoveByName("My Heart whill go on")
	if ret == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() faild.", err)
	}
}

