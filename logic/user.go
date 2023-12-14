package logic

import (
	"context"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/lyonnee/go-gin-template/infra/log"
	"github.com/lyonnee/go-gin-template/repository"
	"github.com/lyonnee/go-gin-template/util"
	"go.uber.org/zap"
	"golang.org/x/crypto/argon2"
)

// Agron2 hash params
type agron2idHashParams struct {
	salt_len    int
	iterations  uint32
	memory      uint32
	parallelism uint8
	keyLen      uint32
}

var ARGON2ID_HASH_PARAMS = agron2idHashParams{
	salt_len:    16,
	iterations:  2,
	memory:      2 * 1024,
	parallelism: 2,
	keyLen:      32,
}

type UserLogic struct {
	repo *repository.UserRepo
}

func NewUserLogic() *UserLogic {
	return &UserLogic{
		repo: repository.NewUserRepo(),
	}
}

func (logic *UserLogic) Register(ctx context.Context, name string, age uint8, phoneNum string, password string) (uint64, error) {
	hashStr, err := hashPassword(password)
	if err != nil {
		return 0, err
	}
	userEntity, err := logic.repo.InsertOne(ctx, name, age, phoneNum, hashStr)
	if err != nil {
		log.Error("insert user to db failed", zap.Error(err))
		return 0, err
	}

	return userEntity.ID, nil
}

func (logic *UserLogic) Login(ctx context.Context, phoneNum string, password string) (uint64, string, string, error) {
	userEntity, err := logic.repo.GetByPhoneNumber(ctx, phoneNum)
	if err != nil {
		log.Error("query user failed", zap.Error(err))
		return 0, "", "", err
	}

	ok, err := comparePasswordAndHash(password, *userEntity.AuthHash)
	if err != nil {
		return 0, "", "", ERR_LOGIC_EXECUTE_FAIL
	}

	if !ok {
		return 0, "", "", ERR_INVALID_PASSWORD
	}

	return userEntity.ID, *userEntity.Name, *userEntity.PhoneNumber, nil
}

// ================================ private method =================================== //
func hashPassword(password string) (string, error) {
	// 生成随机盐
	salt, err := util.GenerateRandomBytes(ARGON2ID_HASH_PARAMS.salt_len)
	if err != nil {
		return "", err
	}

	// 生成hash
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		ARGON2ID_HASH_PARAMS.iterations,
		ARGON2ID_HASH_PARAMS.memory,
		ARGON2ID_HASH_PARAMS.parallelism,
		ARGON2ID_HASH_PARAMS.keyLen,
	)

	// base64序列化
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// 拼接序列化后的字符串和hash参数
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$t=%d,m=%d,p=%d$%s$%s",
		argon2.Version,
		ARGON2ID_HASH_PARAMS.iterations,
		ARGON2ID_HASH_PARAMS.memory,
		ARGON2ID_HASH_PARAMS.parallelism,
		b64Salt,
		b64Hash)

	return encodedHash, nil
}

func comparePasswordAndHash(password, encodedHash string) (bool, error) {
	// 从hash字符串中提取hash参数
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	// 使用原hash参数和输入的passwd生成新的hash
	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLen)

	// 校验hash
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (*agron2idHashParams, []byte, []byte, error) {
	// 拆分hash字符串
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ERR_INVALID_HASH
	}

	var version int
	// 验证argon2算法的版本
	if _, err := fmt.Sscanf(vals[2], "v=%d", &version); err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ERR_INVALID_HASH
	}

	// 提取hash参数
	p := &agron2idHashParams{}
	if _, err := fmt.Sscanf(vals[3], "t=%d,m=%d,p=%d", &p.iterations, &p.memory, &p.parallelism); err != nil {
		return nil, nil, nil, err
	}

	// 提取salt和salt_len
	salt, err := base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.salt_len = len(salt)

	// 提取hash
	hash, err := base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLen = uint32(len(hash))

	return p, salt, hash, nil
}
