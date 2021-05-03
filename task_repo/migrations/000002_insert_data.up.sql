INSERT INTO tasks (name, description, time_limit, memory) VALUES
('Simple Task 1', 'Given two numbers, find their sum', 1000, 1000),
('Simple Task 2', 'Given an array of integers, find the one that appears an odd number of times', 500, 2000),
('Simple Task 3', 'There is an array with some numbers. All numbers are equal except for one. Try to find it!', 300, 3000)
;

INSERT INTO test_cases (task_id, test_data, answer) VALUES
(1, '1 2', '3'),
(1, '4 5', '9'),
(1, '10 11', '21'),
(1, '0 0', '0'),
(2, '1 2 1 2 2 3 3', '2'),
(2, '1 1 1 1 1 1 1', '1'),
(3, '7 1 1 1 1 2 1 1', '2'),
(3, '3 1 1 3', '3'),
(3, '9 0 0 0 0 0 0 2 0 0', '2')
;

INSERT INTO languages (name)  VALUES ('golang'), ('c++');
