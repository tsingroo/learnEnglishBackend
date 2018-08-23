CREATE DATABASE mp_english_db CHARACTER SET=utf8;
use mp_english_db;
-- 用户表
CREATE TABLE mp_english_users (
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  username VARCHAR(50),
  password VARCHAR(50),
  wx_open_id VARCHAR(50),
  user_type INT -- 0:普通用户,1:老师
);
-- 课程表
CREATE TABLE mp_english_lessons (
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  lesson_name VARCHAR(20),
  memo VARCHAR(500),
  creator_id INT -- 创建人
);

-- 课程内容(知识条目)
CREATE TABLE mp_english_less_knowledge (
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  voice_url VARCHAR(500),
  image_url VARCHAR(500),
  lesson_id INT -- 课程ID
);