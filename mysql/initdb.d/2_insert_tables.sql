USE test_db;

-- projects table test data
INSERT INTO projects
    (id, group_id, key_name, name, leader_id, default_assignee_id, created_at, updated_at)
values
    ("024d71d6-1d03-11ec-a478-0242ac180002", "024d78d6-1d03-11ec-a478-0242ac180002", "AJA", "プロジェクト管理ツール", 1, 1, default, default),
    ("024d7945-1d03-11ec-a478-0242ac180002", "023d76d6-1d03-11ec-a478-0242ac180002", "KANR", "kanry", 1, 2, default, default),
    ("024d61d4-1d03-11ec-a478-0242ac180002", "022d74d6-1d03-11ec-a478-0242ac180002", "DEF", "デフォルトツール", 2, 1, default, default),
    ("024d81d3-1d03-11ec-a478-0242ac180002", "021d72d6-1d03-11ec-a478-0242ac180002", "PIED", "PiedPiper", 3, 1, default, default),
    ("024d31d2-1d03-11ec-a478-0242ac180002", "025d71d6-1d03-11ec-a478-0242ac180002", "NEO", "NeoPiedPiper", 4, 1, default, default);
