package menu

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Conversation struct {
	Uuid     uuid.UUID
	Messages []ChatMessage
}

type ConversationRepository interface {
	SaveConversation(c *Conversation) error
	GetConversation(uuid uuid.UUID) (*Conversation, error)
}

func GetConversationRepository(db *sql.DB) ConversationRepository {
	return NewSqliteConversationRepository(db)
}

type SqliteConversationRepository struct {
	db *sql.DB
}

func NewSqliteConversationRepository(db *sql.DB) *SqliteConversationRepository {
	return &SqliteConversationRepository{db: db}
}

func (r *SqliteConversationRepository) CreateTableIfNotExists() error {
	_, err := r.db.Exec(`
	CREATE TABLE IF NOT EXISTS conversation (
		uuid TEXT PRIMARY KEY,
		messages TEXT
	);`)

	if err != nil {
		return err
	}

	return nil
}

func (r *SqliteConversationRepository) SaveConversation(c *Conversation) error {
	err := r.CreateTableIfNotExists()
	if err != nil {
		return err
	}

	strMessages, err := json.Marshal(c.Messages)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		"INSERT INTO conversation (uuid, messages) VALUES (?, ?)",
		c.Uuid.String(),
		strMessages,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *SqliteConversationRepository) GetConversation(uuid uuid.UUID) (*Conversation, error) {
	err := r.CreateTableIfNotExists()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query("SELECT uuid, messages FROM conversation WHERE uuid = ?", uuid.String())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var c Conversation
		var messages string

		err := rows.Scan(&c.Uuid, &messages)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(messages), &c.Messages)
		if err != nil {
			return nil, err
		}

		return &c, nil
	}

	return nil, nil
}
