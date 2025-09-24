package main

import "testing"

func TestUrlToDirHttp(t *testing.T) {
	got := urlToDir("https://github.com/jasperges/gclone.git")
	want := "github.com/jasperges/gclone"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestUrlToDirHttpNoGit(t *testing.T) {
	got := urlToDir("https://github.com/jasperges/gclone")
	want := "github.com/jasperges/gclone"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestUrlToDirSsh(t *testing.T) {
	got := urlToDir("git@github.com:jasperges/gclone.git")
	want := "github.com/jasperges/gclone"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestUrlToDirSshNoGit(t *testing.T) {
	got := urlToDir("git@github.com:jasperges/gclone")
	want := "github.com/jasperges/gclone"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
