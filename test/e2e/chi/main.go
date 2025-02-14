// Copyright Sheldon Lobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func chihello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello from chi server\n")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/chi-server", chihello)
	_ = http.ListenAndServe(":8080", r)
}
