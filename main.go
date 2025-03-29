package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// This function simulates a slow database call
func fetchSlowData(ctx context.Context) (string, error) {
	select {
	case <-time.After(5 * time.Second):
		// Simulates the slow operation finally returning
		return "✅ Slow data loaded", nil
	case <-ctx.Done():
		// Context expired or canceled
		return "", ctx.Err()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Step 1: Set a timeout context of 2 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel() // always cancel to release resources

	// Step 2: Try to fetch slow data
	data, err := fetchSlowData(ctx)
	if err != nil {
		http.Error(w, "❌ Request canceled or timed out: "+err.Error(), http.StatusRequestTimeout)
		return
	}

	// Step 3: Send result if successful
	fmt.Fprintln(w, data)
}

func main() {
	http.HandleFunc("/data", handler)
	fmt.Println("🌐 Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
