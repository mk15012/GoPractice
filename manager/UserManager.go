package manager

import (
	"GoPractice/database"
	"GoPractice/entry"
	"GoPractice/models"
	"github.com/sirupsen/logrus"
)

type UserManager struct {
	DAO *database.User
	Log *logrus.Logger
}

func NewUserManager(dao *database.User, log *logrus.Logger) *UserManager {
	return &UserManager{DAO: dao, Log: log}
}

func (m *UserManager) GetUserByID(id int64) (*models.Users, error) {
	return m.DAO.GetUserByID(id)
}

func (m *UserManager) CreateUser(entry entry.UserEntry) (*models.Users, error) {
	return m.DAO.CreateUser(entry)
}
