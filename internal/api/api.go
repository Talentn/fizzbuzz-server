package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/talentn/fizzbuzz-server/internal/fizzbuzz"
	"github.com/talentn/fizzbuzz-server/internal/stats"
)

type Server struct {
	stats *stats.Store
}

func NewHandler(store *stats.Store) http.Handler {
	s := &Server{stats: store}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /fizzbuzz", s.handleFizzBuzz)
	mux.HandleFunc("GET /stats", s.handleStats)

	return mux
}

func (s *Server) handleFizzBuzz(w http.ResponseWriter, r *http.Request) {
	params, err := parseParams(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := params.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	s.stats.Record(params)
	writeJSON(w, http.StatusOK, map[string]any{"result": fizzbuzz.Generate(params)})
}

func (s *Server) handleStats(w http.ResponseWriter, r *http.Request) {
	params, hits, ok := s.stats.Top()
	if !ok {
		writeJSON(w, http.StatusOK, map[string]any{"message": "no fizzbuzz request recorded yet"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"most_frequent_request": params,
		"hits":                  hits,
	})
}

func parseParams(r *http.Request) (fizzbuzz.Params, error) {
	q := r.URL.Query()
	var problems []string

	parseInt := func(name string) int {
		raw := q.Get(name)
		if raw == "" {
			problems = append(problems, fmt.Sprintf("missing required parameter %q", name))
			return 0
		}
		v, err := strconv.Atoi(raw)
		if err != nil {
			problems = append(problems, fmt.Sprintf("parameter %q must be an integer, got %q", name, raw))
			return 0
		}
		return v
	}

	params := fizzbuzz.Params{
		Int1:  parseInt("int1"),
		Int2:  parseInt("int2"),
		Limit: parseInt("limit"),
		Str1:  q.Get("str1"),
		Str2:  q.Get("str2"),
	}

	if len(problems) > 0 {
		return fizzbuzz.Params{}, fmt.Errorf("%s", strings.Join(problems, "\n"))
	}
	return params, nil
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]any{"errors": strings.Split(message, "\n")})
}