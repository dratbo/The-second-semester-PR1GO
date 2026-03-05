package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"tech-ip-sem2/services/tasks/internal/client/authclient"
	"tech-ip-sem2/services/tasks/internal/service"
)

type TaskHandlers struct {
	authClient *authclient.Client
	storage    *service.Storage
}

func NewTaskHandlers(authClient *authclient.Client) *TaskHandlers {
	return &TaskHandlers{
		authClient: authClient,
		storage:    service.NewStorage(),
	}
}

func getTokenFromRequest(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", authclient.ErrUnauthorized
	}
	parts := strings.Split(auth, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", authclient.ErrUnauthorized
	}
	return parts[1], nil
}

func (h *TaskHandlers) verifyAuth(w http.ResponseWriter, r *http.Request) (string, bool) {
	token, err := getTokenFromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized: missing or invalid token", http.StatusUnauthorized)
		return "", false
	}

	requestID := r.Header.Get("X-Request-ID")

	subject, err := h.authClient.VerifyToken(r.Context(), token, requestID)
	if err != nil {
		switch err {
		case authclient.ErrUnauthorized:
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
		case authclient.ErrAuthServiceUnavailable:
			http.Error(w, "Authorization service unavailable", http.StatusServiceUnavailable)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return "", false
	}
	return subject, true
}

func (h *TaskHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	if _, ok := h.verifyAuth(w, r); !ok {
		return
	}

	var task service.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	created := h.storage.Create(task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *TaskHandlers) ListTasks(w http.ResponseWriter, r *http.Request) {
	if _, ok := h.verifyAuth(w, r); !ok {
		return
	}

	tasks := h.storage.List()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandlers) GetTask(w http.ResponseWriter, r *http.Request) {
	if _, ok := h.verifyAuth(w, r); !ok {
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/v1/tasks/")
	if id == "" {
		http.Error(w, "Task ID required", http.StatusBadRequest)
		return
	}

	task, ok := h.storage.Get(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	if _, ok := h.verifyAuth(w, r); !ok {
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/v1/tasks/")
	if id == "" {
		http.Error(w, "Task ID required", http.StatusBadRequest)
		return
	}

	var updates service.Task
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	updated, ok := h.storage.Update(id, updates)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (h *TaskHandlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if _, ok := h.verifyAuth(w, r); !ok {
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/v1/tasks/")
	if id == "" {
		http.Error(w, "Task ID required", http.StatusBadRequest)
		return
	}

	if ok := h.storage.Delete(id); !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
