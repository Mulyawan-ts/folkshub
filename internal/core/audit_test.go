package core

import (
	"testing"
)

func TestAuditLogRecordAndQuery(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if err := MigrateAudit(db); err != nil {
		t.Fatalf("migrate: %v", err)
	}

	if err := LogAudit(db, "create", "karyawan", 1, "nama=Budi"); err != nil {
		t.Fatalf("log: %v", err)
	}
	if err := LogAudit(db, "update", "karyawan", 1, "nama=Budi Santoso"); err != nil {
		t.Fatalf("log: %v", err)
	}

	entries, err := AuditEntries(db, "karyawan", 1)
	if err != nil {
		t.Fatalf("query: %v", err)
	}
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
	if entries[0].Action != "create" || entries[1].Action != "update" {
		t.Fatalf("unexpected actions: %+v", entries)
	}
}

func TestAuditEntriesAll(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if err := MigrateAudit(db); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	if err := LogAudit(db, "delete", "absensi", 5, ""); err != nil {
		t.Fatalf("log: %v", err)
	}

	entries, err := AuditEntries(db, "", 0)
	if err != nil {
		t.Fatalf("query: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
}
