package stores

import (
	"maps"
	"slices"

	"github.com/google/uuid"
)

type Pokemon struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PokemonStore interface {
	Create(pokemon *Pokemon) *Pokemon
	Get(userID string) *Pokemon
	GetAll() []*Pokemon
}

// constructor function
func NewInMemoryPokemonStore() *InMemoryPokemonStore {
	return &InMemoryPokemonStore{
		pokemon: map[string]*Pokemon{},
	}
}

type InMemoryPokemonStore struct {
	pokemon map[string]*Pokemon
}

func (s *InMemoryPokemonStore) Create(pokemon *Pokemon) *Pokemon {
	id := uuid.New().String()

	copy := &Pokemon{
		Id:   id,
		Name: pokemon.Name,
	}

	s.pokemon[id] = copy

	return copy
}

func (s *InMemoryPokemonStore) Get(pokemonID string) *Pokemon {
	val, ok := s.pokemon[pokemonID]

	if ok {
		return val
	}

	return nil
}

func (s *InMemoryPokemonStore) GetAll() []*Pokemon {
	vals := slices.Collect(maps.Values(s.pokemon))
	if len(vals) == 0 {
		return []*Pokemon{}
	}
	return vals
}
