/*
 * Copyright (C) 2020 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package web

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type commonData struct {
	Username string
	Error    string
}

func (srv Web) templates(name string) (*template.Template, error) {
	// Parse the templates
	p := filepath.Join(srv.Settings.DocRoot, name)
	t, err := template.ParseFiles(p)
	if err != nil {
		log.Printf("Error loading the application template: %v\n", err)
	}
	return t, err
}

func getUsername(r *http.Request) string {
	username, err := r.Cookie("username")
	if err != nil {
		return ""
	}
	return username.String()
}
