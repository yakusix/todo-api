-- INSERT: добавить задачу
INSERT INTO tasks (title, description) VALUES ('написать резюме', 'очень важно');

-- SELECT: найти задачу по id
SELECT * FROM tasks WHERE id = 2;

-- UPDATE: сменить статус
UPDATE tasks SET status = 'done' WHERE id = 1;

-- DELETE: удалить задачу
DELETE FROM tasks WHERE id = 1;