package externalapistorage

import (
	"errors"

	externalapiproto "github.com/YarikRevich/HideSeek-Server/internal/api/external-api/v1/proto"
)

// var ErrExternalStorageValueDoesNotExist = errors.New("error in external storage does not exist")

type PCs []*externalapiproto.PC
type Weapons []*externalapiproto.Weapon
type Elements []*externalapiproto.Element
type Ammos []*externalapiproto.Ammo

func (p *PCs) Add(pc *externalapiproto.PC) {
	*p = append(*p, pc)
}

func (p *PCs) Remove(pc *externalapiproto.PC) {
	for i, v := range *p {
		if v.Base.Id == pc.Base.Id {
			*p = append((*p)[:i], (*p)[i+1:]...)
			break
		}
	}
}

// func (p PCs) LookFor(pc *externalapiproto.PC)

func (w *Weapons) Add(we *externalapiproto.Weapon) {
	*w = append(*w, we)
}

func (w *Weapons) Remove(pc *externalapiproto.Weapon) {
	for i, v := range *w {
		if v.Base.Id == pc.Base.Id {
			*w = append((*w)[:i], (*w)[i+1:]...)
			break
		}
	}
}

func (e *Elements) Add(el *externalapiproto.Element) {
	*e = append(*e, el)
}

func (e *Elements) Remove(pc *externalapiproto.Element) {
	for i, v := range *e {
		if v.Base.Id == pc.Base.Id {
			*e = append((*e)[:i], (*e)[i+1:]...)
			break
		}
	}
}

func (a *Ammos) Add(am *externalapiproto.Ammo) {
	*a = append(*a, am)
}

func (a *Ammos) Remove(am *externalapiproto.Ammo) {
	for i, v := range *a {
		if v.Base.Id == am.Base.Id {
			*a = append((*a)[:i], (*a)[i+1:]...)
			break
		}
	}
}

type WorldStorage map[string]*externalapiproto.World
type MapStorage map[string]*externalapiproto.Map
type PCStorage map[string]*PCs
type ElementStorage map[string]*Elements
type WeaponStorage map[string]*Weapons
type AmmoStorage map[string]*Ammos

func (ws *WorldStorage) Add(key string, value *externalapiproto.World) {
	(*ws)[key] = value
}

func (ws *WorldStorage) Remove(key string) {
	delete((*ws), key)
}

func (ws *WorldStorage) Get(key string) *externalapiproto.World {
	v, ok : = (*ws)[key]
	if !ok {
		return &externalapiproto.World{}
	}
	return v
}

func (ms *MapStorage) Add(key string, value *externalapiproto.Map) {
	(*ms)[key] = value
}

func (ms *MapStorage) Remove(key string) {
	delete((*ms), key)
}

func (ms MapStorage) Get(key string) *externalapiproto.Map {
	return ms[key]
}

func (ps *PCStorage) Add(key string, value *externalapiproto.PC) {
	if _, ok := (*ps)[key]; !ok {
		(*ps)[key] = new(PCs)
	}
	(*ps)[key].Add(value)
}

func (ps *PCStorage) Remove(key string) {
	delete((*ps), key)
}

func (ps *PCStorage) Get(key string) PCs {
	v, ok := (*ps)[key]
	if !ok {
		return PCs{}
	}
	return *v
}

func (ps *PCStorage) Length() int {
	return len(*ps)
}

func (es *ElementStorage) Add(key string, value *externalapiproto.Element) {
	if _, ok := (*es)[key]; !ok {
		(*es)[key] = new(Elements)
	}
	(*es)[key].Add(value)
}

func (es *ElementStorage) Remove(key string) {
	delete((*es), key)
}

func (es *ElementStorage) Get(key string) Elements {
	v, ok := (*es)[key]
	if !ok {
		return Elements{}
	}
	return *v
}

func (ws *WeaponStorage) Add(key string, value *externalapiproto.Weapon) {
	if _, ok := (*ws)[key]; !ok {
		(*ws)[key] = new(Weapons)
	}
	(*ws)[key].Add(value)
}

func (ws *WeaponStorage) Remove(key string) {
	delete((*ws), key)
}

func (ws *WeaponStorage) Get(key string) Weapons {

	v, ok := (*ws)[key]
	if !ok {
		return Weapons{}
	}
	return *v
}

func (as *AmmoStorage) Add(key string, value *externalapiproto.Ammo) {
	if _, ok := (*as)[key]; !ok {
		(*as)[key] = new(Ammos)
	}
	(*as)[key].Add(value)
}

func (as *AmmoStorage) Remove(key string) {
	delete((*as), key)
}

func (as *AmmoStorage) Get(key string) Ammos {
	v, ok := (*as)[key]
	if !ok {
		return Ammos{}
	}
	return *v
}

type ExternalApiStorage struct {
	WorldStorage
	MapStorage
	PCStorage
	ElementStorage
	WeaponStorage
	AmmoStorage
}

func NewExternalApiStorage() *ExternalApiStorage {
	return &ExternalApiStorage{
		make(WorldStorage),
		make(MapStorage),
		make(PCStorage),
		make(ElementStorage),
		make(WeaponStorage),
		make(AmmoStorage),
	}
}
