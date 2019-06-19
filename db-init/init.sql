CREATE TABLE `pets`
(
  `id` int
(11) NOT NULL AUTO_INCREMENT,
  `name` varchar
(128) NOT NULL,
  `photo_urls` text,
  `status` varchar
(64) NOT NULL,
  PRIMARY KEY
(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1