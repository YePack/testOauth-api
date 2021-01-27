package db

import (
	"github.com/gocql/gocql"
	"github.com/yepack/testOauth-api/src/client/cassandra"
	"github.com/yepack/testOauth-api/src/domain/access_token"
	"github.com/yepack/testUtils-api/rest_errors"
	"errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?,?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestErr)
	Create(access_token.AccessToken) rest_errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) rest_errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, rest_errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&result.AccessToken,
		&result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("no access token found with given id")
		}
		return nil, rest_errors.NewInternalServerError("error to get current id",
			errors.New("db error"))
	}
	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) rest_errors.RestErr {

	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId,
		at.ClientId, at.Expires).Exec(); err != nil {
		return rest_errors.NewInternalServerError("error with saving access token in db", err)
	}
	return nil
}
func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) rest_errors.RestErr {

	if err := cassandra.GetSession().Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec(); err != nil {
		return rest_errors.NewInternalServerError("error with updating current resource",
			errors.New("db error"))
	}
	return nil
}
