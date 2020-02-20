package model

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/verkatech/neko/pkg/config"
	"github.com/verkatech/neko/pkg/errors"
	"github.com/verkatech/neko/pkg/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Token is where we hold our tokens and token's owner id
type Token struct {
	ObjectId *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Value    string              `bson:"value,omitempty" json:"value"`
	User     *primitive.ObjectID `bson:"userId,omitempty" json:"-"`
	Created  *time.Time          `bson:"created,omitempty" json:"created"`
	Expires  *time.Time          `bson:"expires,omitempty" json:"expires"`
}

// Create generates a new unique for the specified user
func (token *Token) Create() error {

	randomString := strings.RandomString{}
	randomString.IncludeLowercaseAlphabet()
	randomString.IncludeUppercaseAlphabet()
	randomString.Generate(30)
	token.Value = randomString.Value

	now := time.Now()
	tokenttl := time.Hour * time.Duration(config.Server.TokenTimeToLive)
	expires := time.Now().Add(tokenttl)
	token.Created = &now
	token.Expires = &expires

	tokens := m.Collection(mongoTokensCollection)
	res, insertErr := tokens.InsertOne(context.Background(), token)
	if insertErr != nil {
		log.Info("failed inserting new token into mongo db")
		return errors.New("insertingTokenFailed", "Failed inserting new token", insertErr)
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		token.ObjectId = &oid
	}
	log.Info("successfully inserted new token into mongo db")
	return nil
}

func GetTokenUser(value string) (*User, error) {
	log.WithFields(log.Fields{"value": value}).Debug("looking if token with such value exists")

	var token Token
	tokens := m.Collection(mongoTokensCollection)

	err := tokens.FindOne(context.Background(), bson.M{"value": value}).Decode(&token)
	if err != nil {
		log.WithFields(log.Fields{"value": value}).Info("token with such value doesn't exists")
		return nil, errors.NotFoundError("token with this value doesn't exists")
	}

	user, err := GetUserById(token.User)

	if err != nil {
		log.WithFields(log.Fields{"token": token}).Info("user associated with this token doesn't exists")
		return nil, errors.NotFoundError("token with this value doesn't exists")
	}

	log.WithFields(log.Fields{"token": token}).Info("successfully found user associated with this token")
	return user, nil
}

const (
	mongoTokensCollection string = "tokens"
	mongoTokensValueIndex string = "uniqueValueTokensIndex"
)
