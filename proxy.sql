-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- 主机： localhost:3306
-- 生成日期： 2019-04-15 13:54:12
-- 服务器版本： 5.7.23
-- PHP 版本： 5.5.38

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- 数据库： `test`
--

-- --------------------------------------------------------

--
-- 表的结构 `proxy`
--

CREATE TABLE `proxy` (
  `id` int(10) UNSIGNED NOT NULL,
  `ip` text NOT NULL,
  `port` varchar(8) NOT NULL,
  `addr` text NOT NULL,
  `time` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `proxy`
--

INSERT INTO `proxy` (`id`, `ip`, `port`, `addr`, `time`) VALUES
(478, '54.36.44.253', '1080', '中国 上海 上海 UCloud云计算', '2019-04-14 07:06:19'),
(489, '103.52.135.126', '8080', '中国 吉林 吉林 电信', '2019-04-14 07:09:01'),
(519, '89.27.165.167', '8080', '中国 陕西 延安 电信', '2019-04-14 08:09:52'),
(523, '117.212.91.147', '8080', '中国 广东 深圳 阿里云', '2019-04-14 08:10:10'),
(524, '52.41.29.48', '8080', '中国 上海 上海 互联港湾', '2019-04-14 08:10:12'),
(528, '93.190.137.146', '8080', '中国 江苏 淮安 电信', '2019-04-14 08:10:15'),
(533, '52.83.146.11', '3128', '中国 宁夏 中卫 Amazon/西部云基地', '2019-04-14 08:10:22'),
(538, '117.191.11.78', '8080', '中国 新疆 喀什 移动', '2019-04-14 08:10:29'),
(562, '190.60.69.162', '8080', '中国 辽宁 鞍山 电信', '2019-04-14 08:11:19'),
(566, '114.95.187.239', '8060', '中国 上海 上海 电信', '2019-04-14 08:11:28');

--
-- 转储表的索引
--

--
-- 表的索引 `proxy`
--
ALTER TABLE `proxy`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `proxy`
--
ALTER TABLE `proxy`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1543;
