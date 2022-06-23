# MySQL file

```sql
CREATE DATABASE
IF
	NOT EXISTS service DEFAULT CHARACTER 
	SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';

INSERT INTO `tag` VALUES (1, 'JavaScript', 1636437749, 'Forest', 1636437749, '', 0, 0, 1);
INSERT INTO `tag` VALUES (2, 'Go', 1639526364, 'Forest', 1639526364, '', 0, 0, 1);
INSERT INTO `tag` VALUES (3, 'Node', 1639526364, 'Forest', 1639526364, '', 0, 0, 1);
INSERT INTO `tag` VALUES (4, 'React', 1639526672, 'Forest', 1639526672, '', 0, 0, 1);
INSERT INTO `tag` VALUES (5, 'Next', 1639526713, 'Forest', 1639526713, '', 0, 0, 1);
INSERT INTO `tag` VALUES (6, 'Vue', 1639526738, 'Forest', 1639526738, '', 0, 0, 1);
INSERT INTO `tag` VALUES (7, 'Nuxt', 1639526745, 'Forest', 1639526738, '', 0, 0, 1);
INSERT INTO `tag` VALUES (8, 'TypeScript', 1639526755, 'Forest', 1639526738, '', 0, 0, 1);

DROP TABLE IF EXISTS `auth`;
CREATE TABLE `auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(64) DEFAULT '' COMMENT 'Secret',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='认证管理';

INSERT INTO `auth` VALUES (NULL, 1, 'Forest', 'ybMOwVkh0CsR9KLBD5IU1vET2YFfW4XuNrH8x6aAQeipnlotGmZJzqcSdjgP73', 0, 'Forest', 0, '', 0, 0);

DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT '文章ID',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签ID',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';

DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '文章简述',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  `content` longtext COMMENT '文章内容',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

INSERT INTO `article` VALUES (NULL, 'Go Web', '编程不是“纸上谈兵”。编程语言的学习离不开动手实践，因此学习任何一门编程语言的第一步都是要拥有一个这门编程语言的开发环境，这样我们才可以动手编码，理论与实践结合，不仅加速学习效率，还能取得更好的学习效果；接下来学习下如何安装和配置Go语言开发环境。', 'https://img.php.cn/upload/article/000/000/033/5e02dca2731ac901.jpg', '## Go 语言是怎样诞生的？\r\nGo 语言的创始人有三位，分别是图灵奖获得者、C 语法联合发明人、Unix 之父肯·汤普森（Ken Thompson），Plan 9 操作系统领导者、UTF-8 编码的最初设计者罗伯·派克（Rob Pike），以及 Java 的 HotSpot 虚拟机和 Chrome 浏览器的 JavaScript V8 引擎的设计者之一罗伯特·格瑞史莫（Robert Griesemer）。\r\n\r\n![](https://files.mdnice.com/user/8213/e6175ffe-d536-46fd-9716-5f25fd7361f3.png)\r\n<p style=\"margin-top: -40px; color: #666;font-size: 14px\">（从左到右分别是Robert Griesemer、Rob Pike和Ken Thompson）</p>\r\n\r\n他们三个人在 2007 年 9 月 20 日下午的一次普通讨论，就这么成为了计算机编程语言领域的一次著名历史事件，开启了一个新编程语言的历史。\r\n\r\n那天下午，在谷歌山景城总部的那间办公室里，罗伯·派克启动了一个 C++ 工程的编译构建。按照以往的经验判断，这次构建大约需要一个小时。利用这段时间，罗伯·派克和罗伯特·格瑞史莫、肯·汤普森坐在一处，交换了关于设计一门新编程语言的想法。\r\n\r\n之所以有这种想法，是因为当时的谷歌内部主要使用 C++ 语言构建各种系统，但 C++ 的巨大复杂性、编译构建速度慢以及在编写服务端程序时对并发支持的不足，让三位大佬觉得十分不便，他们就想着设计一门新的语言。在他们的初步构想中，这门新语言应该是能够给程序员带来快乐、匹配未来硬件发展趋势并适合用来开发谷歌内部大规模网络服务程序的。\r\n\r\n趁热打铁！在第一天的简短讨论后，第二天这三位大佬又在谷歌总部的“雅温得（Yaounde）”会议室里具体讨论了这门新语言的设计。会后罗伯特·格瑞史莫发出了一封题为“prog lang discussion”的电邮，对这门新编程语言的功能特性做了初步的归纳总结：\r\n\r\n![](https://files.mdnice.com/user/8213/956bf555-dc06-4bbf-8721-e2b9a7a64109.png)\r\n\r\nGo语言第一版特性设计稿这封电邮对这门新编程语言的功能特性做了归纳总结。主要思路是，在C语言的基础上，修正一些明显的缺陷，删除一些被诟病较多的特性，增加一些缺失的功能，比如，使用 import 替代 include、去掉宏、增加垃圾回收、支持接口等。这封电邮成为了这门新语言的第一版特性设计稿，三位大佬在这门语言的一些基础语法特性上达成了初步一致。\r\n\r\n9 月 25 日，罗伯·派克在一封回复电邮中把这门新编程语言命名为“go”：新编程语言被命名为“**go**”。\r\n\r\n![](https://files.mdnice.com/user/8213/b6db66cb-53cd-49be-97f5-9765b4c9b583.png)\r\n\r\n\r\n在罗伯·派克的心目中，“go”这个单词短小、容易输入并且在组合其他字母后便可以用来命名 Go 相关的工具，比如编译器（goc）、汇编器（goa）、链接器（gol）等（go 的早期版本曾如此命名 go 工具链，但后续版本撤销了这种命名方式，仅保留 go 这一统一的工具链名称 ）。\r\n\r\n很多 Go 语言初学者经常称这门语言为 Golang，其实这是不对的：“Golang”仅应用于命名 Go 语言官方网站，而且当时没有用 go.com 纯粹是这个域名被占用了而已。\r\n\r\n## 从“三人行”到“众人拾柴”\r\n\r\n经过早期讨论，Go 语言的三位作者在语言设计上达成初步一致后，便开启了 Go 语言迭代设计和实现的过程。\r\n\r\n2008 年初，Unix 之父肯·汤普森实现了第一版 Go 编译器，用于验证之前的设计。这个编译器先将 Go 代码转换为 C 代码，再由 C 编译器编译成二进制文件。\r\n\r\n到 2008 年年中，Go 的第一版设计就基本结束了。这时，同样在谷歌工作的伊恩·泰勒（Ian Lance Taylor）为 Go 语言实现了一个 gcc 的前端，这也是 Go 语言的第二个编译器。\r\n\r\n伊恩·泰勒的这一成果不仅仅是一种鼓励，也证明了 Go 这一新语言的可行性 。有了语言的第二个实现，对 Go 的语言规范和标准库的建立也是很重要的。随后，伊恩·泰勒以团队的第四位成员的身份正式加入 Go 语言开发团队，后面也成为了 Go 语言，以及其工具设计和实现的核心人物之一。\r\n\r\n罗斯·考克斯（Russ Cox）是 Go 核心开发团队的第五位成员，也是在 2008 年加入的。进入团队后，罗斯·考克斯利用函数类型是“一等公民”，而且它也可以拥有自己的方法这个特性巧妙设计出了 http 包的HandlerFunc类型。这样，我们通过显式转型就可以让一个普通函数成为满足http.Handler接口的类型了。\r\n\r\n不仅如此，罗斯·考克斯还在当时设计的基础上提出了一些更泛化的想法，比如io.Reader和io.Writer接口，这就奠定了 Go 语言的 I/O 结构模型。后来，罗斯·考克斯成为 Go 核心技术团队的负责人，推动 Go 语言的持续演化。\r\n\r\n到这里，Go 语言最初的核心团队形成，Go 语言迈上了稳定演化的道路。\r\n\r\n2009 年 10 月 30 日，罗伯·派克在 Google Techtalk 上做了一次有关 Go 语言的演讲“The Go Programming Language”，这也是 Go 语言第一次公之于众。十天后，也就是 2009 年 11 月 10 日，谷歌官方宣布 Go 语言项目开源，之后这一天也被 Go 官方确定为 Go 语言的诞生日。\r\n\r\n![](https://files.mdnice.com/user/8213/fb18e948-33a4-4867-b378-0285dfd1b9e3.png)\r\n\r\n在 Go 语言项目开源后，Go 语言也迎来了自己的“吉祥物”，是一只由罗伯·派克夫人芮妮·弗伦奇（Renee French）设计的地鼠，从此地鼠（gopher）也就成为了世界各地 Go 程序员的象征，Go 程序员也被昵称为 Gopher，在后面的课程中，我会直接使用 Gopher 指代 Go 语言开发者。\r\n\r\n![](https://files.mdnice.com/user/8213/25b54f72-27e6-45b8-9821-59ea79f74930.png)\r\n\r\nGo 语言项目的开源使得 Go 语言吸引了全世界开发者的目光，再加上 Go 三位作者在业界的影响力以及谷歌这座大树的加持，更多有才华的程序员加入到 Go 核心开发团队中，更多贡献者开始为 Go 语言项目添砖加瓦。于是，Go 在宣布开源的当年，也就是 2009 年，就成为了著名编程语言排行榜 TIOBE 的年度最佳编程语言。\r\n\r\n2012 年 3 月 28 日，Go 1.0 版本正式发布，同时 Go 官方发布了“Go 1 兼容性”承诺：只要符合 Go 1 语言规范的源代码，Go 编译器将保证向后兼容（backwards compatible），也就是说我们使用新版编译器也可以正确编译用老版本语法编写的代码。\r\n\r\n![](https://files.mdnice.com/user/8213/a915f726-7991-42b6-9ba4-18b1efb7e233.png)\r\n\r\n从此，Go 语言发展得非常迅猛。从正式开源到现在，十一年的时间过去了，Go 语言发布了多个大版本更新，逐渐成熟。\r\n\r\n![](https://files.mdnice.com/user/8213/58f7282f-7aae-47c9-9414-e67b5e6f0b23.png)\r\n', 0, 'Forest', 0, '', 0, 0, 1);

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL  COMMENT '用户id',
  `article_id` int(11) NOT NULL COMMENT '文章ID',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  `content` longtext COMMENT '评论内容',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `avatar` varchar(255) DEFAULT '' COMMENT '头像',
  `blog` varchar(255) DEFAULT '' COMMENT '博客地址',
  `email` varchar(255) DEFAULT '' COMMENT '邮箱',
  `html_url` varchar(255) DEFAULT '' COMMENT 'GitHub主页',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `phone` varchar(20) DEFAULT '' COMMENT '电话',
  `created_on` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  INDEX `idx_user_deleted_on`(`deleted_on`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

INSERT INTO `user` VALUES 
  (NULL, 'https://avatars.githubusercontent.com/u/56526369?v=4','https://www.clin.pro','7674254lin@gmail.com','https://github.com/Forest-211','Forest','',1639540296,NULL,NULL),
  (NULL, 'https://cdn.nlark.com/yuque/0/2019/jpeg/anonymous/1571542066833-36c2d16c-4fbe-4b8a-a73f-4e02ba2bb6cf.jpeg','https://www.changlin93.com','767425412@qq.com','https://github.com/Forest-211','lin','',1639544456,NULL,NULL)
```