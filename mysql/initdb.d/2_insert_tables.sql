USE test_db;

-- projects table test data
INSERT INTO projects
    (id, group_id, key_name, name, leader, default_assignee, created_date, updated_date)
values
    (1, 1, "AJA", "プロジェクト管理ツール", 1, 1, default, default),
    (2, 1, "KANR", "kanry", 1, 2, default, default),
    (3, 1, "DEF", "デフォルトツール", 2, 1, default, default),
    (4, 1, "PIED", "PiedPiper", 3, 1, default, default),
    (5, 2, "NEO", "NeoPiedPiper", 4, 1, default, default);
