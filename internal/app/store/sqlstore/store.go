package sqlstore

import (
	"database/sql"

	"github.com/JustVlad124/EWallet/internal/app/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db                *sql.DB
	walletRepository  *WalletRepository
	historyRepository *HistoryRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Wallet() store.WalletRepository {
	if s.walletRepository != nil {
		return s.walletRepository
	}

	s.walletRepository = &WalletRepository{
		stort: s,
	}

	return s.walletRepository
}

func (s *Store) History() store.HistoryRepository {
	if s.historyRepository != nil {
		return s.historyRepository
	}

	s.historyRepository = &HistoryRepository{
		store: s,
	}

	return s.historyRepository
}
