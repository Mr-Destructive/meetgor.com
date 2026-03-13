---
date: 2025-09-06T00:00:00Z
description: 'Database System: Relational Model and Algebra'
image_url: https://i.ytimg.com/vi/7NPIENPr-zk/hqdefault.jpg
link: https://youtu.be/7NPIENPr-zk
newsletter: 2025-09-06-techstructive-weekly-58
slug: database-system-relational-model-and-algebra
source: newsletter
status: published
tags:
    - links
title: 'Database System: Relational Model and Algebra'
type: links
---


## Commentary

- Prepared a detailed notes as I was watching
- data model - how to define the relation, rules
- schema - defining the specifics of the data model, details
- relational (the primary kind, we only think of this as real dbs) key value - graph - document - column  - array - and more
- Initial idea: Writing cobol in 1970s to get data, literally telling the db to how to get the data. Not ideal as you don't know if that would be ideal way to get data in all cases, example artist and album, you write for each artist, find albums, but what if the number of albums by each artist keep growing and more people ar accessing different artists data together, so not sure how to get data for any query. In SQL however, we dont say how to get data, we say what and from where to get data
- Relational Model
- Structure
- Integrity (constraints)
- Manipulation
- Its upon the database system how to query/mutate the data
- Components
- Database storage(bits)
- Physical storage (pages, files, etc)
- Logical Schema (schema, constraints)
- External Schema (common table data)
- Application
- Everything below application is a db system
- Relation is a set, relation of attributes that represent entities
- Tuple is a set of attribute values in relation. Primary key, uniquely identify a tuple in a relation. Foreign key, related tuples(attributes) across relations. Constraints, conditions must hold for any tuple in a relation.
- Relational Algebra
- Select (where conditions)
- projection (select with what to extract optional modify the selected values)
- union (all, must have same number of attributes)
- intersection (same but common)
- difference (same but difference)
- product (cartesian product, cross join)
- join (natural, without params, common ones, same attribute name, can use params like using **on**)
