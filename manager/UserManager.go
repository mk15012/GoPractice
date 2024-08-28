package manager

import (
	"GoPractice/dao"
	"GoPractice/entry"
	"GoPractice/models"
	"github.com/sirupsen/logrus"
)

type UserManager struct {
	DAO *dao.User
	Log *logrus.Logger
}

func NewUserManager(dao *dao.User, log *logrus.Logger) *UserManager {
	return &UserManager{DAO: dao, Log: log}
}

func (m *UserManager) GetUserByID(id int64) (*models.Users, error) {
	return m.DAO.GetUserByID(id)
}

func (m *UserManager) CreateUser(entry entry.UserEntry) (*models.Users, error) {
	return m.DAO.CreateUser(entry)
}
