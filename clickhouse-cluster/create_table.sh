CREATE TABLE default.test_table ON CLUSTER '{cluster}'    (
    id UInt64,
    column1 String
  )
  ENGINE = ReplicatedMergeTree('/clickhouse/tables/test_table', '{replica}' )
  ORDER BY (id);


insert into default.test_table (id, column1) values (21, 'hello'), (22, 'world')

