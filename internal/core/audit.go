package core

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrNilDB        = errors.New("db tidak boleh nil")
	ErrInvalidAudit = errors.New("action dan entity wajib diisi")
)

// AuditEntry mencatat aktivitas CRUD pada aplikasi HRIS.
type AuditEntry struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Action    string    `json:"action" gorm:"not null;index"`
	Entity    string    `json:"entity" gorm:"not null;index"`
	EntityID  int       `json:"entity_id" gorm:"not null;index"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (AuditEntry) TableName() string {
	return "audit_log"
}

// MigrateAudit membuat tabel audit_log jika belum ada.
func MigrateAudit(db *gorm.DB) error {
	return db.AutoMigrate(&AuditEntry{})
}

// LogAudit menyimpan satu entri audit.
func LogAudit(db *gorm.DB, action, entity string, entityID int, detail string) error {
	if db == nil {
		return ErrNilDB
	}
	if action == "" || entity == "" {
		return ErrInvalidAudit
	}
	return db.Create(&AuditEntry{
		Action:   action,
		Entity:   entity,
		EntityID: entityID,
		Detail:   detail,
	}).Error
}

// AuditEntries mengambil entri audit. Filter entity/entityID opsional:
// entity kosong mengembalikan semua entity, entityID 0 mengembalikan semua id.
func AuditEntries(db *gorm.DB, entity string, entityID int) ([]AuditEntry, error) {
	if db == nil {
		return nil, ErrNilDB
	}
	q := db.Model(&AuditEntry{})
	if entity != "" {
		q = q.Where("entity = ?", entity)
	}
	if entityID != 0 {
		q = q.Where("entity_id = ?", entityID)
	}
	var entries []AuditEntry
	if err := q.Order("id asc").Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}
