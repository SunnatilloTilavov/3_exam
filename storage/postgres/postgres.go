package postgres

import (
 "context"
 "fmt"
 "clone/rent_car_us/config"
 "clone/rent_car_us/storage"
 "clone/rent_car_us/storage/redis"
 "time"

 _ "github.com/lib/pq"

 "github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
 Pool *pgxpool.Pool
 cfg    config.Config
redis  storage.IRedisStorage
}

func New(ctx context.Context, cfg config.Config,redis storage.IRedisStorage) (storage.IStorage, error) {
 url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
  cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

 pgPoolConfig, err := pgxpool.ParseConfig(url)
 if err != nil {
  return nil, err
 }

 pgPoolConfig.MaxConns = 100
 pgPoolConfig.MaxConnLifetime = time.Hour

 newPool, err := pgxpool.NewWithConfig(context.Background(), pgPoolConfig)
 if err != nil {
  fmt.Println("error while connecting to db", err.Error())
  return nil, err
 }

 return Store{
  Pool: newPool,
  cfg:    cfg,
  redis:  redis,
 }, nil
}

func (s Store) CloseDB() {
 s.Pool.Close()
}

func (s Store) Car() storage.ICarStorage {
 newCar := NewCar(s.Pool)

 return &newCar
}

func (s Store) Customer() storage.ICustomerStorage {
 newCustomer := NewCustomer(s.Pool)

 return &newCustomer
}

func (s Store) Order() storage.IOrderStorage {
 NewOrder := NewOrder(s.Pool)

 return &NewOrder
}

func (s Store) Redis() storage.IRedisStorage {
	return redis.New(s.cfg)
}

