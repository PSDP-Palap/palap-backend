package config

import (
	"log"

	"github.com/supabase-community/supabase-go"
)

func InitializeSupabase() *supabase.Client {
	client, err := supabase.NewClient("", "", &supabase.ClientOptions{})

	if err != nil {
		log.Printf("Failed to initailize the client: %v", err)
	}

	return client
}
