CREATE KEYSPACE swapi
  WITH REPLICATION = {
   'class' : 'SimpleStrategy',
   'replication_factor' : 1
  };

create table swapi.planet
(
    id      uuid primary key,
    climate text,
    name    text,
    terrain text
)
    with caching = {'keys': 'ALL', 'rows_per_partition': '20000'}
     and compaction = {'max_threshold': '32', 'min_threshold': '4', 'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy'}
     and compression = {'class': 'org.apache.cassandra.io.compress.LZ4Compressor', 'chunk_length_in_kb': '64'}
     and dclocal_read_repair_chance = 0.1;

create materialized view swapi.planet_by_name
as
select name, id, climate, terrain
from planet
where name IS NOT NULL
primary key (name, id)
with caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
 and compaction = {'max_threshold': '32', 'min_threshold': '4', 'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy'}
 and compression = {'class': 'org.apache.cassandra.io.compress.LZ4Compressor', 'chunk_length_in_kb': '64'}
 and dclocal_read_repair_chance = 0.1;
