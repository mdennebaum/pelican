// handler.go

package server

import (
  "fmt"
  "sync"
  "github.com/mdennebaum/pelican/user"
)

type UserHandler struct {
  users map[string]*user.User
  sync.RWMutex
}

func NewUserHandler() *UserHandler {
  return &UserHandler{
    users: make(map[string]*user.User),
  }
}

func (uh *UserHandler) Create(user *user.User) (*user.User, error) {
  uh.Lock()
  defer uh.Unlock()
  uh.users[user.Id] = user
  return user, nil
}

func (uh *UserHandler) Read(userId string) (*user.User, error) {
    user, ok := uh.users[userId]
    if !ok {
        return nil, fmt.Errorf("user with ID '%s' does not exist", userId)
    }
    return user, nil
}

func (uh *UserHandler) Update(user *user.User) (*user.User, error) {
    uh.Lock()
    defer uh.Unlock()
    uh.users[user.Id] = user
    return user, nil
}

func (uh *UserHandler) Destroy(userId string) error {
    if _, ok := uh.users[userId]; ok {
        uh.Lock()
        defer uh.Unlock()
        delete(uh.users, userId)
    }
    return nil
}

func (uh *UserHandler) Fetch() ([]*user.User, error) {
    var users []*user.User
    for _, user := range uh.users {
        users = append(users, user)
    }
    return users, nil
}

func (uh *UserHandler) Reset() error {
    uh.Lock()
    defer uh.Unlock()
    uh.users = make(map[string]*user.User)
    return nil
}
