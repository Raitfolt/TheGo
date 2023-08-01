package main

import "testing"

func Test_hasHttpPrefix(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "google.com",
			s:    "google.com",
			want: false,
		},
		{
			name: "http://google.com",
			s:    "http://google.com",
			want: true,
		},
		{
			name: "https://google.com",
			s:    "https://google.com",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasHttpPrefix(tt.s); got != tt.want {
				t.Errorf("hasHttpPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasHttpsPrefix(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "google.com",
			s:    "google.com",
			want: false,
		},
		{
			name: "http://google.com",
			s:    "http://google.com",
			want: false,
		},
		{
			name: "https://google.com",
			s:    "https://google.com",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasHttpsPrefix(tt.s); got != tt.want {
				t.Errorf("hasHttpsPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
