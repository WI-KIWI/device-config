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

package network

import (
	"reflect"
	"testing"
)

func Test_validateAddress(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   uint32
		wantErr bool
	}{
		{"valid-ipv4", args{"192.168.1.100/24"}, "192.168.1.100", 24, false},
		{"valid-ipv4-no-mask", args{"192.168.1.100"}, "192.168.1.100", 24, false},
		{"invalid-ipv4-mask-format", args{"192.168.1.100/255.255.255.0"}, "192.168.1.100", 24, false},
		{"invalid-ipv4-bad-mask", args{"192.168.1.100/255"}, "192.168.1.100", 24, false},
		{"invalid-ipv4-bad-addr", args{"192.168.1.999"}, "", 0, true},
		{"invalid-ipv4-empty", args{""}, "", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := validateAddress(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateAddress() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("validateAddress() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_isInList(t *testing.T) {
	tests := []struct {
		name     string
		needle   string
		haystack []string
		want     bool
	}{
		{"in-lis1", "a", []string{"a", "b", "c"}, true},
		{"in-list2", "b", []string{"a", "b", "c"}, true},
		{"not-in-list2", "z", []string{"a", "b", "c"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringInSlice(tt.needle, tt.haystack); got != tt.want {
				t.Errorf("stringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_virtualInterfaces(t *testing.T) {
	tests := []struct {
		name   string
		sysdev string
		want   []string
	}{
		{"valid", "../../config", []string{"config.go", "config_test.go"}},
		{"invalid-path", "does-not-exist", []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := virtualInterfaces(tt.sysdev); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("virtualInterfaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
