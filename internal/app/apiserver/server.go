package apiserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JustVlad124/EWallet/internal/app/models"
	"github.com/JustVlad124/EWallet/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/api/v1/wallet", s.handleWalletCreate()).Methods("POST")
	s.router.HandleFunc("/api/v1/wallet/{walletId}", s.handleWalletFindByID()).Methods("GET")
	s.router.HandleFunc("/api/v1/wallet/{walletId}/send", s.handleHistoryTransferTo()).Methods("POST")
	s.router.HandleFunc("/api/v1/wallet/{walletId}/history", s.handleHistoryFindAllTransfers()).Methods("GET")
}

func (s *server) handleWalletCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wallet := &models.Wallet{
			Balance: 100.0,
		}
		if err := s.store.Wallet().Create(wallet); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, wallet)
	}
}

func (s *server) handleWalletFindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		wallet, err := s.store.Wallet().FindByID(vars["walletId"])
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w, r, http.StatusOK, wallet)
	}
}

func (s *server) handleHistoryTransferTo() http.HandlerFunc {
	type request struct {
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		vars := mux.Vars(r)

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if req.Amount < 0 {
			s.respond(w, r, http.StatusBadRequest, "Amount must be more than 0")
			return
		}

		// если исходящий кошелек не найден
		fromWallet, err := s.store.Wallet().FindByID(vars["walletId"])
		if err != nil {
			s.respond(w, r, http.StatusNotFound, "Исходящий кошелек не найден")
			return
		}

		// если целевой кошеле не найден или на исходящем нет нужной суммы
		toWallet, err := s.store.Wallet().FindByID(req.To)
		if err != nil || fromWallet.Balance-req.Amount < 0 {
			s.respond(w, r, http.StatusNotFound, "Ошибка в пользовательском запросе или ошибка перевода")
			return
		}

		params := map[string]interface{}{"balance": fromWallet.Balance - req.Amount}
		if err := s.store.Wallet().Update(fromWallet, params); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		params = map[string]interface{}{"balance": toWallet.Balance + req.Amount}
		if err := s.store.Wallet().Update(toWallet, params); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		transferData := &models.History{
			Time:   time.Now().Format(time.RFC3339),
			From:   fromWallet.ID,
			To:     toWallet.ID,
			Amount: req.Amount,
		}
		if err := s.store.History().TransferTo(transferData); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, "Перевод успешно совершен")
	}
}

func (s *server) handleHistoryFindAllTransfers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		transfers, err := s.store.History().FindAllTransferByID(vars["walletId"])
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w, r, http.StatusOK, transfers)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
