package dal

import (
	"context"
	"database/sql"

	"github.com/LyonNee/app-layout/entity"
	"github.com/LyonNee/app-layout/pkg/log"
	"go.uber.org/zap"
)

func InsertOneUser(ctx context.Context, conn *sql.Conn, user entity.User) (uint64, error) {
	result, err := conn.ExecContext(ctx, "INSERT INTO users (name,age,phone_number,password,create_at,update_at) VALUES (?,?,?,?,?,?);",
		user.Name,
		user.Age,
		user.PhoneNumber,
		user.Password,
		user.CreateAt,
		user.UpdateAt,
	)
	if err != nil {
		log.ZapLogger().Error("数据库插入数据失败", zap.Error(err))
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		log.ZapLogger().Warn("获取insert id失败", zap.Error(err))
		return 0, nil
	}

	return uint64(insertId), nil
}

func QueryUserByPhoneNum(ctx context.Context, conn *sql.Conn, phoneNum string) (*entity.User, error) {
	var userEntity entity.User

	row := conn.QueryRowContext(ctx, "SELECT id,name,age,phone_number,password,create_at,update_at FROM users WHERE phone_number=?;", phoneNum)
	if err := row.Scan(
		&userEntity.ID,
		&userEntity.Name,
		&userEntity.Age,
		&userEntity.PhoneNumber,
		&userEntity.CreateAt,
		&userEntity.UpdateAt,
	); err != nil {
		if err == sql.ErrNoRows {
			log.ZapLogger().Warn("查询了不存在的blockchain数据", zap.String("phone_number", phoneNum))
			return nil, err
		}

		log.ZapLogger().Error("数据复制失败", zap.Error(err))
		return nil, err
	}

	return &userEntity, nil
}
