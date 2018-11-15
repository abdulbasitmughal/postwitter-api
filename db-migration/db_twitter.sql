CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(100) NOT NULL COMMENT 'Email address of user and always be unique',
  `password` varchar(45) DEFAULT NULL,
  `token` text COMMENT 'token issued to user',
  `time_tag` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;

CREATE TABLE `post` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL COMMENT 'user id of post creator from user''s table',
  `message` varchar(500) DEFAULT NULL COMMENT 'message allowed 500 letters',
  `time_tag` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'creation time',
  PRIMARY KEY (`id`),
  KEY `post_fk1_idx` (`user_id`),
  CONSTRAINT `post_fk1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;