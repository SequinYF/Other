package library

import (
	"errors"
	"fmt"
	"strconv"
)
type MusicEntry struct {
	Id string
	Name string
	Artist string
	Source string
	Type string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager)Len() int  {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index > len(m.musics) {
		return nil, errors.New("Index out of range")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, mm := range m.musics {
		if name == mm.Name {
			return &mm
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}


func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]

	if index < len(m.musics) - 1 {
		m.musics = append(m.musics[:len(m.musics) - 1], m.musics[index + 1:]...)
	} else if index == 0 {
		m.musics = make([]MusicEntry, 0)
	} else {
		m.musics = m.musics[:index - 1]
	}

	return removedMusic
}

func (m *MusicManager)RemoveByName(name string) *MusicEntry{

	music := m.Find(name)
	if music == nil {
		fmt.Println("no such item to remove", name)
	} else {
		index, _ := strconv.Atoi(music.Id)
		ret := m.Remove(index)
		if ret == nil {
			fmt.Println("cant remove item", name)
		} else {
			fmt.Println("remove sucessed!")
			return ret
		}
	}
	return nil
}
