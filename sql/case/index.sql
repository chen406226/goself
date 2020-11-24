CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `author_id` int(11) NOT NULL,
  `category_id` int(255) NOT NULL,
  `views` int(255) NOT NULL,
  `comments` int(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `content` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*catetoryid为1 commentid》1 wiews最多的articleid
EXPLAIN SELECT id,author_id FROM article where article.comments =1 AND article.category_id = 1 ORDER BY views desc LIMIT 1;
范围以后的索引会失效>1
$ EXPLAIN SELECT id,author_id FROM article where article.comments > 1 AND article.category_id = 1 ORDER BY views desc LIMIT 1;
$ show index from artic
$ // alter table 'article' add index idx_art_ccv('category_id','comments','views');
$ create index idx_art_ccv on article(category_id,comments,views);
$ drop index idx_art_ccv on article
2次尝试 好
$ create index idx_art_cv on article(category_id,views);
*/
CREATE TABLE if not exists `class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `card` int(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

CREATE TABLE `book` (
  `bookid` int(11) NOT NULL AUTO_INCREMENT,
  `card` int(255) NOT NULL,
  PRIMARY KEY (`bookid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into class(card) value(floor(1+(rang()*20)))

select * from book inner join class on book.card = class.card
/*
1
$ EXPLAIN select * from class left join book on book.card = class.card
好 create index card_book on book(card)
drop indtx card_book on book
create index card_book on class(card)
*/
/*
三表
加上上面的两个表
*/
CREATE TABLE `phone` (
  `phoneid` int(11) NOT NULL AUTO_INCREMENT,
  `card` int(255) NOT NULL,
  PRIMARY KEY (`phoneid`),
  UNIQUE KEY `phoneid` (`phoneid`,`card`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*
using join buffer
SELECT * from class inner JOIN book ON class.card=book.card inner JOIN phone on book.card = phone.card
如果用leftjoin   row全选all
ALTER TABLE `book` add INDEX Y (`card`)
ALTER TABLE `phone` add INDEX Y (`card`)
三表都是右 边 两个  加index
*/

/*索引失效
*/
CREATE TABLE `staffs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `NAME` varchar(255) NOT NULL,
  `age` int(11) NOT NULL DEFAULT '0',
  `pos` varchar(255) NOT NULL,
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '入职时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

alter table staffs add index idx_staffsnameAgePos(name,age,pos)




