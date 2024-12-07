alter table history RENAME COLUMN result to prediction;

alter TABLE history add column recommendation TEXT after prediction;

ALTER Table history MODIFY COLUMN image JSON;

ALTER TABLE history DROP FOREIGN KEY history_ibfk_1;


alter table users modify column user_id int AUTO_INCREMENT;