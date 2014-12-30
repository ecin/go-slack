package slack

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
)

type MessageResponse struct {
  Channel string
  Error string
  Ok bool
}

type User struct {
  Id string
  Name string
  Deleted bool
  Profile struct {
    Email string
  }
}

type UserResponse struct {
  Error string
  Ok bool
  User User
}

type UsersResponse struct {
  Error string
  Ok bool
  Members []User
}

// Alias for usersResponse.Members
func (response UsersResponse) Users() []User {
  return response.Members
}

const (
  POST_MESSAGE_ENDPOINT = "https://slack.com/api/chat.postMessage"
  USER_INFO_ENDPOINT = "https://slack.com/api/users.info"
  USER_LIST_ENDPOINT = "https://slack.com/api/users.list"
)

type UserRequest struct {
  Token string
  User string
}

type SlackClient struct {
  Token string
}

func NewSlackClient(token string) SlackClient {
  client := SlackClient{
    token,
  }

  return client
}

func (client SlackClient) SendMessage(channel string, text string, botName string) MessageResponse {
  params := url.Values{}
  params.Add("token", client.Token)
  params.Add("channel", channel)
  params.Add("text", text)
  params.Add("username", botName)

  endpoint := fmt.Sprintf("%s?%s", POST_MESSAGE_ENDPOINT, params.Encode())

  response, _ := http.Get(endpoint)
  body, _ := ioutil.ReadAll(response.Body)

  var messageResponse MessageResponse
  json.Unmarshal(body, &messageResponse)

  return messageResponse
}

func (client SlackClient) GetUsers() UsersResponse {
  params := url.Values{}
  params.Add("token", client.Token)

  endpoint := fmt.Sprintf("%s?%s", USER_LIST_ENDPOINT, params.Encode())

  response, _ := http.Get(endpoint)
  body, _ := ioutil.ReadAll(response.Body)

  var usersResponse UsersResponse
  json.Unmarshal(body, &usersResponse)

  return usersResponse
}

func (client SlackClient) GetUser(userId string) UserResponse {
  params := url.Values{}
  params.Add("token", client.Token)
  params.Add("user", userId)

  endpoint := fmt.Sprintf("%s?%s", USER_INFO_ENDPOINT, params.Encode())

  // TODO: handle errors
  response, _ := http.Get(endpoint)
  body, _ := ioutil.ReadAll(response.Body)

  var userResponse UserResponse
  json.Unmarshal(body, &userResponse)

  return userResponse
}
