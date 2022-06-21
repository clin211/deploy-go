package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 导航
	r.GET("/api/v1/nav", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"nav": map[string]map[string]string{
					"article":  {"id": "1", "title_zh": "文章", "title_en": "article", "url": "/article", "position": "1"},
					"resource": {"id": "2", "title_zh": "资源", "title_en": "resource", "url": "/article", "position": "2"},
					"notes":    {"id": "3", "title_zh": "笔记", "title_en": "notes", "url": "/notes", "position": "3"},
					"archive":  {"id": "4", "title_zh": "归档", "title_en": "archive", "url": "/archive", "position": "4"},
					"about":    {"id": "5", "title_zh": "关于", "title_en": "about", "url": "/about", "position": "5"},
				},
			},
		})
	})

	// 文章列表
	type Article struct {
		ID          string   `json:"id"`
		Creation    string   `json:"creation"`
		Cover       string   `json:"cover"`
		Title       string   `json:"title"`
		Content     string   `json:"content"`
		CreateTime  string   `json:"create_time"`
		UpdateTime  string   `json:"update_time"`
		Tags        []string `json:"tags"`
		Author      string   `json:"author"`
		Views       string   `json:"views"`
		Favorite    string   `json:"favorite"`
		Comments    string   `json:"comments"`
		Description string   `json:"description"`
	}
	data := Article{
		ID:          "1",
		Creation:    "原创",
		Cover:       "https://img.yzcdn.cn/vant/cat.jpeg",
		Title:       "模块化、commonJS、AMD、CMD",
		Content:     "文章1内容",
		CreateTime:  "2022/04/06 23:39:05",
		UpdateTime:  "2022/04/07 22:04:47",
		Tags:        []string{"模块化", "commonJS", "AMD", "CMD"},
		Author:      "Forest",
		Views:       "100",
		Favorite:    "0",
		Comments:    "10",
		Description: "随着前端应用的日益复杂化，我们的项目已经逐渐膨胀到了不得不花大量时间去管理的程度。而模块化就是一种最主流的项目组织方式，它通过把复杂的代码按照功能划分为不同的模块单独维护，从而提高开发效率、降低维护成本。 随着前端应用的日益复杂化，我们的项目已经逐渐膨胀到了不得不花大量时间去管理的程度。而模块化就是一种最主流的项目组织方式",
	}
	r.GET("/api/v1/article/list", func(c *gin.Context) {
		list := []Article{}
		for i := 0; i < 10; i++ {
			data.ID = strconv.Itoa(i + 1000)
			list = append(list, data)
		}
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"list": list,
			},
		})
	})

	r.GET("/api/v1/article/detail/:id", func(c *gin.Context) {})

	type Notes struct {
		ID          int    `json:"id"`
		Cover       string `json:"cover"`
		Time        string `json:"time"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	notes := []Notes{
		{
			ID:          1,
			Cover:       "https://weixin.changlin93.com/notes/timthumb1.jpg",
			Time:        "三月 16, 2019",
			Title:       "我们都曾畏惧告别",
			Description: "我实际上是个十分口拙的人。而且，特别是在关键时刻尤为口拙。比如告别。告别似乎是一个普遍公认的隆重时刻，也是一个最能够让人感怀的时刻。越是这样的时刻，我就越是畏惧。倒不是怕伤心怕落泪，而是怕说话。我们都曾畏惧告别。",
		},
		{
			ID:          2,
			Cover:       "https://weixin.changlin93.com/notes/timthumb2.jpg",
			Time:        "三月 21, 2019",
			Title:       "终有弱水替沧海，再把相思寄巫山",
			Description: "那个笑起来眼睛弯弯成月牙，灿若星光，梨涡勾起俏皮弧度，耀如暖阳的姑娘，她恋爱了。那个在你耳边喋喋不休谈起心爱的明星，议论着女生间斗角勾心，嘟囔着嘴问不解问题的姑娘，她终于恋爱了。那个与你谈起未来一片向往，面对困难佯装坚强，却还是柔弱肩膀，纯洁的一尘不染的白纸一张的姑娘，她还是恋爱了。结局即在意料之中，又在想象之外。微风揉过，就着一分怀旧，两分温柔，和十分的不再回头。",
		},
		{
			ID:          3,
			Cover:       "https://weixin.changlin93.com/notes/timthumb3.jpg",
			Time:        "六月 01, 2019",
			Title:       "后来的我们",
			Description: "曾经，我以为自己是个很幽默的人，比如逗女孩开心，讲的是这样的笑话。一只蚂蚁第一次出门觅食，突然下起了雨，它不知道怎么回到蚁窝，这时遇见了另外一只蚂蚁，于是问：“你都如何回蚁窝？",
		},
		{
			ID:          4,
			Cover:       "https://weixin.changlin93.com/notes/timthumb4.jpg",
			Time:        "六月 17, 2019",
			Title:       "年少不懂《还珠3》，看懂已是而立年",
			Description: "那时候我觉得黄奕没有赵薇眼睛大，没有赵薇活泼可爱。马伊俐又没有林心如那样楚楚可怜，弱不禁风。何书桓饰演的五阿哥也没有苏有朋身上的少年气息。最重要的是黄晓明饰演的萧剑，完美的败坏了第二部的萧剑在我心中的好感。",
		},
		{
			ID:          5,
			Cover:       "https://weixin.changlin93.com/notes/timthumb5.jpg",
			Time:        "八月 6, 2019",
			Title:       "要么孤独，要么庸俗",
			Description: "刷朋友圈的时候，看到一条由三感音乐故事拍摄的短视频，被文案戳中了泪点; 22岁生日，一个人吃火锅，还好锅底可以点最辣的；187次路过的码头，4次遇到一对情侣，两个人眼中的风景，也不见得更好看；第6次一个人搬家，扔掉了3箱旧东西，很遗憾，好像连回忆也一并被丢弃了。",
		},
		{
			ID:          6,
			Cover:       "https://weixin.changlin93.com/notes/timthumb6.jpg",
			Time:        "十二月 12, 2019",
			Title:       "余生，我们一起吃饭",
			Description: "我和他总是在学校的食堂里相遇。每次遇见他，我总是拉着自己的饭友跑到一个偏僻的角落里就餐。我不想让他看到我吃饭时的样子。可他还是看到了，而且还看得特别认真。余生，我们一起吃饭。",
		},
		{
			ID:          7,
			Cover:       "https://weixin.changlin93.com/notes/timthumb7.jpg",
			Time:        "一月 01, 2020",
			Title:       "我终于瘦下来了，却不再喜欢你",
			Description: "许桐最大的毛病就是贪吃。她尤其喜欢吃甜食，精致的蛋糕、松软的面包、酥脆的饼干都是她的心头好，“甜食可以让人心情愉快”——她这样解释。所以，城市里有头有脸的西点屋她几乎都办了张会员卡。",
		}, {
			ID:          8,
			Cover:       "https://weixin.changlin93.com/notes/timthumb8.jpg",
			Time:        "一月 05, 2020",
			Title:       "说句再见吧，少年",
			Description: "毕业之后，我和大学的很多同学都留在了这座不大的城市里，两个室友在学校附近租了一套小资豪华房，我在北边签了一份朝九晚五的工作，顺理成章的一个人搬了出来，除了熬夜就会加深的黑眼圈，越发暗沉的肤色，还有慢慢爬上眼睑的细纹，我发现自己在逐渐变老，头发已经脱落了很多，夜里睡觉越来越难眠……心里却只空荡荡地挂着一个年轻的背影。",
		},
		{
			ID:          9,
			Cover:       "https://weixin.changlin93.com/notes/timthumb9.jpg",
			Time:        "一月 23, 2020",
			Title:       "失恋，让我成为更好的人",
			Description: "我是在机场认识的栗子姐，跟她在自助充电那儿偶遇，在等待电充满的时间里，我们随便聊了几句。她说自己今年28岁，未婚，一个人经营一家咖啡厅，趁着最近店里不忙，打算来一次短途旅行。失恋，让我成为更好的人。",
		},
		{
			ID:          10,
			Cover:       "https://weixin.changlin93.com/notes/timthumb10.jpg",
			Time:        "二月 5, 2020",
			Title:       "让你放弃和等待，是为了给你最好的安排",
			Description: "这是普普通通的一天。早上起来，她发现家里停电了。于是没办法用热水洗漱，用电吹风吹头发，不能热牛奶，烤面包，只好草草打理一下就出门。刚走进电梯，邻居家养的小狗一下子冲进来扑住，上周刚买的米白长裙上顿时出现两只黑黑的爪印儿。让你放弃和等待，是为了给你最好的。",
		},
		{
			ID:          11,
			Cover:       "https://weixin.changlin93.com/notes/timthumb2.jpg",
			Time:        "九月 16, 2021",
			Title:       "二十多岁的我们，为什么觉得谈恋爱好难",
			Description: "前两天，我和朋友约在咖啡馆聊天虚度假期。有一个朋友为感情的事愁眉不展，她说在生活的压力下，感觉自己变得难以心动，对一切都感到麻木。我也思考起，自己和身边人的感情生活。确实，好像一切都变了，但不是变得难以心动，而是心动之后所面对的一连串现实问题，让我们望而却步了。",
		}, {
			ID:          12,
			Cover:       "https://weixin.changlin93.com/notes/timthumb2.jpg",
			Time:        "九月 18, 2021",
			Title:       "那些完美主义让我们止步不前",
			Description: "前两天，我和朋友约在咖啡馆聊天虚度假期。有一个朋友为感情的事愁眉不展，她说在生活的压力下，感觉自己变得难以心动，对一切都感到麻木。我也思考起，自己和身边人的感情生活。确实，好像一切都变了，但不是变得难以心动，而是心动之后所面对的一连串现实问题，让我们望而却步了。",
		},
	}
	r.GET("/api/v1/notes", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"list": notes,
			},
		})
	})

	// notes details
	r.GET("/api/v1/notes/:id", func(c *gin.Context) {})

	type Comment struct {
		ID       int    `json:"id"`
		Nickname string `json:"nickname"`
		Time     string `json:"time"`
		Content  string `json:"content"`
		Avatar   string `json:"avatar"`
		Position string `json:"position"`
	}
	comments := []Comment{
		{
			ID:       1,
			Nickname: "Forest",
			Time:     "2021-09-06 20:55:23",
			Content:  "生活要过下去，她们只能选择性遗忘紫薇要忘记尔康曾经同意娶慕沙，忘记他戒毒时可怕的样子尔康要忘记那八个月里不堪的自己小燕子要忘记知画母子，忘记永琪曾经的渣男行为永琪要忘记自己的阿哥身份，忘记自己的牺牲萧剑要忘记灭门之恨晴儿要忘记萧剑的几次离弃",
			Avatar:   "http://weixin.changlin93.com/avatar.jpg",
			Position: "四川省成都市"},
		{
			ID:       2,
			Nickname: "Forest",
			Time:     "2021-09-06 20:55:23",
			Content:  "清醒时做事，糊涂时读书，大怒时睡觉，独处时思考；做一个幸福的人，读书，旅行，努力工作，关心身体和心情，成为最好的自己",
			Avatar:   "http://weixin.changlin93.com/avatar.jpg",
			Position: "四川省成都市",
		},
		{
			ID:       3,
			Nickname: "line",
			Time:     "2021-09-06 20:55:23",
			Content:  "我们要面对的苦难和挫折不会因为我们遭遇够多而减少，上天不会因为你的悲惨而悯怜你，它只会安排一道又一道难关来磨炼你的意志就像唐僧西天取经一样，想要取得真经，就不要有返途的念想，你除了继续死磕到底，根本没有其他退路前面的路还很远，你可能会哭，但是一定要走下去，一定不能停",
			Avatar:   "",
			Position: "北京市",
		},
		{
			ID:       4,
			Nickname: "fangfang",
			Time:     "2021-09-06 20:55:23",
			Content:  "世界很大，风景也很美，机会很多人生很短，不要蜷缩在小块阴影里",
			Avatar:   "https://pic1.zhimg.com/80/v2-10bbed439ad748ec70e3e1a8bf5b7a84_1440w.jpg?source=1940ef5c",
			Position: "上海市",
		},
		{
			ID:       43,
			Nickname: "feifei",
			Time:     "2021-09-06 20:55:23",
			Content:  "再美好的时光，都会浓缩为历史;再遥远的等待，只要坚持总会到来",
			Avatar:   "https://pic1.zhimg.com/80/v2-1eb26ad2e7ce2bd451053e1615c7aee1_1440w.jpg?source=1940ef5c",
			Position: "拉萨",
		},
		{
			ID:       143,
			Nickname: "yang",
			Time:     "2021-09-06 20:55:23",
			Content:  " 在你努力的时候，不管多么的狼狈不堪，都是你所有模样中最美的样子",
			Avatar:   "https://pic2.zhimg.com/80/v2-bd66acc5dc8dc94149722c393e218e8a_1440w.jpg?source=1940ef5c",
			Position: "四川省巴中市",
		},
	}
	r.GET("/api/v1/notes/:id/comments", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"list": comments,
			},
		})
	})

	type ResourceItem struct {
		ID       int    `json:"id"`
		Nickname string `json:"nickname"`
		Time     string `json:"time"`
		Content  string `json:"content"`
		Avatar   string `json:"avatar"`
		Position string `json:"position"`
	}
	type Resource struct {
		Title string         `json:"title"`
		List  []ResourceItem `json:"list"`
	}

	r.GET("/api/v1/resource", func(c *gin.Context) {

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
