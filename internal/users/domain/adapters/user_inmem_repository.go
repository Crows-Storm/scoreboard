package adapters

import (
	"context"
	"sync"
	"time"

	domain "github.com/Crows-Storm/scoreboard/internal/users/domain/users"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type MemoryUserRepository struct {
	lock  *sync.RWMutex
	store []*domain.Users
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		lock:  &sync.RWMutex{},
		store: make([]*domain.Users, 0),
	}
}

func (m MemoryUserRepository) Create(_ context.Context, users *domain.Users) (*domain.Users, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	nowTimestamp := time.Now().UnixMilli()
	newUsers := &domain.Users{
		Id:        uuid.NewString(),
		Email:     users.Email,
		Name:      users.Name,
		Avatar:    users.Avatar,
		CreatedAt: nowTimestamp,
		UpdatedAt: nowTimestamp,
	}
	m.store = append(m.store, newUsers)
	logrus.WithFields(logrus.Fields{
		"input_users":        users,
		"store_after_create": m.store,
	})

	return newUsers, nil
}

func (m MemoryUserRepository) Update(ctx context.Context, users *domain.Users, updateFun func(context.Context, *domain.Users) (*domain.Users, error)) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	for i, v := range m.store {
		if v.Id == users.Id {
			//nowTimestamp := time.Now().UnixMilli()
			//updateUsers := &domain.Users{
			//	Id:        users.Id,
			//	Email:     users.Email,
			//	Name:      users.Name,
			//	Avatar:    users.Avatar,
			//	CreatedAt: users.CreatedAt,
			//	UpdatedAt: nowTimestamp,
			//}
			updateUsers, err := updateFun(ctx, users)
			if err != nil {
				return err
			}
			m.store[i] = updateUsers
			logrus.WithFields(logrus.Fields{
				"input_users":        users,
				"store_after_create": m.store,
			})
		} else {
			return domain.NotFoundError{UserId: users.Id}
		}
	}

	return nil
}

func (m MemoryUserRepository) Get(id string) (*domain.Users, error) {
	// xx.getById(id)
	m.lock.RLock()
	defer m.lock.RUnlock()

	for _, v := range m.store {
		if v.Id == id {
			logrus.Debugf("memory_users_repo_get||id=%s||res=%+v", id, *v)
			return v, nil
		}
	}
	return nil, domain.NotFoundError{
		UserId: id,
	}
}
