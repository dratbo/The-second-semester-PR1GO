package service

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	DueDate     string    `json:"due_date,omitempty"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Storage struct {
	mu    sync.RWMutex
	tasks map[string]Task
}

func NewStorage() *Storage {
	return &Storage{
		tasks: make(map[string]Task),
	}
}

func (s *Storage) Create(task Task) Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = uuid.New().String()[:8]
	task.CreatedAt = time.Now()
	task.UpdatedAt = task.CreatedAt

	s.tasks[task.ID] = task
	return task
}

func (s *Storage) List() []Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}

func (s *Storage) Get(id string) (Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.tasks[id]
	return t, ok
}

func (s *Storage) Update(id string, updated Task) (Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.tasks[id]
	if !ok {
		return Task{}, false
	}

	if updated.Title != "" {
		t.Title = updated.Title
	}
	if updated.Description != "" {
		t.Description = updated.Description
	}
	if updated.DueDate != "" {
		t.DueDate = updated.DueDate
	}
	t.Done = updated.Done
	t.UpdatedAt = time.Now()

	s.tasks[id] = t
	return t, true
}

func (s *Storage) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.tasks[id]
	if ok {
		delete(s.tasks, id)
	}
	return ok
}
