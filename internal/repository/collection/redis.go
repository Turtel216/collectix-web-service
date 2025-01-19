package collection

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Turtel216/Go-Microservice/internal/models"
	"github.com/redis/go-redis/v9"
)

var ErrNotExist = errors.New("order does not exist")

type RedisRepo struct {
	Client *redis.Client
}

type FindAllPage struct {
	Size   uint64
	Offset uint64
}

type FindResult struct {
	Collections []models.Collection
	Cursor      uint64
}

func (r *RedisRepo) Insert(ctx context.Context, collection models.Collection) error {
	data, err := json.Marshal(collection)
	if err != nil {
		return fmt.Errorf("failed to encode collection: %w", err)
	}

	key := collectionIdKey(collection.CollectionId)

	// transaction
	txn := r.Client.TxPipeline()

	res := r.Client.SetNX(ctx, key, string(data), 0)
	if err := res.Err(); err != nil {
		txn.Discard() // Failed transaction
		return fmt.Errorf("failed to set: %w", err)
	}

	if err := txn.SAdd(ctx, "collections", key).Err(); err != nil {
		txn.Discard() // Failed transaction
		return fmt.Errorf("failed to add to collections set: %w", err)
	}

	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

func (r *RedisRepo) findById(ctx context.Context, id uint64) (models.Collection, error) {
	key := collectionIdKey(id)

	value, err := r.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return models.Collection{}, ErrNotExist
	} else if err != nil {
		return models.Collection{}, fmt.Errorf("get collection: %w", err)
	}

	var collection models.Collection
	err = json.Unmarshal([]byte(value), &collection)

	return collection, nil
}

func (r *RedisRepo) DeleteById(ctx context.Context, id uint64) error {
	key := collectionIdKey(id)

	txn := r.Client.TxPipeline()

	err := txn.Del(ctx, key).Err()
	if errors.Is(err, redis.Nil) {
		txn.Discard()
		return ErrNotExist
	} else if err != nil {
		txn.Discard()
		return fmt.Errorf("get collection: %w", err)
	}

	if err := txn.SRem(ctx, "collections", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to exec: %w", err)
	}

	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

func (r *RedisRepo) Update(ctx context.Context, collection models.Collection) error {
	data, err := json.Marshal(collection)
	if err != nil {
		return fmt.Errorf("failed to encode collection: %w", err)
	}

	key := collectionIdKey(collection.CollectionId)

	res := r.Client.SetXX(ctx, key, string(data), 0)
	if err := res.Err(); err != nil {
		return fmt.Errorf("failed to set: %w", err)
	}

	return nil
}

func (r *RedisRepo) FindAll(ctx context.Context, page FindAllPage) (FindResult, error) {
	res := r.Client.SScan(ctx, "collections", page.Offset, "*", int64(page.Size))

	keys, cursor, err := res.Result()
	if err != nil {
		return FindResult{}, fmt.Errorf("failed to get collection ids: %w", err)
	}

	if len(keys) == 0 {
		return FindResult{
			Collections: []models.Collection{},
		}, nil
	}

	xs, err := r.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return FindResult{}, fmt.Errorf("failed to get collections: %w", err)
	}

	collections := make([]models.Collection, len(xs))

	for i, x := range xs {
		x := x.(string)
		var collection models.Collection

		err := json.Unmarshal([]byte(x), &collection)
		if err != nil {
			return FindResult{}, fmt.Errorf("failed to decode order json: %w", err)
		}

		collections[i] = collection
	}

	return FindResult{
		Collections: collections,
		Cursor:      cursor,
	}, nil
}

func collectionIdKey(id uint64) string {
	return fmt.Sprintf("collection:%d", id)
}
