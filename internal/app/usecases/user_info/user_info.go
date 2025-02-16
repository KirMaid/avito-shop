package usecases

import (
	"avitoshop/internal/app/entities"
	"avitoshop/internal/app/repositories"
	"context"
	"fmt"
)

type UserInfoUseCase struct {
	userRepo           repositories.UserRepository
	inventoryRepo      repositories.InventoryRepository
	transactionRepo    repositories.TransactionRepository
	goodRepo           repositories.GoodRepository
	redisUserRepo      repositories.RedisUserRepository
	redisInventoryRepo repositories.RedisInventoryRepository
	redisGoodRepo      repositories.RedisGoodRepository
	//redisTransactionRepo repositories.RedisTransactionRepository
}

func NewUserInfoUseCase(
	userRepo repositories.UserRepository,
	inventoryRepo repositories.InventoryRepository,
	transactionRepo repositories.TransactionRepository,
	goodRepo repositories.GoodRepository,
	redisUserRepo repositories.RedisUserRepository,
	redisInventoryRepo repositories.RedisInventoryRepository,
	redisGoodRepo repositories.RedisGoodRepository,
	// redisTransactionRepo repositories.RedisTransactionRepository,
) *UserInfoUseCase {
	return &UserInfoUseCase{
		userRepo:           userRepo,
		inventoryRepo:      inventoryRepo,
		transactionRepo:    transactionRepo,
		goodRepo:           goodRepo,
		redisUserRepo:      redisUserRepo,
		redisInventoryRepo: redisInventoryRepo,
		redisGoodRepo:      redisGoodRepo,
		//redisTransactionRepo: redisTransactionRepo,
	}
}

//func (s *UserInfoUseCase) GetInfo(ctx context.Context, username string) (*UserInfoDTO, error) {
//	user, err := s.redisUserRepo.GetByUsername(ctx, username)
//
//	if err != nil {
//		user, err = s.userRepo.GetByUsername(ctx, username)
//		if err != nil {
//			return nil, fmt.Errorf("failed to get user: %w", err)
//		}
//
//		_ = s.redisUserRepo.SetByUsername(ctx, username, user)
//		_ = s.redisUserRepo.SetById(ctx, user.ID, user)
//	}
//
//	inventoryEntities, err := s.redisInventoryRepo.GetByUser(ctx, user.ID)
//	if err != nil {
//		inventoryEntities, err = s.redisInventoryRepo.GetByUser(ctx, user.ID)
//		if err != nil {
//			return nil, fmt.Errorf("failed to get inventory: %w", err)
//		}
//
//		_ = s.redisInventoryRepo.SetByUser(ctx, user.ID, inventoryEntities)
//	}
//
//	inventory := make([]InventoryDTO, 0, len(inventoryEntities))
//	for _, item := range inventoryEntities {
//		inventory = append(inventory, InventoryDTO{
//			Type:     "fdssdf", //item.Type,
//			Quantity: item.Quantity,
//		})
//	}
//
//	receivedEntities, err := s.transactionRepo.GetReceivedTransactions(ctx, user.ID)
//	if err != nil {
//		return nil, fmt.Errorf("failed to get received transactions: %w", err)
//	}
//
//	sentEntities, err := s.transactionRepo.GetSentTransactions(ctx, user.ID)
//	if err != nil {
//		return nil, fmt.Errorf("failed to get sent transactions: %w", err)
//	}
//
//	userIDSet := make(map[int]struct{})
//	for _, tx := range receivedEntities {
//		userIDSet[tx.SenderID] = struct{}{}
//	}
//	for _, tx := range sentEntities {
//		userIDSet[tx.ReceiverID] = struct{}{}
//	}
//
//	userIDs := make([]int, 0, len(userIDSet))
//	for id := range userIDSet {
//		userIDs = append(userIDs, id)
//	}
//
//	//usernames, missingIds, err := s.redisUserRepo.GetUsernamesByIDs(ctx, userIDs)
//	//
//	//fmt.Printf("Чекаем лут 1 %s\n\n", usernames[1])
//	//fmt.Printf("Чекаем лут 2 %s\n\n", usernames[2])
//	//if err != nil {
//	//	if !errors.Is(err, repositories.ErrCacheMiss) {
//	//		return nil, fmt.Errorf("failed to get usernames from Redis: %w", err)
//	//	}
//	//}
//	//
//	//for _, id := range missingIds {
//	//	fmt.Printf("Гойда missingIds %d\n\n", id)
//	//}
//	//
//	//if len(missingIds) > 0 {
//	usernames := make(map[int]string, len(userIDs))
//
//	usernamesFromDb, err := s.userRepo.GetByIDs(ctx, userIDs)
//	if err != nil {
//		return nil, fmt.Errorf("failed to get usernames from Postgre: %w", err)
//	}
//
//	for _, username := range usernamesFromDb {
//		usernames[username.ID] = username.Username
//		//if err := s.redisUserRepo.SetByUsername(ctx, username.Username, &username); err != nil {
//		//	return nil, fmt.Errorf("failed to set username in Redis: %w", err)
//		//}
//	}
//	//}
//
//	received := make([]ReceivedDTO, 0, len(receivedEntities))
//	for _, tx := range receivedEntities {
//		received = append(received, ReceivedDTO{
//			FromUser: usernames[tx.SenderID],
//			Amount:   tx.Amount,
//		})
//	}
//
//	sent := make([]SentDTO, 0, len(sentEntities))
//	for _, tx := range sentEntities {
//		sent = append(sent, SentDTO{
//			ToUser: usernames[tx.ReceiverID],
//			Amount: tx.Amount,
//		})
//	}
//
//	return &UserInfoDTO{
//		Coins:     user.Balance,
//		Inventory: inventory,
//		CoinHistory: CoinHistoryDTO{
//			Received: received,
//			Sent:     sent,
//		},
//	}, nil
//}

func (s *UserInfoUseCase) GetInfo(ctx context.Context, username string) (*UserInfoDTO, error) {
	user, err := s.getUser(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	inventoryEntities, err := s.getInventory(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory: %w", err)
	}

	inventory, err := s.getInventoryDTOs(ctx, inventoryEntities)
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory DTOs: %w", err)
	}

	receivedEntities, err := s.transactionRepo.GetReceivedTransactions(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get received transactions: %w", err)
	}

	sentEntities, err := s.transactionRepo.GetSentTransactions(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get sent transactions: %w", err)
	}

	usernames, err := s.getUsernames(ctx, receivedEntities, sentEntities)
	if err != nil {
		return nil, fmt.Errorf("failed to get usernames: %w", err)
	}

	received := s.mapReceivedDTOs(receivedEntities, usernames)
	sent := s.mapSentDTOs(sentEntities, usernames)

	return &UserInfoDTO{
		Coins:     user.Balance,
		Inventory: inventory,
		CoinHistory: CoinHistoryDTO{
			Received: received,
			Sent:     sent,
		},
	}, nil
}

func (s *UserInfoUseCase) getUser(ctx context.Context, username string) (*entities.User, error) {
	user, err := s.redisUserRepo.GetByUsername(ctx, username)
	if err == nil {
		return user, nil
	}

	user, err = s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from database: %w", err)
	}

	go func() {
		_ = s.redisUserRepo.SetByUsername(ctx, username, user)
		_ = s.redisUserRepo.SetById(ctx, user.ID, user)
	}()

	return user, nil
}

func (s *UserInfoUseCase) getInventory(ctx context.Context, userID int) ([]entities.Inventory, error) {
	inventoryEntities, err := s.redisInventoryRepo.GetByUser(ctx, userID)
	if err == nil {
		return inventoryEntities, nil
	}

	inventoryEntities, err = s.inventoryRepo.GetByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory from database: %w", err)
	}

	go func() {
		_ = s.redisInventoryRepo.SetByUser(ctx, userID, inventoryEntities)
	}()

	return inventoryEntities, nil
}

func (s *UserInfoUseCase) getInventoryDTOs(ctx context.Context, inventoryEntities []entities.Inventory) ([]InventoryDTO, error) {
	inventory := make([]InventoryDTO, 0, len(inventoryEntities))

	for _, item := range inventoryEntities {
		good, err := s.getGood(ctx, item.GoodID)
		if err != nil {
			return nil, fmt.Errorf("failed to get good: %w", err)
		}

		inventory = append(inventory, InventoryDTO{
			Type:     good.Name,
			Quantity: item.Quantity,
		})
	}

	return inventory, nil
}

func (s *UserInfoUseCase) getGood(ctx context.Context, goodID int) (*entities.Good, error) {
	good, err := s.redisGoodRepo.GetByID(ctx, goodID)
	if err == nil {
		return good, nil
	}

	good, err = s.goodRepo.GetByID(ctx, goodID)
	if err != nil {
		return nil, fmt.Errorf("failed to get good from database: %w", err)
	}

	go func() {
		_ = s.redisGoodRepo.SetByID(ctx, goodID, good)
	}()

	return good, nil
}

func (s *UserInfoUseCase) getUsernames(ctx context.Context, receivedEntities, sentEntities []entities.Transaction) (map[int]string, error) {
	userIDSet := make(map[int]struct{})
	for _, tx := range receivedEntities {
		userIDSet[tx.SenderID] = struct{}{}
	}
	for _, tx := range sentEntities {
		userIDSet[tx.ReceiverID] = struct{}{}
	}

	userIDs := make([]int, 0, len(userIDSet))
	for id := range userIDSet {
		userIDs = append(userIDs, id)
	}

	usernames := make(map[int]string, len(userIDs))
	usernamesFromDb, err := s.userRepo.GetByIDs(ctx, userIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get usernames from database: %w", err)
	}

	for _, user := range usernamesFromDb {
		usernames[user.ID] = user.Username
	}

	return usernames, nil
}

func (s *UserInfoUseCase) mapReceivedDTOs(receivedEntities []entities.Transaction, usernames map[int]string) []ReceivedDTO {
	received := make([]ReceivedDTO, 0, len(receivedEntities))
	for _, tx := range receivedEntities {
		received = append(received, ReceivedDTO{
			FromUser: usernames[tx.SenderID],
			Amount:   tx.Amount,
		})
	}
	return received
}

func (s *UserInfoUseCase) mapSentDTOs(sentEntities []entities.Transaction, usernames map[int]string) []SentDTO {
	sent := make([]SentDTO, 0, len(sentEntities))
	for _, tx := range sentEntities {
		sent = append(sent, SentDTO{
			ToUser: usernames[tx.ReceiverID],
			Amount: tx.Amount,
		})
	}
	return sent
}
