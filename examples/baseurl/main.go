package main

import (
	"context"
	"fmt"
	"log"

	mcp "github.com/leefowlercu/go-mcp-registry/mcp"
)

func main() {
	fmt.Println("Configurable Base URL Example")
	fmt.Println("==============================")

	// Example 1: Default client (uses the official registry)
	fmt.Println("\n1. Using default client (official registry):")
	defaultClient, err := mcp.NewClient(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Base URL: %s\n", defaultClient.BaseURL.String())

	// Try to list a few servers to verify it works
	ctx := context.Background()
	servers, _, err := defaultClient.Servers.List(ctx, &mcp.ServerListOptions{
		ListOptions: mcp.ListOptions{Limit: 3},
	})
	if err != nil {
		fmt.Printf("Error listing servers: %v\n", err)
	} else {
		fmt.Printf("Successfully found %d servers\n", len(servers.Servers))
	}

	// Example 2: Custom base URL
	fmt.Println("\n2. Using custom base URL:")
	customClient, err := mcp.NewClient(nil, mcp.WithBaseURL("https://my-registry.example.com"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Base URL: %s\n", customClient.BaseURL.String())

	// Example 3: Custom base URL without trailing slash (will be added automatically)
	fmt.Println("\n3. Using custom base URL without trailing slash:")
	noSlashClient, err := mcp.NewClient(nil, mcp.WithBaseURL("https://api.example.com/registry"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Base URL: %s\n", noSlashClient.BaseURL.String())

	// Example 4: Custom base URL with path
	fmt.Println("\n4. Using custom base URL with path:")
	pathClient, err := mcp.NewClient(nil, mcp.WithBaseURL("https://api.example.com/v1/mcp"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Base URL: %s\n", pathClient.BaseURL.String())

	// Example 5: Error handling - invalid URL
	fmt.Println("\n5. Error handling with invalid URL:")
	_, err = mcp.NewClient(nil, mcp.WithBaseURL("ftp://invalid.com"))
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// Example 6: Error handling - empty URL
	fmt.Println("\n6. Error handling with empty URL:")
	_, err = mcp.NewClient(nil, mcp.WithBaseURL(""))
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	fmt.Println("\n✅ Configurable base URL example completed successfully!")
}