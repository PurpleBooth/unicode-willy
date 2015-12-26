// It's a web server that prints an ASCII penis.
// Copyright (C) 2015  Billie Thompson <billie@purplebooth.co.uk>
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along with
// this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"log"
	"net/http"
	"html"
)

const port = 8080

const htmlTemplate = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">

  <title>Penis</title>
</head>

<body>
  <pre>%s</pre>
</body>
</html>`

const penis = `                     _,-%/%|
                _,-'    \//%\
            _,-'        \%/|%
          / / )    __,--  /%\
          \__/_,-'%(%  ;  %)%
                  %\%,   %\
                    '--%'      `

const httpContentTypeValue = "text/html; charset=utf-8"
const httpContentTypeHeader = "Content-Type"

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("8=====>")

	w.Header().Set(httpContentTypeHeader, httpContentTypeValue)
	fmt.Fprintf(w, htmlTemplate, html.EscapeString(penis))
}

func main() {
	listeningAddress := fmt.Sprintf(":%d", port)
	log.Println(fmt.Sprintf("Listening on %s", listeningAddress))

	http.HandleFunc("/", handler)
	http.ListenAndServe(listeningAddress, nil)
}