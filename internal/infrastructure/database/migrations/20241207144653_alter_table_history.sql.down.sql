alter table history RENAME COLUMN prediction to result;

alter TABLE history drop column recommendation;

ALTER Table history MODIFY COLUMN image VARCHAR(256);

alter table users modify column user_id int;