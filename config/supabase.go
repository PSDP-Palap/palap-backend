package config

import (
	"log"

	"github.com/supabase-community/supabase-go"
)

func InitializeSupabase() *supabase.Client {
	client, err := supabase.NewClient("https://wexlbxkxfduncbbzbyam.supabase.co", "sb_publishable_rpvjaCWmgOf_BWAYsMhwuQ_9nKcMZra", &supabase.ClientOptions{})

	if err != nil {
		log.Printf("Failed to initailize the client: %v", err)
	}

	return client
}
