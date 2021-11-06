USE test_db;

-- products table test data
INSERT INTO products
    (id, group_id, name, leader_id, created_at, updated_at)
values
    ("4495c574-34c2-4fb3-9ca4-3a7c79c267a6", "024d78d6-1d03-11ec-a478-0242ac180002", "プロジェクト管理ツール", "024d78d6-1d03-11ec-a478-0242ac184402", "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000"),
    ("4487c574-34c2-4fb3-9ca4-3a7c79c267a6", "024d78d6-1d03-11ec-a478-0242ac180002", "test", "024d78d6-1d03-11ec-a478-0242ac184402", "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000");

-- projects table test data
INSERT INTO projects
    (id, group_id, key_name, name, leader_id, default_assignee_id, created_at, updated_at)
values
    ("024d71d6-1d03-11ec-a478-0242ac180002", "024d78d6-1d03-11ec-a478-0242ac180002", "AAA", "管理ツール1", "024d78d6-1d03-11ec-a478-0242ac184402", "024d78d6-1d03-11ec-a478-9242ac180002", default, default),
    ("024d7945-1d03-11ec-a478-0242ac180002", "024d78d6-1d03-11ec-a478-0242ac180002", "BBB", "管理ツール2", "024d78d6-1d03-11ec-a478-0242ac180022", "024d78d6-1d03-11ec-a478-0242ac180502", default, default),
    ("024d61d4-1d03-11ec-a478-0242ac180002", "024d78d6-1d03-11ec-a478-0242ac180002", "CCC", "管理ツール3", "024d78d6-1d03-11ec-a478-0245ac180002", "024d78d6-1d03-11ec-a488-0242ac180002", default, default);
