package command

import (
	"context"
	"crypto/rand"
	db "douyin-easy/cmd/user/dal"
	"douyin-easy/grpc_gen/user"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type CreateUserService struct {
	ctx context.Context
}

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.CreateUserRequest, argon2Params *Argon2Params) (*db.User, error) {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if len(users) != 0 {
		return nil, errors.New("username is null")
	}

	// Pass the plaintext password and parameters to our generateFromPassword
	// helper function.
	passWord, err := generateFromPassword(req.Password, argon2Params)
	if err != nil {
		return nil, err
	}
	user, err := db.CreateUser(s.ctx, &db.User{
		Username: req.Username,
		Password: passWord,
	})
	return user, err
}

// generateFromPassword generate the hash from the password string with salt and iterations values.
// the encrypting algorithm is Argon2id.
func generateFromPassword(password string, argon2Params *Argon2Params) (encodedHash string, err error) {
	salt, err := generateRandomBytes(argon2Params.SaltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, argon2Params.Iterations, argon2Params.Memory, argon2Params.Parallelism, argon2Params.KeyLength)

	// Base64 encode the salt and hashed password.
	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, argon2Params.Memory, argon2Params.Iterations, argon2Params.Parallelism, base64Salt, base64Hash)

	return encodedHash, nil
}

// generateRandomBytes returns a random bytes.
func generateRandomBytes(saltLength uint32) ([]byte, error) {
	buf := make([]byte, saltLength)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
