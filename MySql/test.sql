Use blogsdb;

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `mobile` varchar(30) DEFAULT NULL,
  `password` varchar(45) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `posts` (
  `post_id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `text` text,
  `image_url` varchar(255) DEFAULT NULL,
  `create_at` datetime DEFAULT NULL,
  `show_post` varchar(1) DEFAULT NULL,
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `comments` (
  `comment_id` int(11) NOT NULL,
  `comment_text` text,
  `parent_comm_id` int(11) DEFAULT NULL,
  `create_at` datetime DEFAULT NULL,
  `post_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `reply_count` int(3) DEFAULT NULL,
  `comment_level` int(3) DEFAULT NULL,
  `show_comment` varchar(2) DEFAULT 'Y',
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE `blogsdb`.`comments` 
CHANGE COLUMN `comment_id` `comment_id` INT(11) NOT NULL AUTO_INCREMENT ;

ALTER TABLE `blogsdb`.`posts` 
CHANGE COLUMN `post_id` `post_id` INT(11) NOT NULL AUTO_INCREMENT ;

ALTER TABLE posts
ADD CONSTRAINT FK_POSTUSERS
FOREIGN KEY (user_id) REFERENCES users(user_id);

ALTER TABLE comments
ADD CONSTRAINT FK_COMMENTSPOST
FOREIGN KEY (post_id) REFERENCES posts(post_id);
 
insert into users values ('101', 'Shankar', 'shanmicheal25@gmail.com', '85993401', 'abcd1234', 'AMK');
 
insert into posts (post_id, user_id, text, image_url, create_at, show_post) values (1001, 101, 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Malesuada fames ac turpis egestas sed. Imperdiet nulla malesuada pellentesque elit. Amet consectetur adipiscing elit ut aliquam purus sit amet. Iaculis at erat pellentesque adipiscing commodo elit at imperdiet. Amet porttitor eget dolor morbi non arcu risus quis varius. Sed blandit libero volutpat sed cras ornare. A arcu cursus vitae congue. Cursus in hac habitasse platea. Sit amet nisl purus in mollis nunc. \n Purus ut faucibus pulvinar elementum integer enim neque volutpat ac. Enim diam vulputate ut pharetra. Quisque id diam vel quam elementum pulvinar. Urna nunc id cursus metus. Sit amet justo donec enim diam vulputate ut pharetra. In iaculis nunc sed augue lacus viverra vitae congue. Et sollicitudin ac orci phasellus egestas tellus rutrum tellus pellentesque. Mauris ultrices eros in cursus. Est placerat in egestas erat. Integer quis auctor elit sed vulputate mi sit. Vestibulum rhoncus est pellentesque elit ullamcorper dignissim cras tincidunt lobortis. Egestas purus viverra accumsan in nisl nisi. Diam quis enim lobortis scelerisque. Orci phasellus egestas tellus rutrum tellus pellentesque eu. Vel pharetra vel turpis nunc eget lorem dolor sed. \n\tPhasellus faucibus scelerisque eleifend donec pretium vulputate.', '', now(), 'Y');
insert into posts (post_id, user_id, text, image_url, create_at, show_post) values (1002, 101, 'Sed adipiscing diam donec adipiscing tristique risus nec feugiat in. Feugiat vivamus at augue eget arcu dictum varius. Malesuada proin libero nunc consequat interdum varius sit amet mattis. Facilisi etiam dignissim diam quis enim lobortis scelerisque fermentum. Neque egestas congue quisque egestas diam in arcu cursus euismod. Mattis enim ut tellus elementum. Nulla aliquet porttitor lacus luctus. Massa enim nec dui nunc mattis. Neque sodales ut etiam sit amet. Leo vel fringilla est ullamcorper eget nulla facilisi etiam. Massa tempor nec feugiat nisl pretium. Congue nisi vitae suscipit tellus mauris a diam. Ipsum nunc aliquet bibendum enim facilisis gravida neque convallis. Dis parturient montes nascetur ridiculus mus mauris. Non sodales neque sodales ut etiam sit amet nisl. Sed blandit libero volutpat sed. Morbi tincidunt ornare massa eget egestas purus viverra.', '', now(), 'Y');

