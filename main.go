package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
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
					"notes":    {"id": "3", "title_zh": "杂记", "title_en": "notes", "url": "/notes", "position": "3"},
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

	r.GET("/api/v1/article/detail/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if id <= 3 {
			var content, path string

			switch id {
			case 1:
				path = "./docs/index.md"
			case 2:
				path = "./docs/base.md"
			case 3:
				path = "./docs/nuxt.md"
			case 4:
				path = "./docs/blog-sql.md"
			default:
				path = "./docs/dart-entry.md"
			}
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal("open file failed:", err)
			}
			content = string(bytes)

			data.Content = content
			c.JSON(200, gin.H{
				"code":    http.StatusOK,
				"message": "success",
				"data":    data,
			})
		} else {
			c.JSON(200, gin.H{
				"code":    404,
				"message": "success",
				"data":    "not found",
			})
		}
	})

	// 杂记列表
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
	r.GET("/api/v1/notes/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		list := notes[0]
		if id <= len(notes) {
			list = notes[id]
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data": gin.H{
				"id":   id,
				"list": list,
			},
		})
	})

	// 评论
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

	// 资源列表
	type ResourceItem struct {
		Name        string `json:"name"`
		Cover       string `json:"cover"`
		Link        string `json:"link"`
		Description string `json:"description"`
	}
	type Resource struct {
		Title string         `json:"title"`
		List  []ResourceItem `json:"list"`
	}
	resource := []Resource{
		{
			Title: "前端框架",
			List: []ResourceItem{
				{
					Name:        "React",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201020/react.svg",
					Link:        "https://zh-hans.reactjs.org",
					Description: "React是一个用于构建用户界面的JavaScript库",
				},
				{
					Name:        "Vue2",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201020/vue.svg",
					Link:        "https://cn.vuejs.org",
					Description: "是一套用于构建用户界面的渐进式框架。与其它大型框架不同的是，Vue 被设计为可以自底向上逐层应用。Vue 的核心库只关注视图层，不仅易于上手，还便于与第三方库或既有项目整合。另一方面，当与现代化的工具链以及各种支持类库结合使用时，Vue 也完全能够为复杂的单页应用提供驱动",
				},
				{
					Name:        "Vue3",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201020/vue.svg",
					Link:        "https://v3.cn.vuejs.org",
					Description: "是一套用于构建用户界面的渐进式框架。与其它大型框架不同的是，Vue 被设计为可以自底向上逐层应用。Vue 的核心库只关注视图层，不仅易于上手，还便于与第三方库或既有项目整合。另一方面，当与现代化的工具链以及各种支持类库结合使用时，Vue 也完全能够为复杂的单页应用提供驱动",
				},
				{
					Name:        "Angular",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201020/angular.svg",
					Link:        "https://angular.cn",
					Description: "学会用 Angular 构建应用，然后把这些代码和能力复用在多种多种不同平台的应用上 —— Web、移动 Web、移动应用、原生应用和桌面原生应用",
				},
				{
					Name:        "Next",
					Cover:       "http://weixin.changlin93.com/nextjs-logo.png",
					Link:        "https://nextjs.org/",
					Description: "Next.js 为您提供生产所需的所有功能的最佳开发人员体验：混合静态和服务器渲染、TypeScript 支持、智能捆绑、路由预取等。 无需配置",
				},
				{
					Name:        "Nuxt",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201023/nuxt-icon.png",
					Link:        "https://nuxtjs.org/",
					Description: "使用 NuxtJS 自信地构建您的下一个 Vue.js 应用程序。 一个开源框架，使 Web 开发变得简单而强大",
				},
			},
		},
		{
			Title: "Vue",
			List: []ResourceItem{
				{
					Name:        "Naive",
					Cover:       "http://weixin.changlin93.com/naive.svg",
					Link:        "https://www.naiveui.com/zh-CN/os-theme",
					Description: "Naive UI是一个Vue3的组件库; 全量使用TypeScript编写，和你的TypeScript项目无缝衔接",
				},
				{
					Name:        "Ant Design",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201021/AntDesign.svg",
					Link:        "https://antdv.com/docs/vue/introduce-cn",
					Description: "企业级中后台产品的交互语言和视觉风格、开箱即用的高质量 Vue 组件",
				},
				{
					Name:        "Vuetify",
					Cover:       "http://weixin.changlin93.com/vuetify-logo.svg",
					Link:        "https://vuetifyjs.com/zh-Hans",
					Description: "Vuetify 是一个纯手工精心打造的 Material 样式的 Vue UI 组件库。 不需要任何设计技能 — 创建叹为观止的应用程序所需的一切都触手可及。",
				},
				{
					Name:        "element-plus",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201021/element.ba99f6df.svg",
					Link:        "https://element-plus.gitee.io/#/zh-CN",
					Description: "网站快速成型工具，一套为开发者、设计师和产品经理准备的基于 Vue 3 的桌面端组件库",
				},
				{
					Name:        "Buefy",
					Cover:       "http://weixin.changlin93.com/buefy-logo.png",
					Link:        "https://buefy.org",
					Description: "Buefy是一个基于Vue.js和Bulma的用户界面组件库",
				},
				{
					Name:        "vben",
					Cover:       "http://weixin.changlin93.com/vben.png",
					Link:        "https://anncwb.github.io/vue-vben-admin-doc/guide/introduction.html",
					Description: "Vue-Vben-Admin 是一个基于 Vue3.0、Vite、 Ant-Design-Vue、TypeScript 的后台解决方案，目标是为开发中大型项目提供开箱即用的解决方案",
				},
				{
					Name:        "vant",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201021/Zent.png",
					Link:        "https://youzan.github.io/vant/#/zh-CN/",
					Description: "Vant 是有赞前端团队开源的移动端组件库，于 2017 年开源，已持续维护 4 年时间。Vant 对内承载了有赞所有核心业务，对外服务十多万开发者，是业界主流的移动端组件库之一",
				},
				{
					Name:        "nut",
					Cover:       "http://weixin.changlin93.com/nut-ui.png",
					Link:        "https://nutui.jd.com/#/",
					Description: "NutUI 是京东风格的 Vue 移动端组件库，开发和服务于移动Web界面的企业级产品;70+ 高质量组件（3.0 持续开发中）",
				},
				{
					Name:        "muse",
					Cover:       "http://weixin.changlin93.com/muse.png",
					Link:        "https://muse-ui.org/#/zh-CN",
					Description: "基于 Vue 2.0 优雅的 Material Design UI 组件库",
				},
			},
		},
		{
			Title: "React",
			List: []ResourceItem{
				{
					Name:        "React-Router",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201021/react-router.svg",
					Link:        "https://react-router.docschina.org",
					Description: "react-router组件是 React 的核心功能，其拥有非常强大的声明式编程模型",
				},
				{
					Name:        "Umi",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201021/UmiJS.png",
					Link:        "https://umijs.org/zh-CN",
					Description: "插件化的企业级前端应用框架",
				},
				{
					Name:        "Dva",
					Cover:       "",
					Link:        "https://dvajs.com",
					Description: "dva首先是一个基于redux和redux-saga的数据流方案，然后为了简化开发体验，dva还额外内置了react-router和 fetch，所以也可以理解为一个轻量级的应用框架",
				},
				{
					Name:        "Ant Design",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201021/AntDesign.svg",
					Link:        "https://ant.design/docs/react/introduce-cn",
					Description: "基于 Ant Design 设计体系的 React UI 组件库，主要用于研发企业级中后台产品",
				},
				{
					Name:        "Material",
					Cover:       "http://weixin.changlin93.com/material-ui.svg",
					Link:        "https://material-ui.com/zh",
					Description: "React 组件用于更快速、更简便的 web 开发。你也可以建立你自己的设计系统，或者从 Material Design 开始",
				},
				{
					Name:        "Ant Design Pro",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201021/AntDesign.svg",
					Link:        "https://pro.ant.design/zh-CN",
					Description: "开箱即用的中台前端/设计解决方案",
				},
				{
					Name:        "Zarm Design",
					Cover:       "http://weixin.changlin93.com/zarm.svg",
					Link:        "https://zarm.gitee.io/#/docs/introduce",
					Description: "Zarm 是众安科技基于React、React-Native研发的一款适用于企业级的移动端UI组件库",
				},
			},
		},
		{
			Title: "Mobile",
			List: []ResourceItem{
				{
					Name:        "React Ntive",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201021/react.16ed8348.svg",
					Link:        "https://reactnative.dev",
					Description: "React Native将原生开发的最佳部分与React相结合,致力于成为构建用户界面的JavaScript和React来编写原生应用的框架",
				},
				{
					Name:        "Flutter",
					Cover:       "http://weixin.changlin93.com/flutter.png",
					Link:        "https://flutter.dev",
					Description: "Flutter是一个由谷歌开发的开源移动应用软件开发工具包，用于为Android、iOS、Windows、Mac、Linux、Google Fuchsia开发应用",
				},
				{
					Name:        "Weex",
					Cover:       "http://weixin.changlin93.com/weex.svg",
					Link:        "http://doc.weex.io/zh/",
					Description: "Weex 是一个可以使用现代化的 Web 技术开发高性能原生应用的框架",
				},
				{
					Name:        "Android",
					Cover:       "http://weixin.changlin93.com/android.png",
					Link:        "https://developer.android.com",
					Description: "Android 提供各种尖端工具和功能，可帮助您针对数十亿用户日常使用的手机、平板电脑、电视和汽车来构建应用",
				},
				{
					Name:        "Swift",
					Cover:       "http://weixin.changlin93.com/swift.png",
					Link:        "https://developer.apple.com/cn/swift/",
					Description: "Swift 是一种强大直观的编程语言，适用于 macOS、iOS、watchOS 和 Apple tvOS 等。编写 Swift 代码的过程充满了乐趣和互动。Swift 语法简洁，但表现力强，更包含了开发者喜爱的现代功能。Swift 代码从设计上保证安全，同时还能开发出运行快如闪电的软件。",
				},
			},
		},
		{
			Title: "Node.js",
			List: []ResourceItem{
				{
					Name:        "Express",
					Link:        "https://expressjs.com",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201023/expressjs.png",
					Description: "用于 Node.js 的快速、独立、简约的 Web 框架",
				},
				{
					Name:        "Koa",
					Link:        "https://koajs.com",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201023/koa.f64b4807.svg",
					Description: "Koa 是由 Express 背后的团队设计的新 Web 框架，旨在成为 Web 应用程序和 API 的更小、更具表现力和更健壮的基础。 通过利用异步函数，Koa 允许您放弃回调并大大增加错误处理。 Koa 在其核心中没有捆绑任何中间件，它提供了一套优雅的方法，使编写服务端变得快速和更优的开发体验",
				},
				{
					Name:        "Nest",
					Link:        "https://nestjs.com",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201023/1603446826638_nest.png",
					Description: "用于构建高效、可靠和可扩展的服务器端应用程序的渐进式 Node.js 框架",
				},
				{
					Name:        "typeorm",
					Link:        "https://typeorm.io/#/",
					Cover:       "http://weixin.changlin93.com/typeorm.png",
					Description: "TypeORM 是一个ORM框架，它可以运行在 NodeJS、Browser、Cordova、PhoneGap、Ionic、React Native、Expo 和 Electron 平台上，可以与 TypeScript 和 JavaScript (ES5,ES6,ES7,ES8)一起使用",
				},
				{
					Name:        "Egg",
					Link:        "https://eggjs.org/zh-cn",
					Cover:       "http://weixin.changlin93.com/egg.svg",
					Description: "Egg.js 为企业级框架和应用而生，我们希望由 Egg.js 孕育出更多上层框架，帮助开发团队和开发人员降低开发和维护成本",
				},
				{
					Name:        "Think",
					Link:        "https://thinkjs.org",
					Cover:       "http://weixin.changlin93.com/think-logo.png",
					Description: "ThinkJS 是一款面向未来开发的 Node.js 框架，整合了大量的项目最佳实践，让企业级开发变得更简单、高效",
				},
				{
					Name:        "Puppeteer",
					Link:        "https://pptr.dev",
					Cover:       "http://weixin.changlin93.com/pupeteer.png",
					Description: "Puppeteer 是一个 Node 库，它提供了一个高级 API 来通过 DevTools 协议控制 Chrome 或 Chromium; Puppeteer 默认无头运行，但可以配置为运行完整（非无头）Chrome 或 Chromium",
				},
				{
					Name:        "npm",
					Link:        "https://www.npmjs.com",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201023/download.png",
					Description: "npm 是 JavaScript package 管理工具，在这里可以找到可复用代码，并以强大的全新方式进行聚合",
				},
				{
					Name:        "yarn",
					Link:        "https://yarnpkg.com/",
					Cover:       "http://weixin.changlin93.com/yarn.svg",
					Description: "npm 是 JavaScript package 管理工具，在这里可以找到可复用代码，并以强大的全新方式进行聚合",
				},
			},
		},
		{
			Title: "状态管理",
			List: []ResourceItem{
				{
					Name:        "Redux",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201023/1603451517463_redux.5ae1af16.svg",
					Link:        "https://redux.js.org",
					Description: "JS 应用的状态容器，提供可预测的状态管理;集中管理应用的状态和逻辑可以让你开发出强大的功能;可与任何 UI 层框架搭配使用，并且有 庞大的插件生态 来实现你的需求。",
				},
				{
					Name:        "Mobx",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201023/1603451452369_mobx.1e58fcaa.png",
					Link:        "https://mobx.js.org/README.html",
					Description: "MobX 是一个经过实战测试的库，它通过透明地应用函数式反应式编程 (TFRP) 使状态管理变得简单和可扩展。",
				},
				{
					Name:        "Vuex",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201023/1603451627144_vue-logo.svg",
					Link:        "https://next.vuex.vuejs.org/zh/index.html",
					Description: "Vuex 是一个专为 Vue.js 应用程序开发的状态管理模式 + 库。它采用集中式存储管理应用的所有组件的状态，并以相应的规则保证状态以一种可预测的方式发生变化",
				},
				{
					Name:        "RxJS",
					Cover:       "http://weixin.changlin93.com/rxjs-logo.png",
					Link:        "https://rxjs.dev",
					Description: "RxJS 是一个使用可观察序列编写异步和基于事件的程序的库",
				},
			},
		},
		{
			Title: "小程序",
			List: []ResourceItem{
				{
					Name:        "Taro",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201020/Taro.png",
					Link:        "https://preactjs.com",
					Description: "Taro是一个开放式跨端跨框架解决方案，支持使用 React/Vue/Nerv等框架来开发微信、京东、百度、支付宝、字节跳动、QQ 小程序、H5、RN 等应用。",
				},
				{
					Name:        "uni-app",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201020/uni.png",
					Link:        "https://uniapp.dcloud.io",
					Description: "uni-app 是一个使用Vue.js开发所有前端应用的框架，开发者编写一套代码，可发布到iOS、Android、Web（响应式）、以及各种小程序（微信/支付宝/百度/头条/QQ/快手/钉钉/淘宝）、快应用等多个平台。",
				},
				{
					Name:        "kbone",
					Cover:       "",
					Link:        "https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/extended/kbone",
					Description: "一个致力于微信小程序和Web端同构的解决方案。",
				},
				{
					Name:        "Chameleon",
					Cover:       "http://mingme.oss-cn-beijing.aliyuncs.com/image/poster/20201021/chameleon.png",
					Link:        "http://cml.didi.cn",
					Description: "一个跨多端开发的统一解决方案，它可以像变色龙一样适应不同的环境。",
				},
				{
					Name:        "Lin UI",
					Cover:       "http://weixin.changlin93.com/lin-ui.png",
					Link:        "https://doc.mini.talelin.com/",
					Description: "Lin UI 是基于微信小程序原生语法实现的组件库。遵循简洁，易用的设计规范。",
				},
			},
		},
		{
			Title: "Go",
			List: []ResourceItem{
				{
					Name:        "go",
					Link:        "https://golang.org",
					Cover:       "http://weixin.changlin93.com/go.png",
					Description: "Golang 具备很强的语言表达能力，支持静态类型安全，能够快速编译大型项目；同时也能够让开发人员访问底层操作系统，极力挖掘计算机CPU资源，还提供了强大的网络编程和并发编程支持。",
				},
				{
					Name:        "gin",
					Link:        "https://gin-gonic.com",
					Cover:       "http://weixin.changlin93.com/gin-logo.png",
					Description: "Gin 是一个用 Go (Golang) 编写的HTTP web框架, 它是一个类似于martini但拥有更好性能的API框架, 优于http router，速度提高了近40倍。如果你需要极好的性能，使用Gin吧!",
				},
				{
					Name:        "beego",
					Link:        "https://beego.me",
					Cover:       "http://weixin.changlin93.com/beego.png",
					Description: "beego 是一个快速开发Go应用的HTTP框架，它可以用来快速开发API、Web及后端服务等各种应用，是一个RESTful的框架，主要设计灵感来源于tornado、sinatra和flask这三个框架，但是结合了Go本身的一些特性（interface、struct嵌入等）而设计的一个框架。",
				},
				{
					Name:        "iris",
					Link:        "https://www.iris-go.com",
					Cover:       "http://weixin.changlin93.com/iris-logo.png",
					Description: "iris 是一个通过Go编写的快速的，简单的，但是功能齐全和非常有效率的web框架。",
				},
				{
					Name:        "go-micro",
					Link:        "https://github.com/asim/go-micro",
					Cover:       "",
					Description: "Go Micro是一个分布式系统开发框架。",
				},
				{
					Name:        "echo",
					Link:        "https://echo.labstack.com",
					Cover:       "http://weixin.changlin93.com/echo.png",
					Description: "高性能、可扩展、极简的Go web框架。",
				},
				{
					Name:        "fiber",
					Link:        "https://gofiber.io",
					Cover:       "http://weixin.changlin93.com/fiber.png",
					Description: "高性能、可扩展、极简的Go web框架。",
				},
				{
					Name:        "gorm",
					Link:        "https://gorm.io/index.html",
					Cover:       "http://weixin.changlin93.com/gorm.png",
					Description: "GORM 是 Go 语言的 ORM 包，功能强大，调用方便、支持关联查询、支持多种关系数据库",
				},
				{
					Name:        "xorm",
					Link:        "https://gitea.com/xorm/xorm/src/branch/master/README_CN.md",
					Cover:       "",
					Description: "xorm 是一个简单而强大的Go语言ORM库; 通过它可以使数据库操作非常简便。",
				},
				{
					Name:        "logrus",
					Link:        "https://github.com/sirupsen/logrus",
					Cover:       "",
					Description: "logrus是目前 GitHub 上 star 数量最多的日志包，它的优点是功能强大、性能高效、高度灵活，还提供了自定义插件的功能。",
				},
				{
					Name:        "zap",
					Link:        "https://pkg.go.dev/go.uber.org/zap",
					Cover:       "",
					Description: "zap是uber开源的日志包，以高性能著称，很多公司的日志包都是基于zap改造而来。",
				},
			},
		},
	}
	r.GET("/api/v1/resource", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data":    resource,
		})
	})

	// 分类
	r.GET("/api/v1/catalog", func(c *gin.Context) {
		catalogs := []string{
			"all",
			"html",
			"css",
			"javascript",
			"vue",
			"react",
			"node",
			"小程序",
			"工程化",
			"go",
			"mongoDB",
			"mysql",
			"postgreSQL",
			"other",
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data":    catalogs,
		})
	})

	// 归档
	r.GET("/api/v1/archive", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data":    "Efforts to develop......",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
