package analytics_test

import (
	"context"
	"teknologi-umum-bot/analytics"
	"testing"
	"time"
)

func TestGetAllUserID(t *testing.T) {
	defer Cleanup(DB, Redis)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	err := Redis.SAdd(ctx, "analytics:users", "Adam", "Bobby", "Clifford").Err()
	if err != nil {
		t.Error(err)
	}

	deps := &analytics.Dependency{
		Redis: Redis,
	}

	users, err := deps.GetAllUserID(ctx)
	if err != nil {
		t.Error(err)
	}

	if len(users) != 3 {
		t.Error("Expected 3 users, got ", len(users))
	}
}

func TestGetAllUserMap(t *testing.T) {
	defer Cleanup(DB, Redis)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	tx := Redis.TxPipeline()
	defer tx.Close()
	tx.SAdd(ctx, "analytics:users", "1", "2", "3")
	tx.HSet(ctx, "analytics:1", "username", "adam", "display_name", "Adam", "counter", 1)
	tx.HSet(ctx, "analytics:2", "username", "bobby45", "display_name", "Bobby", "counter", 5)
	tx.HSet(ctx, "analytics:3", "username", "clifford77", "display_name", "Clifford", "counter", 3)

	_, err := tx.Exec(ctx)
	if err != nil {
		t.Error(err)
	}

	deps := &analytics.Dependency{
		Redis: Redis,
	}

	users, err := deps.GetAllUserMap(ctx)
	if err != nil {
		t.Error(err)
	}

	if len(users) != 3 {
		t.Error("Expected 3 users, got ", len(users))
	}
}
