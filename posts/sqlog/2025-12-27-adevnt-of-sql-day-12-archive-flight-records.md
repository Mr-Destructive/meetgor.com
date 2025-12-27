---
type: "sqlog"
title: "Advent of SQL 2025 Day 12: Archive Flight Records"
slug: "advent-of-sql-2025-day-12"
date: 2025-12-27
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Advent of SQL - Day 12, Archive Flight Records

We are on Day 12! Phew its almost done! Just 3 days more!

Let's get the SQL!

```sql
DROP TABLE IF EXISTS archive_records;

CREATE TABLE archive_records (
    id INT PRIMARY KEY,
    title TEXT,
    description TEXT
);

INSERT INTO archive_records (id, title, description) VALUES
(1, 'Flight Stabilization Prototype Analysis', 'This report details the latest advancements in stabilizing aerial maneuvers for enchanted sleighs. Initial tests yielded promising results, showcasing a marked decrease in turbulence during airborne navigation.'),
(2, 'Lift Calibration Incident Log', 'During a routine lift calibration test, the sleigh experienced an unexpected upward surge, causing it to hover dangerously close to the workshop ceiling. Subsequent analysis revealed a miscalculation in the weight distribution formula, prompting a thorough review of all aerodynamic coefficients.'),
(3, 'Aerial Aspirations: The Great Flop', 'Despite the initial excitement surrounding the design of our feather-laden airborne contraption, the prototype proved less than buoyant. The unexpected descent resulted in a rather spectacular cloud of glitter and twigs, serving as a vivid reminder that not all dreams of flight take wing as intended.'),
(4, 'Optimized Sled Dynamics', 'This document explores advanced mechanics behind sled propulsion and movement efficiency. Through extensive calculations and enchanted material tests, the elves aim to refine turning capabilities and enhance downhill speed.'),
(5, 'Reindeer Harness Design Flaws', 'The latest prototype of our reindeer harness was found to inhibit mobility, causing undue strain on the animals during testing. Observations indicated that the weight distribution was poorly calibrated, necessitating a complete redesign for optimal comfort and performance.'),
(6, 'Streamlined Cargo Routing System', 'In our latest endeavor, we have implemented a magical algorithm to optimize the route taken by our toy-laden sleighs. This groundbreaking update minimizes transport time between the workshop and delivery points, ensuring that holiday cheer reaches every home even faster!'),
(7, 'Toy Durability Testing Protocols', 'The elves meticulously conducted stress evaluations on the latest toy prototypes to determine their resilience under various conditions. Initial findings indicate that while some designs withstood rigorous play, others required reinforcements to avoid premature wear and tear.'),
(8, 'Weather Resistance Breakthrough: Frost Shielding', 'In our recent experiments, we discovered an innovative composite material that effectively repels moisture while withstanding extreme cold. This newfound frost shielding could revolutionize our outdoor toys, ensuring they remain both functional and enchanting, even in the harshest winter conditions.'),
(9, 'Safety Compliance Check Overview', 'In the pursuit of enchantment and joy, this document outlines the mandatory safety compliance measures for all workshop operations. Each elf must adhere strictly to these guidelines to ensure the safe transport and handling of our delicate prototypes, thus preventing any unforeseen magical mishaps.'),
(10, 'Sleigh Skim Mechanism Upgrade', 'The experiment involved enchanting the underside of the sleigh with a whispering wind charm to achieve unprecedented speeds. Results were alarming, leading to uncontrollable flight trajectories and sudden descents—definitely do not attempt again.'),
(11, 'Caution: Enchanted Toy Prototype', 'This design incorporates a reactive magic component that may unpredictably animate in the presence of mischief. Ensure all test environments are secured against spontaneous giggles and potential chaos.'),
(12, 'Workshop Experiment Safety Checklist', 'Before embarking on any workshop experiments, ensure all safety goggles are securely fastened to prevent debris from interfered visions. Always double-check that the workspace is free of clutter, as unexpected accidents can arise from even the smallest flurry of trinkets and tools.'),
(13, 'Cocoa Bean Roasting Innovations', 'This review explores the latest techniques in roasting cocoa beans to achieve unparalleled flavor profiles. Adjustments to temperature and timing have led to a delightful spectrum of aromas, promising to elevate our confectionery creations to new heights.'),
(14, 'Intricate Snowflake Ornament Design', 'This design blueprint outlines the geometric intricacies of a multi-faceted snowflake ornament, emphasizing a balance between elegance and structural integrity. Each arm is meticulously patterned to reflect light, creating a shimmering effect that dances with the seasons, while ensuring optimal symmetry for enchanting visual appeal.'),
(15, 'Gift Box Assembly Prototype', 'This prototype outlines the intricate process of assembling the enchanted gift boxes designed to withstand the whims of time and space. Each step must be meticulously executed to ensure that every box not only sparkles with joy but also maintains its magical properties through every unwrapping.'),
(16, 'Wrap-It-Up: Innovative Designs', 'This experimental report explores various materials and techniques for creating enchanted wrapping paper that enhances the gift-giving experience. Initial findings suggest that incorporation of shimmering elven dust can amplify the aesthetic appeal while maintaining structural integrity during airborne delivery.'),
(17, 'Magical Confection Fusion Results', 'The experimental concoction blended sugar crystals with essence of starlight, resulting in a luminescent treat that sparkles enchantingly. However, a curious side effect was noted: excessive giggling among taste testers, raising questions about potential airborne laughter.'),
(18, 'Luminous Ornament Crafting Techniques', 'In our continuous pursuit of radiance, this document outlines innovative methods for creating ornaments that glow with enchantment. Engaging both traditional techniques and modern enchantments, each design is meant to instill joy and sparkle during the festive season.'),
(19, 'Magical Energy Conduction Analysis', 'Recent experiments have shown that the flow of magical energy through crystalline conduits behaves unpredictably under varying lunar phases. Further investigation into the correlation between ambient mana levels and energy stability is necessary to optimize enchantment potency.'),
(20, 'Elven Workshop Organization Protocols', 'The implementation of open shelving systems has significantly increased accessibility to essential materials, thus enhancing workflow efficiency. Furthermore, the organization of tools into color-coded bins ensures that each elf can swiftly locate their required implements without disrupting the harmony of the workshop.');

```
 
We have just one table and a couple of text like columns. That's it, looks like a string searching problem.

Let's head to the problem statement!


## Problem

> Using the `archive_records` table, search both the `title` and `description` fields for the term "fly". Make sure that you also match for words like "flying", "flight", etc. Boost the results where the term appears in the title and lastly, rank the results by relevance (most relevant first). Provide the elves the top 5 most relevant archived records back.

It is a text search use case problem indeed!

We need to find and rank the records matching the word `fly`, `flying`, `flight` like those. Kind of tricky if we miss any terms that are not hard-coded.

We need to boost the search term in `title`, so there is more weightage if the term appears in title than in description. Makes sense!

Let's start simple and move into full-text-search in SQLite!


### Simple String Matching

We start by a simple nested `CASE WHEN THEN` condition. We check if the `title` has `fly`, `flying`, `flight`, etc then we set the rank as `2` and add the score from `description` as `1` if the same terms appear in description.

So,
- If the search term (fly, flight, etc) appears **only** in title, score is `2`
- If the search term appears **only** in description, score is `1`
- If the search term appears in **both** title and description, the score is `3` since we are adding the scores.
- If the search term doesn't appear at all, then the score remains `0`.

We simply assign the score based on the appearance of the search term and then order the result based on the computed `rank` and list the top 5.

```sql
 SELECT 
    id,
    title,
    description,
    (
        CASE WHEN LOWER(title) LIKE '%fly%' OR LOWER(title) LIKE '%flight%' OR LOWER(title) LIKE '%flying%' THEN 2 ELSE 0 END +
        CASE WHEN LOWER(description) LIKE '%fly%' OR LOWER(description) LIKE '%flight%' OR LOWER(description) LIKE '%flying%' THEN 1 ELSE 0 END
    ) AS rank
FROM archive_records
ORDER BY rank DESC, id ASC
LIMIT 5;
```

However, we can limit the search space by eliminating the computation of ranking all the records with a `WHERE` clause since it can optimise only the relevant records before giving out the result set.

Hence, we filter only in the cases where the `title` and `description` have the relevant word and then use the condition to order by the `rank` score that we computed and list the top `5` ones.

```sql
SELECT 
    id,
    title,
    description,
    (
        -- Title matches worth 2 points (boosted relevance)
        CASE WHEN LOWER(title) LIKE '%fly%' OR LOWER(title) LIKE '%flight%' OR LOWER(title) LIKE '%flying%' THEN 2 ELSE 0 END +
        -- Description matches worth 1 point
        CASE WHEN LOWER(description) LIKE '%fly%' OR LOWER(description) LIKE '%flight%' OR LOWER(description) LIKE '%flying%' THEN 1 ELSE 0 END
    ) AS rank
FROM archive_records
WHERE 
    LOWER(title) LIKE '%fly%' OR 
    LOWER(title) LIKE '%flight%' OR 
    LOWER(title) LIKE '%flying%' OR
    LOWER(description) LIKE '%fly%' OR 
    LOWER(description) LIKE '%flight%' OR 
    LOWER(description) LIKE '%flying%'
ORDER BY rank DESC, id ASC
LIMIT 5;
```

We can wrap it in a CTE, but that doesn't does anything differently though.

With CTE as below, we assign a common score like `1` or `0` based on the `title` or description match. Then in the outer query, we can assign a weight to the rank if the `title` matched as `2` or `description` matched as `1` and add it up to get the same `rank` for that record.

```sql
WITH search_results AS (
    SELECT 
        id,
        title,
        description,
        CASE WHEN LOWER(title) LIKE '%fly%' OR LOWER(title) LIKE '%flight%' OR LOWER(title) LIKE '%flying%' 
            THEN 1 ELSE 0 END AS title_match,
        CASE WHEN LOWER(description) LIKE '%fly%' OR LOWER(description) LIKE '%flight%' OR LOWER(description) LIKE '%flying%' 
            THEN 1 ELSE 0 END AS desc_match
    FROM archive_records
    WHERE 
        LOWER(title) LIKE '%fly%' OR LOWER(title) LIKE '%flight%' OR LOWER(title) LIKE '%flying%' OR
        LOWER(description) LIKE '%fly%' OR LOWER(description) LIKE '%flight%' OR LOWER(description) LIKE '%flying%'
)
SELECT 
    id,
    title,
    description,
    (title_match * 2 + desc_match * 1) AS rank
FROM search_results
ORDER BY rank DESC, id ASC
LIMIT 5;
```

### Assign score based on matching frequency

We can also make it better by counting how many times the term appeared in the relevant column and then compute a cumulative rank. This way it becomes more aggressive on the most relevant docs only.

So, we write something like this.

- Compute the length of the full column example as title with `LENGTH(LOWER(title)` then subtract the length of the word left after removing `fly`, or `flight`, or others. It calculates the difference in length between the original title and the one with 'fly' removed. Each 'fly' is 3 characters, so the difference divided by 3 gives the number of occurrences of 'fly'. This `(LENGTH(LOWER(title)) - LENGTH(REPLACE(LOWER(title), 'fly', ''))) / 3
` counts how many times 'fly' appears. Similarly for `flight` and `flying` it is `6`, and so on.
- If that appears in title it is multiplied by `2`
- If it appears in description, it is kept as is (multiplied by `1`) but you can weight it accordingly, as there could be multiple column to weigh.

```sql
SELECT 
    id,
    title,
    description,
    (

        ((LENGTH(LOWER(title)) - LENGTH(REPLACE(LOWER(title), 'fly', ''))) / 3 +
         (LENGTH(LOWER(title)) - LENGTH(REPLACE(LOWER(title), 'flight', ''))) / 6 +
         (LENGTH(LOWER(title)) - LENGTH(REPLACE(LOWER(title), 'flying', ''))) / 6
        ) * 2
        +

        ((LENGTH(LOWER(description)) - LENGTH(REPLACE(LOWER(description), 'fly', ''))) / 3 +
         (LENGTH(LOWER(description)) - LENGTH(REPLACE(LOWER(description), 'flight', ''))) / 6 +
         (LENGTH(LOWER(description)) - LENGTH(REPLACE(LOWER(description), 'flying', ''))) / 6
        )
    ) AS rank
FROM archive_records
ORDER BY rank DESC, id ASC
LIMIT 5;
```

You can always add a where clause to reduce the search space upfront:

```sql
-- Simplified frequency-based ranking
SELECT 
    id,
    title,
    description,
    (
        ((LENGTH(LOWER(title)) - LENGTH(REPLACE(LOWER(title), 'fly', ''))) / 3 +
         (LENGTH(LOWER(title)) - LENGTH(REPLACE(LOWER(title), 'flight', ''))) / 6 +
         (LENGTH(LOWER(title)) - LENGTH(REPLACE(LOWER(title), 'flying', ''))) / 6
        ) * 2
        +
        ((LENGTH(LOWER(description)) - LENGTH(REPLACE(LOWER(description), 'fly', ''))) / 3 +
         (LENGTH(LOWER(description)) - LENGTH(REPLACE(LOWER(description), 'flight', ''))) / 6 +
         (LENGTH(LOWER(description)) - LENGTH(REPLACE(LOWER(description), 'flying', ''))) / 6
        )
    ) AS rank
FROM archive_records
WHERE 
    LOWER(title) LIKE '%fly%' OR LOWER(title) LIKE '%flight%' OR LOWER(title) LIKE '%flying%' OR
    LOWER(description) LIKE '%fly%' OR LOWER(description) LIKE '%flight%' OR LOWER(description) LIKE '%flying%'
ORDER BY rank DESC, id ASC
LIMIT 5;
```

This would give us similar results but slightly different as we are assigning rank on the number of times each term appears and not just if it appears once or not.

We can even make the list of words and their weight as a CTE and dynamically use it in the actual query:

```sql
WITH keywords AS (
    SELECT 'fly' AS term, 2.0 AS title_weight, 1.0 AS desc_weight
    UNION ALL
    SELECT 'flight', 2.0, 1.0
    UNION ALL
    SELECT 'flying', 2.0, 1.0
),
ranked AS (
    SELECT 
        a.id,
        a.title,
        a.description,
        SUM(
            ((LENGTH(LOWER(a.title)) - LENGTH(REPLACE(LOWER(a.title), k.term, ''))) / LENGTH(k.term)) * k.title_weight +
            ((LENGTH(LOWER(a.description)) - LENGTH(REPLACE(LOWER(a.description), k.term, ''))) / LENGTH(k.term)) * k.desc_weight
        ) AS rank
    FROM archive_records a
    CROSS JOIN keywords k
    GROUP BY a.id
)
SELECT id, title, description, rank
FROM ranked
ORDER BY rank DESC, id ASC
LIMIT 5;

```

Here we have 2 CTEs now:
- `keywords` defining the words to search for and its weight
- `ranked` defining the computed sum of weights with frequency on the relevant columns.

As you can see, by just adding the keyword in the CTE along with its weight, the rest of the query can work without changing anything.


### Full Text Search

We can now also look at the [FTS](https://sqlite.org/fts5.html) or Full Text Search in SQLite.

In SQLite, we can create a `VIRTUAL TABLE` which is like
- A table computed on the fly
- Doesn't exist as physical entity in the database

We have [fts5](https://sqlite.org/fts5.html) which we can use to match, equality, or expressions in functions to search for text in columns.

```sql
CREATE VIRTUAL TABLE IF NOT EXISTS archive_fts USING fts5(
    title, 
    description,
    content=archive_records
);
```

We need to provide the columns that we want to search against. In this case we want to search for `title` and `description` as the column. The `content=archive_records`, links the FTS table to a real table `archive_records` instead of storing its own copy of all text, the FTS table indexes the data from archive_records.

Then we want to insert all the records from the `archive_records` for it to make aware of the records existing in the actual table.

```sql
INSERT OR IGNORE INTO archive_fts(rowid, title, description)
SELECT id, title, description FROM archive_records;
```

Then we can query it like so:

```sql
SELECT * FROM archive_fts WHERE archive_fts MATCH 'fly';
```

We have a couple of option to search against like
1. MATCH
2. Boolean Operator
3. Expressions and Wildcards

We simply can filter the records with this 

```sql
SELECT *
FROM archive_fts
JOIN archive_records ON archive_fts.rowid = archive_records.id
WHERE archive_fts MATCH 'fly OR flight OR flying';
```

We have added `MATCH 'fly OR flight OR flying'` to only limit the search space on those keywords.


```sql
SELECT 
    archive_records.id,
    archive_records.title,
    archive_records.description,
    (
        CASE WHEN archive_records.title LIKE '%fly%' OR archive_records.title LIKE '%flight%' OR archive_records.title LIKE '%flying%' THEN 2 ELSE 0 END +
        CASE WHEN archive_records.description LIKE '%fly%' OR archive_records.description LIKE '%flight%' OR archive_records.description LIKE '%flying%' THEN 1 ELSE 0 END
    ) AS rank
FROM archive_fts
JOIN archive_records ON archive_fts.rowid = archive_records.id
WHERE archive_fts MATCH 'fly OR flight OR flying'
ORDER BY rank DESC, archive_records.id ASC
LIMIT 5;
```

This simply will give us the relevant results by doing a fuzzy search which is way better then `LIKE` or `%` wildcard operators in the columns.

So, that are the approaches I like taking on day 12.

It was fun! Working with full text search for the first time.

On to day 13!


