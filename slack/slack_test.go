package slack

import (
  "os"
  "testing"
)

func TestGetUser(t *testing.T) {
  token := os.Getenv("SLACK_TOKEN")
  client := NewSlackClient(token)

  response := client.GetUser("fake")

  if response.Error != "user_not_found" {
    t.Errorf("Expected fake user id to generate error, got %v", response)
  }

  userId := "U0389942F"
  response = client.GetUser(userId)

  if response.Error != "" {
    t.Errorf("Didn't expect errors when fetching user %s, got %v", userId, response)
  }
}

func TestGetUsers(t *testing.T) {
  token := os.Getenv("SLACK_TOKEN")
  client := NewSlackClient(token)

  response := client.GetUsers()

  if response.Error != "" {
    t.Errorf("Didn't expect errors when fetching users, got %v", response)
  }
}

func TestSendMessage(t *testing.T) {
  token := os.Getenv("SLACK_TOKEN")
  client := NewSlackClient(token)

  response := client.SendMessage("#general", "Hello from Golang!", "slackash")

  if response.Error != "" {
    t.Errorf("Didn't expect errors when sending message, got %v", response)
  }

}
