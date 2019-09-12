mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 5.7.27, for Linux (x86_64)
--
-- Host: localhost    Database: whacos
-- ------------------------------------------------------
-- Server version	5.7.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `blog_article`
--

DROP TABLE IF EXISTS `blog_article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text COMMENT '内容',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `blog_auth`
--

DROP TABLE IF EXISTS `blog_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `blog_tag`
--

DROP TABLE IF EXISTS `blog_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_dict`
--

DROP TABLE IF EXISTS `t_dict`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_dict` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `parent_id` bigint(20) DEFAULT NULL COMMENT '父字典ID',
  `key` varchar(16) DEFAULT NULL COMMENT '字典名',
  `value` varchar(32) DEFAULT NULL COMMENT '字典值',
  `dict_type` varchar(16) DEFAULT NULL COMMENT '字典类型',
  `sort_num` int(4) DEFAULT NULL COMMENT '排序号',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建人',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '修改人',
  `updated_time` datetime DEFAULT NULL COMMENT '修改时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `del_flag` int(2) DEFAULT NULL COMMENT '删除标记',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`) USING BTREE COMMENT 'parent_id 普通索引',
  KEY `idx_dict_type` (`dict_type`) USING BTREE COMMENT 'dict_type 普通索引',
  KEY `idx_updated_by` (`updated_by`) USING BTREE COMMENT 'updated_by 普通索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='字典表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_file`
--

DROP TABLE IF EXISTS `t_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_file` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `file_name` varchar(64) DEFAULT NULL COMMENT '文件名',
  `key_id` varchar(32) DEFAULT NULL COMMENT 'keyId',
  `suffix` varchar(8) DEFAULT NULL COMMENT '文件后缀',
  `size` int(16) DEFAULT NULL COMMENT '文件大小',
  `address` varchar(64) DEFAULT NULL COMMENT '文件地址',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建人',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '修改人',
  `updated_time` datetime DEFAULT NULL COMMENT '修改时间',
  `remark` varchar(128) DEFAULT NULL COMMENT '备注',
  `del_flag` int(2) DEFAULT NULL COMMENT '删除标记',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_key_id` (`key_id`) USING BTREE COMMENT 'key_id 唯一索引',
  KEY `idx_created_by` (`created_by`) USING BTREE COMMENT 'created_by 普通索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文件表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_log`
--

DROP TABLE IF EXISTS `t_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `operation` varchar(64) DEFAULT NULL COMMENT '操作类型',
  `times` int(8) DEFAULT NULL COMMENT '耗费时长',
  `method` varchar(64) DEFAULT NULL COMMENT '请求方法',
  `params` varchar(500) DEFAULT NULL COMMENT '请求参数',
  `ip_address` varchar(32) DEFAULT NULL COMMENT 'ip地址',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建人',
  `created_username` varchar(32) DEFAULT NULL COMMENT '用户名',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_created_by` (`created_by`) USING BTREE COMMENT '创建人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='日志表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_sys_dept`
--

DROP TABLE IF EXISTS `t_sys_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_sys_dept` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `parent_id` bigint(20) DEFAULT NULL COMMENT '父部门Id',
  `name` varchar(128) DEFAULT NULL COMMENT '部门名称',
  `sort_num` int(4) DEFAULT NULL COMMENT '排序号',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建人',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '修改人',
  `updated_time` datetime DEFAULT NULL COMMENT '修改时间',
  `remark` varchar(128) DEFAULT NULL COMMENT '备注',
  `del_flag` int(2) DEFAULT NULL COMMENT '删除标记',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`) USING BTREE COMMENT 'parent_id 普通索引',
  KEY `idx_updated_by` (`updated_by`) USING BTREE COMMENT 'updated_by普通索引'
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='系统部门表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_sys_menu`
--

DROP TABLE IF EXISTS `t_sys_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_sys_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `parent_id` bigint(20) DEFAULT NULL COMMENT '父菜单ID',
  `name` varchar(64) DEFAULT NULL COMMENT '菜单名称',
  `url` varchar(128) DEFAULT NULL COMMENT '菜单url',
  `permissions` varchar(255) DEFAULT NULL COMMENT '菜单权限',
  `menu_type` int(2) DEFAULT NULL COMMENT '菜单类型',
  `icon` varchar(64) DEFAULT NULL COMMENT '菜单图标',
  `order_num` int(4) DEFAULT NULL COMMENT '排序号',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建人',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '修改人',
  `updated_time` datetime DEFAULT NULL COMMENT '修改时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `del_flag` int(2) DEFAULT NULL COMMENT '删除标记',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`) USING BTREE COMMENT 'parent_id 普通索引',
  KEY `idx_updated_by` (`updated_by`) USING BTREE COMMENT 'updated_by普通索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统菜单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_sys_role`
--

DROP TABLE IF EXISTS `t_sys_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_sys_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `parent_id` bigint(20) DEFAULT NULL COMMENT '父角色ID',
  `code` varchar(64) DEFAULT NULL COMMENT '角色自定义ID',
  `name` varchar(64) DEFAULT NULL COMMENT '角色名',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建人',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '修改人',
  `updated_time` datetime DEFAULT NULL COMMENT '修改时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `del_flag` int(2) DEFAULT NULL COMMENT '删除标记',
  PRIMARY KEY (`id`),
  KEY `uk_parent_id` (`parent_id`) USING BTREE COMMENT 'parent_id 普通索引',
  KEY `idx_updated_by` (`updated_by`) USING BTREE COMMENT 'updated_by 普通索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_sys_role_menu`
--

DROP TABLE IF EXISTS `t_sys_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_sys_role_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_menu_id` (`role_id`,`menu_id`) USING BTREE COMMENT 'role_id_menu_id 唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统角色菜单关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_sys_user`
--

DROP TABLE IF EXISTS `t_sys_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_sys_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部门Id',
  `username` varchar(64) DEFAULT NULL COMMENT '用户名（唯一）',
  `real_name` varchar(64) DEFAULT NULL COMMENT '真实姓名',
  `password` varchar(64) DEFAULT NULL COMMENT '密码',
  `email` varchar(64) DEFAULT NULL COMMENT '邮箱',
  `mobile_phone` varchar(16) DEFAULT NULL COMMENT '电话号码',
  `id_card` varchar(20) DEFAULT NULL COMMENT '身份号码',
  `status` int(4) DEFAULT NULL COMMENT '用户状态',
  `gender` int(2) DEFAULT NULL COMMENT '性别',
  `birth` datetime DEFAULT NULL COMMENT '出生日期',
  `avatar_id` bigint(20) DEFAULT NULL COMMENT '用户头像Id',
  `hobby` varchar(64) DEFAULT NULL COMMENT '爱好',
  `work_address` varchar(128) DEFAULT NULL COMMENT '工作详细地址',
  `province` varchar(16) DEFAULT NULL COMMENT '省',
  `city` varchar(16) DEFAULT NULL COMMENT '市',
  `district` varchar(32) DEFAULT NULL COMMENT '区',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建人',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '修改人',
  `updated_time` datetime DEFAULT NULL COMMENT '修改时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `del_flag` int(2) DEFAULT NULL COMMENT '删除标记',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`) USING BTREE COMMENT '用户名唯一索引',
  UNIQUE KEY `uk_email` (`email`) USING BTREE COMMENT '邮箱唯一索引',
  UNIQUE KEY `uk_mobile_phone` (`mobile_phone`) USING BTREE COMMENT '电话号码唯一索引',
  UNIQUE KEY `uk_id_card` (`id_card`) USING BTREE COMMENT '身份号码唯一索引',
  KEY `idx_dept_id` (`dept_id`) USING BTREE COMMENT '部门ID普通索引',
  KEY `idx_updated_by` (`updated_by`) USING BTREE COMMENT '修改人普通索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_sys_user_plus`
--

DROP TABLE IF EXISTS `t_sys_user_plus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_sys_user_plus` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户Id',
  `bank_name` varchar(64) DEFAULT NULL COMMENT '银行名',
  `bank_card` varchar(32) DEFAULT NULL COMMENT '银行号',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建人',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '修改人',
  `updated_time` datetime DEFAULT NULL COMMENT '修改时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `del_flag` int(2) DEFAULT NULL COMMENT '删除标记',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`) USING BTREE COMMENT 'user_id 普通索引',
  KEY `idx_updated_by` (`updated_by`) USING BTREE COMMENT 'updated_by普通索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统用户拓展表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_sys_user_role`
--

DROP TABLE IF EXISTS `t_sys_user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_sys_user_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` bigint(20) NOT NULL COMMENT '用户Id',
  `role_id` bigint(20) NOT NULL COMMENT '角色Id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_role_id` (`user_id`,`role_id`) USING BTREE COMMENT 'user_role_id 唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统用户角色关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_task`
--

DROP TABLE IF EXISTS `t_task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_task` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `task_name` varchar(64) DEFAULT NULL COMMENT '任务名称',
  `cron_expression` varchar(16) DEFAULT NULL COMMENT 'cron 表达式',
  `model_name` varchar(64) DEFAULT NULL COMMENT '模块名称',
  `method_name` varchar(64) DEFAULT NULL COMMENT '方法名称',
  `concurrent_flag` varchar(4) DEFAULT NULL COMMENT '并发标记',
  `status` varchar(2) DEFAULT NULL COMMENT '任务状态',
  `task_group` varchar(16) DEFAULT NULL COMMENT '任务组',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建人',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '修改人',
  `updated_time` datetime DEFAULT NULL COMMENT '修改时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `del_flag` int(2) DEFAULT NULL COMMENT '删除标记',
  PRIMARY KEY (`id`),
  KEY `idx_task_name` (`task_name`) USING BTREE COMMENT 'task_name 普通索引',
  KEY `idx_updated_by` (`updated_by`) USING BTREE COMMENT 'updated_by 普通索引'
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='定时任务表';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-09-12  7:30:44
