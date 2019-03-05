package config

import "testing"

func TestLoadConfig(t *testing.T) {

	want := Config{
		ListenAddress: "localhost:8889",
		RelationalDB:  "mysql",
		RelationalDSN: "dean:Dean#168168@tcp(192.168.1.12:3306)/porker?collation=utf8mb4_unicode_ci",
	}

	got, err := LoadConfig("config.json")
	if err != nil {
		t.Fatalf("load configuration file err: %v\n", err)
	}

	if *got != want {
		t.Fatalf("got != want, got %v, want %v\n", *got, want)
	}

}
