package service

import (
	"fmt"
	"mingin/model"
	"sync"
)

func ListUser(offset, limit int) ([]*model.UserInfo, int64, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(offset, limit)
	if err != nil {
		return nil, count, err
	}
	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserInfo, len(users)),
	}

	errChan := make(chan error, 1)
	finChan := make(chan bool, 1)

	// 并发加快数据处理速度
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			if err != nil {
				errChan <- err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.Id] = &model.UserInfo{
				ID:    u.Id,
				Email: u.Email,
				// Phone:     u.Phone,
				Nickname:  u.Nickname,
				Role:      u.Role,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
				SayHello:  fmt.Sprintf("Hello %s", u.Nickname),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finChan)
	}()

	select {
	case <-finChan:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}
