package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.String(http.StatusOK, "PONG")
}

type Mine struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
	Status   string `json:"status"`
	Sign     string `json:"sign"`
	Avatar   string `json:"avatar"`
}

type FriendList struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
	Status   string `json:"status"`
	Sign     string `json:"sign"`
	Avatar   string `json:"avatar"`
}

type Friend struct {
	Groupname string       `json:"groupname"`
	Id        int          `json:"id"`
	List      []FriendList `json:"list"`
}

type UserListResult struct {
	Mine   Mine     `json:"mine"`
	Friend []Friend `json:"friend"`
}

func GetUserList(c *gin.Context) {

	//err, user := service.FindUserById(1)
	//if err != nil {
	//	fmt.Printf("err %v", err)
	//}
	//fmt.Println(user)
	//friends := UserListResult{
	//	Mine: Mine{
	//		Username: "纸飞机",
	//		Id:       100000,
	//		Status:   "online",
	//		Sign:     "在深邃的编码世界，做一枚轻盈的纸飞机",
	//		Avatar:   "http://res.layui.com/images/fly/avatar/00.jpg",
	//	},
	//	Friend: []Friend{
	//		{
	//			Groupname: "知名人物",
	//			Id:        0,
	//			List: []FriendList{
	//				{
	//					Username: "贤心",
	//					Id:       100001,
	//					Status:   "online",
	//					Sign:     "这些都是测试数据，实际使用请严格按照该格式返回",
	//					Avatar:   "http://tva1.sinaimg.cn/crop.0.0.118.118.180/5db11ff4gw1e77d3nqrv8j203b03cweg.jpg",
	//				},
	//				{
	//					Username: "贤心",
	//					Id:       100002,
	//					Status:   "online",
	//					Sign:     "这些都是测试数据，实际使用请严格按照该格式返回",
	//					Avatar:   "http://tva1.sinaimg.cn/crop.0.0.118.118.180/5db11ff4gw1e77d3nqrv8j203b03cweg.jpg",
	//				},
	//			},
	//		},
	//	},
	//}
	//c.Header("Access-Control-Allow-Origin", "*")
	//response.OkWithData(friends, c)
}
