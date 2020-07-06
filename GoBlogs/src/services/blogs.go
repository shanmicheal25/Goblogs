package services

import (
	u "GoBlogs/apiHelpers"
	"GoBlogs/models"
	r "GoBlogs/resources"
	"fmt"
	"time"

	"github.com/biezhi/gorm-paginator/pagination"
)

//BlogsService struct
type BlogsService struct {
	Posts []models.Posts
}

//BlogService struct
type BlogService struct {
	Posts models.Posts
}

type BlogsCommService struct {
	Posts []r.Posts
}

//BlogsList function returns the list
func (blogs BlogsService) BlogsFilter(page int, limit int) map[string]interface{} {

	postArr := blogs.Posts
	fmt.Println(postArr)
	//comments := models.Comments{}  Related(&comments)

	//paginator := models.GetDB().Preload("Comments").Find(&postArr)
	paginator := pagination.Paging(&pagination.Param{
		DB:      models.GetDB().Preload("Comments").Where("show_post = 'Y'"),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"post_id asc"},
		ShowSQL: true,
	}, &postArr)

	if paginator == nil {
		response := u.Message(1, "Invalid request")
		return response
	}

	fmt.Println("________________________________________")
	fmt.Println(paginator.Records)
	// myEventChan := paginator.Records.(string)
	//json.Unmarshal([]byte(paginator.Records), results)
	//fmt.Println(myEventChan)

	// for _, r := range rec {
	// 	buildTree(r.comments, 0, 0)
	// }

	fmt.Println("________________________________________")

	response := u.Message(0, "Post Blogs Result ")
	response["BlogsData"] = paginator
	return response
}

//BlogsCommFilter function returns the list
func (blogs BlogsCommService) BlogsCommFilter(page int, limit int) map[string]interface{} {

	postArr := blogs.Posts
	paginator := pagination.Paging(&pagination.Param{
		DB: models.GetDB().Select([]string{
			"posts.post_id",
			"posts.text",
			"posts.user_id",
			"posts.image_url",
			"posts.comments"}).Joins("LEFT OUTER JOIN (SELECT * FROM comments t3 where t3.comment_level=1) comments ON comments.post_id = posts.post_id").Find(&postArr).Related("comments"),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"post_id asc"},
		ShowSQL: true,
	}, &postArr)

	if paginator == nil {
		response := u.Message(1, "Invalid request")
		return response
	}

	response := u.Message(0, "Post Blogs Result ")
	response["BlogsData"] = paginator
	return response
}

// Blog Insert record .
func (blog BlogService) BlogsPost() map[string]interface{} {

	blogs := blog.Posts

	blogs.CreateAt = time.Now()
	blogs.ShowPost = "Y"

	err := models.GetDB().Create(&blogs).Error
	if err != nil {
		fmt.Println("Error ", err)
		response := u.Message(1, "Invalid request")
		return response
	}

	response := u.Message(200, "Blogs Posted successfully")
	response["BlogsData"] = blogs
	return response

}

func buildTree(tbl []models.Comments, parent uint, depth int) {
	for _, r := range tbl {
		if r.ParentCommID == parent {
			for i := 0; i < depth; i++ {
				fmt.Print("--")
			}
			fmt.Print(r.CommentText, "\n\n")
			buildTree(tbl, r.CommentID, depth+1)
		}
	}
}
