package routes

import (
	"fmt"
	"log/slog"
	"net/http"

	"go-example/stores"
)

func HandlePostPokemon(
	logger *slog.Logger,
	pokemonStore stores.PokemonStore,
) http.Handler {

	type request struct {
		Name string `json:"name"`
	}

	type response struct {
		Message string          `json:"message"`
		Pokemon *stores.Pokemon `json:"pokemon"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("handle post pokémon")

		decoded, _, err := decode[request](r, nil)
		if err != nil {
			m := Response{Message: err.Error()}
			encode(w, r, http.StatusBadRequest, m)
			return
		}

		name := "Squirtle"
		if decoded.Name != "" {
			name = decoded.Name
		}

		pokemon := pokemonStore.Create(&stores.Pokemon{Name: name})

		resp := response{
			Message: fmt.Sprintf("Hello, %s!", name),
			Pokemon: pokemon,
		}
		encode(w, r, http.StatusOK, resp)
	})
}

func HandleGetPokemon(
	logger *slog.Logger,
	pokemonStore stores.PokemonStore,
) http.Handler {

	type response struct {
		Pokemon []*stores.Pokemon `json:"pokemon"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("handle get pokémon")

		pokemonID := r.URL.Query().Get("id")

		resp := response{
			Pokemon: []*stores.Pokemon{},
		}

		if pokemonID == "" {
			logger.Info("get all pokémon")
			resp.Pokemon = pokemonStore.GetAll()
		} else {
			logger.Info("get one pokémon")
			pokemon := pokemonStore.Get(pokemonID)
			if pokemon != nil {
				resp.Pokemon = []*stores.Pokemon{pokemon}
			}
		}

		encode(w, r, http.StatusOK, resp)
	})
}
