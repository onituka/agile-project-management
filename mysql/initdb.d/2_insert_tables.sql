USE test_db;

-- projects table test data
INSERT INTO projects
    (id, group_id, key_name, name, leader, default_assignee, created_date, updated_date)
values
    ("024d71d6-1d03-11ec-a478-0242ac180002", 1, "AJA", "プロジェクト管理ツール", 1, 1, default, default),
    ("024d7945-1d03-11ec-a478-0242ac180002", 1, "KANR", "kanry", 1, 2, default, default),
    ("024d61d4-1d03-11ec-a478-0242ac180002", 1, "DEF", "デフォルトツール", 2, 1, default, default),
    ("024d81d3-1d03-11ec-a478-0242ac180002", 1, "PIED", "PiedPiper", 3, 1, default, default),
    ("024d31d2-1d03-11ec-a478-0242ac180002", 2, "NEO", "NeoPiedPiper", 4, 1, default, default);
