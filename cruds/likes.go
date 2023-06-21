package cruds

import (
	"errors"
	"oc-2023/db"
	// "github.com/Kyutech-C3/toybox-server/のデータベース"
)

func GiveLike(workId string, userId string) ( /*ToyBoxに返すWorkId,*/ err error) {
	// workIdの型があっているか
	// workはToyBoxから、userIdはフロントのuuidから
	if err = db.Psql.Where("work_id = ? AND user_id = ?", workId, userId).First(&db.Likes{}).Error; err == nil {
		err = errors.New("Like already exists.")
		return
	}

	lk := db.Likes{
		WorkID: workId,
		UserID: userId,
	}
	if err = db.Psql.Create(lk).Error; err != nil {
		return
	}
	// ここでToyBoxから作品を取得する関数 GetPost参考
	return
}

func DeleteLike(workId string, userId string) ( /*ToyBoxに返すWorkId,*/ err error) {
	if err = db.Psql.Where("work_id = ? AND user_id = ?", workId, userId).First(&db.Likes{}).Error; err != nil {
		err = errors.New("Like not found")
		return
	}

	lk := db.Likes{
		WorkID: workId,
		UserID: userId,
	}
	if err = db.Psql.Delete(lk).Error; err != nil {
		return
	}
	// ここで
	return
}
