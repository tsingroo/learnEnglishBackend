CREATE DATABASE mp_english_db;
use mp_english_db;
-- 用户表
CREATE TABLE mp_english_Users (
  Id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  Username VARCHAR(50),
  Password VARCHAR(50),
  WxOpenId VARCHAR(50),
  UserType INT -- 0:普通用户,1:老师
);
-- 课程表
CREATE TABLE mp_english_Lessons (
  Id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  LessonName VARCHAR(20),
  Memo VARCHAR(500),
  CreatorId INT -- 创建人
);

-- 课程内容(知识条目)
CREATE TABLE mp_english_Less_Knowledge (
  Id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  VoiceUrl VARCHAR(500),
  ImageUrl VARCHAR(500),
  LessonId INT -- 课程ID
);