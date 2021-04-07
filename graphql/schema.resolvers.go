package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gqllasttry/graphql/generated"
	"gqllasttry/graphql/model"
	"gqllasttry/postgres"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (r *chatResolver) Members(ctx context.Context, obj *model.Chat) ([]*model.User, error) {
	chatID, err := uuid.Parse(obj.ID)
	if err != nil {
		return nil, err
	}

	dbmembers, err := r.DB.GetChatMembers(ctx, chatID)
	if err != nil {
		return nil, err
	}

	var result []*model.User
	for _, dbuser := range dbmembers {
		result = append(result, &model.User{
			ID:          dbuser.UserID.String(),
			Username:    dbuser.Username,
			DateCreated: dbuser.DateCreated,
			DateUpdated: dbuser.DateUpdated,
		})
	}

	return result, nil
}

func (r *chatResolver) MembersCount(ctx context.Context, obj *model.Chat) (int, error) {
	chatID, err := uuid.Parse(obj.ID)
	if err != nil {
		return 0, err
	}

	count, err := r.DB.CountChatMembers(ctx, chatID)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *chatResolver) Messages(ctx context.Context, obj *model.Chat) ([]*model.Message, error) {
	chatID, err := uuid.Parse(obj.ID)
	if err != nil {
		return nil, err
	}

	dbmessages, err := r.DB.GetChatMessages(ctx, chatID)
	if err != nil {
		return nil, err
	}

	var result []*model.Message
	for _, dbmessage := range dbmessages {
		result = append(result, &model.Message{
			ID:          dbmessage.MessageID.String(),
			Text:        dbmessage.Content,
			SenderID:    dbmessage.UserID.String(),
			ChatID:      dbmessage.ChatID.String(),
			DateCreated: dbmessage.DateCreated,
		})
	}

	return result, nil
}

func (r *chatResolver) MessageCount(ctx context.Context, obj *model.Chat) (int, error) {
	chatID, err := uuid.Parse(obj.ID)
	if err != nil {
		return 0, err
	}

	count, err := r.DB.CountChatMessages(ctx, chatID)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *messageResolver) Sender(ctx context.Context, obj *model.Message) (*model.User, error) {
	userID, err := uuid.Parse(obj.SenderID)
	if err != nil {
		return nil, err
	}

	dbuser, err := r.DB.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:          dbuser.UserID.String(),
		Username:    dbuser.Username,
		DateCreated: dbuser.DateCreated,
		DateUpdated: dbuser.DateUpdated,
	}, nil
}

func (r *messageResolver) Chat(ctx context.Context, obj *model.Message) (*model.Chat, error) {
	chatID, err := uuid.Parse(obj.ChatID)
	if err != nil {
		return nil, err
	}

	dbchat, err := r.DB.GetChatByID(ctx, chatID)
	if err != nil {
		return nil, err
	}

	return &model.Chat{
		ID:          dbchat.ChatID.String(),
		Title:       dbchat.Title,
		DateCreated: dbchat.DateCreated,
		DateUpdated: dbchat.DateUpdated,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.CreateUserInput) (*model.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := r.DB.CreateUser(ctx, postgres.CreateUserParams{
		Username: input.Username,
		Password: string(hashedPass),
	})
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:          user.UserID.String(),
		Username:    user.Username,
		DateCreated: user.DateCreated,
		DateUpdated: user.DateUpdated,
	}, nil
}

func (r *mutationResolver) SendMessage(ctx context.Context, input *model.SendMessageInput) (*model.Message, error) {
	senderID, err := uuid.Parse(input.SenderID)
	if err != nil {
		return nil, err
	}

	chatID, err := uuid.Parse(input.ChatID)
	if err != nil {
		return nil, err
	}

	dbmessage, err := r.DB.CreateMessage(ctx, postgres.CreateMessageParams{
		UserID:  senderID,
		ChatID:  chatID,
		Content: input.Text,
	})
	if err != nil {
		return nil, err
	}

	return &model.Message{
		ID:          dbmessage.MessageID.String(),
		Text:        dbmessage.Content,
		SenderID:    dbmessage.UserID.String(),
		ChatID:      dbmessage.ChatID.String(),
		DateCreated: dbmessage.DateCreated,
	}, nil
}

func (r *mutationResolver) CreateChat(ctx context.Context, input *model.CreateChatPayload) (*model.Chat, error) {
	dbchat, err := r.DB.CreateChat(ctx, input.Title)
	if err != nil {
		return nil, err
	}

	return &model.Chat{
		ID:          dbchat.ChatID.String(),
		Title:       dbchat.Title,
		DateCreated: dbchat.DateCreated,
		DateUpdated: dbchat.DateUpdated,
	}, nil
}

func (r *mutationResolver) JoinChat(ctx context.Context, input *model.JoinChatPayload) (*model.Chat, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, err
	}

	chatID, err := uuid.Parse(input.ChatID)
	if err != nil {
		return nil, err
	}

	err = r.DB.CreateChatMember(ctx, postgres.CreateChatMemberParams{
		UserID: userID,
		ChatID: chatID,
	})
	if err != nil {
		return nil, err
	}

	dbchat, err := r.DB.GetChatByID(ctx, chatID)
	if err != nil {
		return nil, err
	}

	return &model.Chat{
		ID:          dbchat.ChatID.String(),
		Title:       dbchat.Title,
		DateCreated: dbchat.DateCreated,
		DateUpdated: dbchat.DateUpdated,
	}, nil
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	dbusers, err := r.DB.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var result []*model.User
	for _, dbuser := range dbusers {
		result = append(result, &model.User{
			ID:          dbuser.UserID.String(),
			Username:    dbuser.Username,
			DateCreated: dbuser.DateCreated,
			DateUpdated: dbuser.DateUpdated,
		})
	}

	return result, nil
}

func (r *queryResolver) AllMessages(ctx context.Context) ([]*model.Message, error) {
	dbmessages, err := r.DB.GetAllMessages(ctx)
	if err != nil {
		return nil, err
	}

	var result []*model.Message
	for _, dbmessage := range dbmessages {
		result = append(result, &model.Message{
			ID:          dbmessage.MessageID.String(),
			Text:        dbmessage.Content,
			SenderID:    dbmessage.UserID.String(),
			ChatID:      dbmessage.ChatID.String(),
			DateCreated: dbmessage.DateCreated,
		})
	}

	return result, nil
}

func (r *queryResolver) AllChats(ctx context.Context) ([]*model.Chat, error) {
	dbchats, err := r.DB.GetAllChats(ctx)
	if err != nil {
		return nil, err
	}

	var result []*model.Chat
	for _, dbchat := range dbchats {
		result = append(result, &model.Chat{
			ID:          dbchat.ChatID.String(),
			Title:       dbchat.Title,
			DateCreated: dbchat.DateCreated,
			DateUpdated: dbchat.DateUpdated,
		})
	}

	return result, nil
}

func (r *userResolver) MessagesCount(ctx context.Context, obj *model.User) (int, error) {
	userID, err := uuid.Parse(obj.ID)
	if err != nil {
		return 0, err
	}

	count, err := r.DB.CountUserMessages(ctx, userID)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

// Chat returns generated.ChatResolver implementation.
func (r *Resolver) Chat() generated.ChatResolver { return &chatResolver{r} }

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type chatResolver struct{ *Resolver }
type messageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
