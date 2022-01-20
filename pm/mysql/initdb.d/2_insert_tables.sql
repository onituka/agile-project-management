USE test_db;

-- products table test data
INSERT INTO products
    (id, group_id, name, leader_id, created_at, updated_at)
values
    ("4495c574-34c2-4fb3-9ca4-3a7c79c267a6", "024d78d6-1d03-11ec-a478-0242ac180002", "プロジェクト管理ツール", "024d78d6-1d03-11ec-a478-0242ac184402", "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000"),
    ("4487c574-34c2-4fb3-9ca4-3a7c79c267a6", "024d78d6-1d03-11ec-a478-0242ac180002", "test", "024d78d6-1d03-11ec-a478-0242ac184402", "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000"),
    ("5487c574-34c2-4fb3-9ca4-3a7c79c267a6", "124d78d6-1d03-11ec-a478-0242ac180002", "test2", "024d78d6-1d03-11ec-a478-0242ac184402", "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000");

-- product_notes table test data
INSERT INTO product_notes
    (id, product_id, group_id, title, content, created_by, updated_by, created_at, updated_at)
values
    ("52dfc0d0-748e-11ec-88fd-acde48001122", "4495c574-34c2-4fb3-9ca4-3a7c79c267a6", "024d78d6-1d03-11ec-a478-0242ac180002", "ノート", "note", "024d78d6-1d03-11ec-a478-0242ac184402", "024d78d6-1d03-11ec-a478-0242ac184402", "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000"),
    ("62dfc0d0-748e-11ec-88fd-acde48001122", "4495c574-34c2-4fb3-9ca4-3a7c79c267a6", "024d78d6-1d03-11ec-a478-0242ac180002", "test", "test", "024d78d6-1d03-11ec-a478-0242ac184402", "024d78d6-1d03-11ec-a478-0242ac184402", "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000"),
    ("72dfc0d0-748e-11ec-88fd-acde48001122", "5487c574-34c2-4fb3-9ca4-3a7c79c267a6", "124d78d6-1d03-11ec-a478-0242ac180002", "test2", "note", "024d78d6-1d03-11ec-a478-0242ac184402", "024d78d6-1d03-11ec-a478-0242ac184402", "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000");

-- projects table test data
INSERT INTO projects
    (id, product_id, group_id, key_name, name, leader_id, default_assignee_id, trashed_at, created_at, updated_at)
values
    ("024d71d6-1d03-11ec-a478-0242ac180002", "4495c574-34c2-4fb3-9ca4-3a7c79c267a6", "024d78d6-1d03-11ec-a478-0242ac180002", "AAA", "管理ツール1", "024d78d6-1d03-11ec-a478-0242ac184402", "024d78d6-1d03-11ec-a478-9242ac180002", NULL, "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000"),
    ("024d7945-1d03-11ec-a478-0242ac180002", "4495c574-34c2-4fb3-9ca4-3a7c79c267a6", "024d78d6-1d03-11ec-a478-0242ac180002", "BBB", "管理ツール2", "024d78d6-1d03-11ec-a478-0242ac180022", "024d78d6-1d03-11ec-a478-0242ac180502", NULL, "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000"),
    ("024d61d4-1d03-11ec-a478-0242ac180002", "4487c574-34c2-4fb3-9ca4-3a7c79c267a6", "024d78d6-1d03-11ec-a478-0242ac180002", "CCC", "管理ツール3", "024d78d6-1d03-11ec-a478-0245ac180002", "024d78d6-1d03-11ec-a488-0242ac180002", NULL, "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000");

-- project_notes tables test data
INSERT INTO project_notes
    (id, product_id, project_id, group_id, title, content, created_by, updated_by, created_at, updated_at)
VALUES
    ("777d71d6-1d03-11ec-a478-0242ac180002", "4495c574-34c2-4fb3-9ca4-3a7c79c267a6", "024d71d6-1d03-11ec-a478-0242ac180002", "024d78d6-1d03-11ec-a478-0242ac180002", "管理ツール調査", "testデータ１", "024d78d6-1d03-11ec-a478-0242ac184402", "024d78d6-1d03-11ec-a478-9242ac180002", "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000"),
    ("888d71d6-1d03-11ec-a478-0242ac180002", "4495c574-34c2-4fb3-9ca4-3a7c79c267a6", "024d71d6-1d03-11ec-a478-0242ac180002", "024d78d6-1d03-11ec-a478-0242ac180002", "売り上げまとめ", "testデータ2", "024d78d6-1d03-11ec-a478-0242ac184402", "024d78d6-1d03-11ec-a478-0242ac184402", "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000"),
    ("999d71d6-1d03-11ec-a478-0242ac180002", "4495c574-34c2-4fb3-9ca4-3a7c79c267a6", "024d71d6-1d03-11ec-a478-0242ac180002", "024d78d6-1d03-11ec-a478-0242ac180002", "言語まとめ", "testデータ3", "024d78d6-1d03-11ec-a478-0242ac184502", "024d78d6-1d03-11ec-a478-9242ac180002", "2021-11-05 00:00:00.000000", "2021-11-05 00:00:00.000000");
