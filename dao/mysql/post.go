package mysql

import (
	"database/sql"
	"fmt"
	"web_app/models"
)

//CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post (post_id,title,content,author_id,community_id)
               values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

//GetPostById 根据id查询单个帖子数据
func GetPostById(pid int64) (data *models.Post, err error) {
	data = new(models.Post)
	sqlStr := `select 
	post_id,title,content,author_id,community_id,create_time
	from post
	where post_id = ?
	`
	if err := db.Get(data, sqlStr, pid); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return data, err
}
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select 
		post_id,title,content,author_id,community_id,create_time
		from post
		ORDER BY create_time
		DESC 
		limit ?,?  
	`
	fmt.Println(page, size)
	//第一个问号表示从那个开始，第二个问号表示一个页面显示几个
	posts = make([]*models.Post, 0, 3) //不要写成make（[]*modles.post,2） 0:表示长度，2表示容量
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	//fmt.Println(err)
	return
}
