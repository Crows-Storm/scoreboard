package adapters

import (
	"context"
	"sync"
	"time"

	domain "github.com/Crows-Storm/scoreboard/internal/user/domain/user"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type MemoryUserRepository struct {
	lock  *sync.RWMutex
	store []*domain.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	s := make([]*domain.User, 0)
	s = append(s, &domain.User{
		Id:        "apple",
		Email:     "sanderQiu@good.com",
		Name:      "Sander",
		Avatar:    "https://baidu.com/img.jpg",
		CreatedAt: time.Now().UnixMilli() - 3600_000,
		UpdatedAt: time.Now().UnixMilli(),
	})
	return &MemoryUserRepository{
		lock: &sync.RWMutex{},
		//store: make([]*domain.User, 0),
		store: s,
	}
}

func (m MemoryUserRepository) Create(_ context.Context, user *domain.User) (*domain.User, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	nowTimestamp := time.Now().UnixMilli()
	newUsers := &domain.User{
		Id:        uuid.NewString(),
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		Avatar:    user.Avatar,
		CreatedAt: nowTimestamp,
		UpdatedAt: nowTimestamp,
	}
	m.store = append(m.store, newUsers)
	logrus.WithFields(logrus.Fields{
		"input_user":         user,
		"store_after_create": m.store,
	})

	return newUsers, nil
}

func (m MemoryUserRepository) Update(ctx context.Context, user *domain.User, updateFun func(context.Context, *domain.User) (*domain.User, error)) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	for i, v := range m.store {
		if v.Id == user.Id {
			// give updateFun domain.User from store, Updates will be performed by updateFun
			updateUsers, err := updateFun(ctx, v)
			if err != nil {
				return err
			}
			m.store[i] = updateUsers
			logrus.WithFields(logrus.Fields{
				"input_user":         user,
				"store_after_create": m.store,
			})
		} else {
			return domain.NotFoundError{UserId: user.Id}
		}
	}

	return nil
}

func (m MemoryUserRepository) Get(ctx context.Context, id string) (*domain.User, error) {
	// xx.getById(id)
	m.lock.RLock()
	defer m.lock.RUnlock()

	for _, v := range m.store {
		if v.Id == id {
			// reset password
			v.Password = ""
			logrus.Debugf("memory_user_repo_get||id=%s||res=%+v", id, *v)
			return v, nil
		}
	}
	return nil, domain.NotFoundError{
		UserId: id,
	}
}

func (m MemoryUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for _, v := range m.store {
		if v.Name == username {
			logrus.Debugf("memory_user_repo_get||username=%s||res=%+v", username, *v)
			return v, nil
		}
	}
	return nil, domain.NotFoundError{
		UserId: "getUser by username: " + username,
	}
}
