package jwt_black_service

import (
	"gorm.io/gorm"

	"go-gin-test/global"
	"go-gin-test/model"
)

// @title    JsonInBlacklist
// @description   create jwt blacklist
// @param     jwtList         model.JwtBlacklist
// @return    err             error

func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.DB.Create(&jwtList).Error
	return
}

// @title    IsBlacklist
// @description   check if the Jwt is in the blacklist or not, 判断JWT是否在黑名单内部
// @param     jwt             string
// @param     jwtList         model.JwtBlacklist
// @return    err             error

func IsBlacklist(jwt string, jwtList model.JwtBlacklist) bool {
	err := global.DB.Where("jwt = ?", jwt).First(&jwtList).Error

	if err == gorm.ErrRecordNotFound { // 不在Jwt黑名单
		return false
	}
	return true
}

// @title    GetRedisJWT
// @description   Get user info in redis
// @param     userName        string
// @return    err             error
// @return    redisJWT        string

func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.REDIS.Get(userName).Result()
	return err, redisJWT
}

// @title    SetRedisJWT 有效期7天
// @description   set jwt into the Redis
// @param     jwtList         model.JwtBlacklist
// @param     userName        string
// @return    err             error

func SetRedisJWT(jwtList model.JwtBlacklist, userName string) (err error) {
	err = global.REDIS.Set(userName, jwtList.Jwt, 1000*1000*1000*60*60*24*7).Err()
	return err
}
