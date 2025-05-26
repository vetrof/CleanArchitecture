package service

import "testing"

import (
	"CleanArchitecture/repo"
)

func TestCreateAndListUsers(t *testing.T) {
	// Используем фейковый репозиторий
	fakeRepo := repo.NewFakeUserRepository()

	// Создаем сервис с фейком
	service := NewUserService(fakeRepo)

	// Добавляем пользователей
	service.CreateUser("Alice")
	service.CreateUser("Bob")

	// Получаем список
	users := service.ListUsers()

	// Проверяем
	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
	if users[0].Name != "Alice" || users[1].Name != "Bob" {
		t.Errorf("unexpected users: %+v", users)
	}
}
