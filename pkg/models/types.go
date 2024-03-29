package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	// ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();not null;primaryKey"`
	// Nickname  string    `gorm:"not null;default:null;primaryKey"`
	ID        uuid.UUID     `gorm:"not null;default:null;primaryKey"`
	Username  string        `gorm:"not null;unique;default:null"`
	Password  string        `gorm:"not null;default:null"`
	Chats     []GenericChat `gorm:"many2many:user_chats;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Chat struct {
	// ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();not null;primaryKey"`
	ID            uuid.UUID `gorm:"not null;default:null;primaryKey"`
	User1ID       uuid.UUID `gorm:"not null;default:null"`
	User2ID       uuid.UUID `gorm:"not null;default:null"`
	User1         User      `gorm:"foreignKey:User1ID"`
	User2         User      `gorm:"foreignKey:User2ID"`
	LastMessageAt time.Time `gorm:"not null;default:null"`
	LastMessage   string    `gorm:"not null;default:null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type Message struct {
	// ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();not null;primaryKey"`
	ID          uuid.UUID `gorm:"not null;default:null;primaryKey"`
	Content     string    `gorm:"not null;default:null"`
	ChatID      uuid.UUID `gorm:"not null;default:null"`
	SenderID    uuid.UUID `gorm:"not null;default:null"`
	Attachments []Attachment
	Sender      User `gorm:"foreignKey:SenderID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Attachment struct {
	// ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();not null;primaryKey"`
	ID        uuid.UUID `gorm:"not null;default:null;primaryKey"`
	Path      string    `gorm:"not null;default:null"`
	MessageID uuid.UUID `gorm:"not null;default:null"`
	Message   Message   `gorm:"foreignKey:MessageID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type GenericChat struct {
	// ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();not null;primaryKey"`
	ID            uuid.UUID `gorm:"not null;default:null;primaryKey"`
	Name          string    `gorm:"default:null"`
	LastMessage   string    `gorm:"not null;default:null"`
	LastSenderID  uuid.UUID `gorm:"not null;default:null"`
	LastMessageAt time.Time `gorm:"not null;default:null"`
	IsGroup       bool      `gorm:"not null;"`
	Users         []User    `gorm:"many2many:user_chats;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

func (chat *GenericChat) GetName(userID uuid.UUID) string {
	if chat.IsGroup {
		return chat.Name
	}

	var username string
	for _, user := range chat.Users {
		if userID != user.ID {
			username = user.Username
		}
	}

	return username
}

func (chat *GenericChat) GetLastMessage(userID uuid.UUID) string {
	if !chat.IsGroup {
		return chat.LastMessage
	}

	if userID == chat.LastSenderID {
		return fmt.Sprintf("Eu: %s", chat.LastMessage)
	}

	var username string
	for _, user := range chat.Users {
		if chat.LastSenderID == user.ID {
			username = user.Username
		}
	}

	return fmt.Sprintf("%s: %s", username, chat.LastMessage)
}

// type User struct {
// 	gorm.Model
// 	Nickname string
// 	Username string
// 	Password string
// }

// type Message struct {
// 	gorm.Model
// 	Content    string
// 	SenderID   uint
// 	ReceiverID uint
// }

// type Chat struct {
// 	ContactNickname string
// 	ContactID       string
// 	LastMessage     string
// 	LastMessageAt   time.Time
// }
