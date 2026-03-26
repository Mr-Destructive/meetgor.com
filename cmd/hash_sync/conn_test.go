package main

import "testing"

func TestBuildLibsqlConnStr(t *testing.T) {
	tests := []struct {
		name   string
		dbURL  string
		token  string
		expect string
	}{
		{
			name:   "host with token",
			dbURL:  "notes-mr-destructive.aws-ap-south-1.turso.io",
			token:  "tok",
			expect: "libsql://notes-mr-destructive.aws-ap-south-1.turso.io?authToken=tok",
		},
		{
			name:   "host without token",
			dbURL:  "notes-mr-destructive.aws-ap-south-1.turso.io",
			token:  "",
			expect: "libsql://notes-mr-destructive.aws-ap-south-1.turso.io",
		},
		{
			name:   "libsql url with token",
			dbURL:  "libsql://notes-mr-destructive.aws-ap-south-1.turso.io",
			token:  "tok",
			expect: "libsql://notes-mr-destructive.aws-ap-south-1.turso.io?authToken=tok",
		},
		{
			name:   "libsql url with existing query",
			dbURL:  "libsql://notes-mr-destructive.aws-ap-south-1.turso.io?tls=1",
			token:  "tok",
			expect: "libsql://notes-mr-destructive.aws-ap-south-1.turso.io?tls=1&authToken=tok",
		},
		{
			name:   "https url with token",
			dbURL:  "https://notes-mr-destructive.aws-ap-south-1.turso.io",
			token:  "tok",
			expect: "https://notes-mr-destructive.aws-ap-south-1.turso.io?authToken=tok",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildLibsqlConnStr(tt.dbURL, tt.token)
			if got != tt.expect {
				t.Fatalf("expected %q, got %q", tt.expect, got)
			}
		})
	}
}
