package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/aexrun/htmx-go-templ-echo-pg/internal/models"
	"github.com/aexrun/htmx-go-templ-echo-pg/templates"
)

type Handler struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Handler {
	return &Handler{db: db}
}

func render(c echo.Context, status int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	if err := t.Render(c.Request().Context(), buf); err != nil {
		return err
	}
	return c.HTML(status, buf.String())
}

func (h *Handler) Home(c echo.Context) error {
	todos, err := h.listTodos(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return render(c, http.StatusOK, templates.Home(todos))
}

func (h *Handler) CreateTodo(c echo.Context) error {
	title := c.FormValue("title")
	if title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "title is required")
	}

	var todo models.Todo
	err := h.db.QueryRow(c.Request().Context(),
		"INSERT INTO todos (title) VALUES ($1) RETURNING id, title, completed, created_at",
		title,
	).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return render(c, http.StatusCreated, templates.TodoItem(todo))
}

func (h *Handler) ToggleTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	var todo models.Todo
	err = h.db.QueryRow(c.Request().Context(),
		"UPDATE todos SET completed = NOT completed WHERE id = $1 RETURNING id, title, completed, created_at",
		id,
	).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return render(c, http.StatusOK, templates.TodoItem(todo))
}

func (h *Handler) DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	_, err = h.db.Exec(c.Request().Context(), "DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) HealthCheck(c echo.Context) error {
	if err := h.db.Ping(c.Request().Context()); err != nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"status": "unhealthy", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}

func (h *Handler) listTodos(ctx context.Context) ([]models.Todo, error) {
	rows, err := h.db.Query(ctx, "SELECT id, title, completed, created_at FROM todos ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("query todos: %w", err)
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan todo: %w", err)
		}
		todos = append(todos, t)
	}
	return todos, nil
}
