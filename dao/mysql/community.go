package mysql

import (
	"database/sql"
	"web_app/models"

	"go.uber.org/zap"
)

// 查找 GetCommunityList 全部数据

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in database")
			err = nil
		}
	}
	return
}

//GetCommunityDetailById 根据ID查询社区详情
func GetCommunityDetailById(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `Select
			community_id, community_name, introduction, create_time
			from community
			where community_id = ?`
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}
