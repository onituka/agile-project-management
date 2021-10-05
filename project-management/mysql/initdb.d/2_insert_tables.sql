USE test_db;

-- projects table test data
INSERT INTO projects
    (id, group_id, key_name, name, leader_id, default_assignee_id, created_at, updated_at)
values
    ("024d71d6-1d03-11ec-a478-0242ac180002", "024d78d6-1d03-11ec-a478-0242ac180002", "AJA", "プロジェクト管理ツール", "024d78d6-1d03-11ec-a478-0242ac184402", "024d78d6-1d03-11ec-a478-9242ac180002", default, default),
    ("024d7945-1d03-11ec-a478-0242ac180002", "023d76d6-1d03-11ec-a478-0242ac180002", "KANR", "kanry", "024d78d6-1d03-11ec-a478-0242ac180022", "024d78d6-1d03-11ec-a478-0242ac180502", default, default),
    ("024d61d4-1d03-11ec-a478-0242ac180002", "022d74d6-1d03-11ec-a478-0242ac180002", "DEF", "デフォルトツール", "024d78d6-1d03-11ec-a478-0245ac180002", "024d78d6-1d03-11ec-a488-0242ac180002", default, default),
    ("024d81d3-1d03-11ec-a478-0242ac180002", "021d72d6-1d03-11ec-a478-0242ac180002", "PIED", "PiedPiper", "024d78d6-1d03-11ec-a478-0242ac380002", "024d78d6-1d03-11ec-a478-0842ac180002", default, default),
    ("024d31d2-1d03-11ec-a478-0242ac180002", "025d71d6-1d03-11ec-a478-0242ac180002", "NEO", "NeoPiedPiper", "024d78d6-1d03-11ec-a478-0246ac180002", "024d78d6-1d03-11ec-a498-0242ac180002", default, default);
