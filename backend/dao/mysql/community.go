package mysql

import (
	"database/sql"
	"errors"
	"github.com/webforum/models"
	"go.uber.org/zap"
)

func GetCommunityNameByID(idStr string) (community *models.Community, err error) {
	community = new(models.Community)
	sqlStr := `select community_id ,community_name from community where community_id = ?`
	err = db.Get(community, sqlStr, idStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(ErrorInvalidID)
		}
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		return nil, errors.New(ErrorQueryFailed)
	}
	return

}
