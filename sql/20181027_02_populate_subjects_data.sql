-- subjects data
INSERT INTO mp_english_subjects 
VALUES(1, 'shapes are everywhere', '', 'all about shapes', 1, 1);
INSERT INTO mp_english_subjects 
VALUES(2, 'my family', '', 'my family', 2, 1);

-- subject knowledge data
INSERT INTO mp_english_subject_knowledge(description, image_url, question, answer, display_order, subject_id)
VALUES('circle', 'http://ph1cnfby2.bkt.clouddn.com/18-10-27/38775799.jpg', 
'what shape is this', 'it is a circle', 1, 1);
INSERT INTO mp_english_subject_knowledge(description, image_url, question, answer, display_order, subject_id)
VALUES('crescent', 'http://ph1cnfby2.bkt.clouddn.com/18-10-27/82601651.jpg', 
'what shape is this', 'it is a crescent', 2, 1);
INSERT INTO mp_english_subject_knowledge(description, image_url, question, answer, display_order, subject_id)
VALUES('diamond', 'http://ph1cnfby2.bkt.clouddn.com/18-10-27/38052335.jpg', 
'what shape is this', 'it is a diamond', 3, 1);
INSERT INTO mp_english_subject_knowledge(description, image_url, question, answer, display_order, subject_id)
VALUES('rectangle', 'http://ph1cnfby2.bkt.clouddn.com/18-10-27/39961790.jpg', 
'what shape is this', 'it is a rectangle', 4, 1);
INSERT INTO mp_english_subject_knowledge(description, image_url, question, answer, display_order, subject_id)
VALUES('square', 'http://ph1cnfby2.bkt.clouddn.com/18-10-27/65954920.jpg', 
'what shape is this', 'it is  square', 5, 1);
INSERT INTO mp_english_subject_knowledge(description, image_url, question, answer, display_order, subject_id)
VALUES('triangle', 'http://ph1cnfby2.bkt.clouddn.com/18-10-27/79299394.jpg', 
'what shape is this', 'it is a triangle', 6, 1);
INSERT INTO mp_english_subject_knowledge(description, image_url, question, answer, display_order, subject_id)
VALUES('oval', 'http://ph1cnfby2.bkt.clouddn.com/18-10-27/44807990.jpg', 
'what shape is this', 'it is an oval', 7, 1);
