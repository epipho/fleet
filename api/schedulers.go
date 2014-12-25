/*
   Copyright 2014 CoreOS, Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package api

import (
	"errors"
	"log"
	"net/http"
	"path"

	"github.com/coreos/fleet/client"
)

func wireUpSchedulersResource(mux *http.ServeMux, prefix string, cAPI client.API) {
	base := path.Join(prefix, "schedulers")
	sr := schedulersResource{cAPI, base}
	mux.Handle(base, &sr)
	mux.Handle(base+"/", &sr)

	log.Printf("Registering Stuff")
}

type schedulersResource struct {
	cAPI     client.API
	basePath string
}

func (sr *schedulersResource) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if isCollectionPath(sr.basePath, req.URL.Path) {
		switch req.Method {
		case "GET":
		default:
			sendError(rw, http.StatusMethodNotAllowed, errors.New("only GET supported against this resource"))
		}
	} else if _, ok := isItemPath(sr.basePath, req.URL.Path); ok {
		switch req.Method {
		case "GET":
		case "PUT":
		case "DELETE":
		default:
			sendError(rw, http.StatusMethodNotAllowed, errors.New("only GET, PUT, and DELETE supported against this resource"))
		}
	} else {
		sendError(rw, http.StatusForbidden, nil)
	}
}
