CREATE TABLE mp_english_subjects(
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  description VARCHAR(50),
  image_url varchar(500),
  memo varchar(500),
  display_order int,
  creator_id int
);

CREATE TABLE mp_english_subject_knowledge(
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  description VARCHAR(500),
  image_url VARCHAR(500),
  question VARCHAR(500),
  answer VARCHAR(500),
  display_order int,
  subject_id int
);