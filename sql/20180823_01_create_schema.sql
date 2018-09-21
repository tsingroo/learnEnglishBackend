CREATE DATABASE mp_english_db CHARACTER SET=utf8;
use mp_english_db;
-- 用户表
CREATE TABLE mp_english_users (
  ID INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  USERNAME VARCHAR(50),
  PASSWORD VARCHAR(50),
  WX_OPEN_ID VARCHAR(50),
  USER_TYPE INT -- 0:普通用户,1:老师
);
-- 课程表
CREATE TABLE mp_english_lessons (
  ID INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  LESSON_NAME VARCHAR(20),
  IMAGE_URL VARCHAR(500),
  MEMO VARCHAR(500),
  CREATOR_ID INT -- 创建人
);

-- 课程内容(知识条目)
CREATE TABLE mp_english_less_knowledge (
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  word VARCHAR(500),
  voice_url VARCHAR(500),
  image_url VARCHAR(500),
  lesson_id INT -- 课程ID
);