package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
)

type messageType string

const (
	MESSAGE messageType = "MESSAGE" // Change to UpperCase to difference from Message struct
	LIKE    messageType = "LIKE"
)

type Message struct {
	// Id        string            `json:"id"` // Not needed for the data set
	UserId      string            `json:"UserId"`
	Message     MessageValue      `json:"Message"` // change to MessageValue to unmarshal object value
	Type        messageType       `json:"Type"`
	HasRead     map[string]string `json:"HasRead"`
	TimeStamp   int64             `json:"TimeStamp"`
	IsDeleted   bool              `json:"IsDeleted"`
	Collections map[string]string `json:"__collections__"`
}

type MessageValue struct {
	DataType string `json:"__datatype__"`
	Value    string `json:"value"`
}

func mergeMessages(messagesA, messagesB []Message) []Message {
	merged := append(messagesA, messagesB...)
	sort.Slice(merged, func(i, j int) bool {
		return merged[i].TimeStamp < merged[j].TimeStamp
	})
	return merged
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	dataSetA, err := os.ReadFile("dataSetA.json")
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
		return
	}

	dataSetB, err := os.ReadFile("dataSetB.json")
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
		return
	}

	var messagesA, messagesB []Message
	if err := json.Unmarshal([]byte(dataSetA), &messagesA); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal([]byte(dataSetB), &messagesB); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Merge the messages
	merged := mergeMessages(messagesA, messagesB)

	// Convert the merged messages to JSON
	mergedJSON, err := json.Marshal(merged)
	if err != nil {
		http.Error(w, "Failed to marshal merged messages", http.StatusInternalServerError)
		return
	}

	// Set the response content type and write the merged JSON to the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(mergedJSON)
}

func main() {
	http.HandleFunc("/messages", handleMessages)
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
