---
type: "sqlog"
title: "Advent of SQL 2025 Day 15: Confirmation Phrase Dispatches"
slug: "advent-of-sql-2025-day-15"
date: 2025-12-28
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Advent of SQL, Day 15 - Confirmation Phrase Dispatches

We are on the final day of Advent of SQL!

I can't believe it, I completed it! (with some help of course, I ~~don't~~ didn't know SQL very well. But this 15 days flipped it around.). I feel good and fresh to start the year to go deep in SQLite and databases!

Let's solve the last day first, if it has somethings reamining to teach us!

We need to make a few changes.

SQLite doesn't have `->>` operator for extracting JSON values from the columns, so we need to use `json_extract` instead.

So, replace this 9th line from and with:

```diff
-    marker_letter TEXT GENERATED ALWAYS AS (payload ->> 'marker') STORED,
+    marker_letter TEXT GENERATED ALWAYS AS (json_extract(payload, '$.marker')) STORED,
```

Once we have it, we also need to remove the marker `::jsonb` from the each insert row, we can do that with `sed` or any other utility you like.

```
sed "s/'\(.*\)'::jsonb/'\1'/g" day15-inserts.sql > day15-inserts-sqlite.sql
```

```sql
DROP TABLE IF EXISTS incoming_dispatches;
DROP TABLE IF EXISTS system_dispatches;

CREATE TABLE system_dispatches (
    id SERIAL PRIMARY KEY,
    system_id TEXT NOT NULL,
    dispatched_at TIMESTAMP NOT NULL,
    payload JSONB NOT NULL,
    marker_letter TEXT GENERATED ALWAYS AS (json_extract(payload, '$.marker')) STORED,
    UNIQUE (system_id, dispatched_at, payload)
);

CREATE TABLE incoming_dispatches (
    system_id TEXT,
    dispatched_at TIMESTAMP,
    payload JSONB
);

INSERT INTO system_dispatches (system_id, dispatched_at, payload) VALUES
('SYS-0081', '2025-12-21T06:02:26', '{"marker": "T", "source": "secondary"}'),
('SYS-0137', '2025-12-19T06:03:21', '{"marker": "D", "source": "secondary"}'),
('SYS-0237', '2025-12-19T06:23:37', '{"marker": "D", "source": "secondary"}'),
('SYS-0006', '2025-12-24T18:10:16', '{"marker": "K", "source": "secondary"}'),
('SYS-0170', '2025-12-19T06:17:24', '{"marker": "I", "source": "secondary"}'),
('SYS-0224', '2025-12-19T06:23:24', '{"marker": "S", "source": "secondary"}'),
('SYS-0007', '2025-12-24T18:10:06', '{"marker": "O", "source": "secondary"}'),
('SYS-0035', '2025-12-23T15:55:34', '{"marker": "C", "source": "secondary"}'),
('SYS-0010', '2025-12-23T15:55:09', '{"marker": "L", "source": "secondary"}'),
('SYS-0037', '2025-12-23T15:55:36', '{"marker": "D", "source": "secondary"}'),
('SYS-0001', '2025-12-24T08:02:00', '{"marker": "X", "source": "primary"}'),
('SYS-0225', '2025-12-19T06:23:25', '{"marker": "Q", "source": "secondary"}'),
('SYS-0095', '2025-12-21T06:02:40', '{"marker": "D", "source": "secondary"}'),
('SYS-0021', '2025-12-23T15:55:20', '{"marker": "W", "source": "secondary"}'),
('SYS-0001', '2025-12-24T08:10:00', '{"marker": "A", "source": "primary"}'),
('SYS-0009', '2025-12-24T08:02:08', '{"marker": "U", "source": "primary"}'),
('SYS-0119', '2025-12-19T06:03:03', '{"marker": "E", "source": "secondary"}'),
('SYS-0142', '2025-12-19T06:03:26', '{"marker": "O", "source": "secondary"}'),
('SYS-0048', '2025-12-23T15:55:30', '{"marker": "I", "source": "secondary"}'),
('SYS-0228', '2025-12-19T06:23:28', '{"marker": "G", "source": "secondary"}'),
('SYS-0092', '2025-12-21T06:02:37', '{"marker": "K", "source": "secondary"}'),
('SYS-0008', '2025-12-24T08:02:07', '{"marker": "B", "source": "primary"}'),
('SYS-0111', '2025-12-19T06:02:55', '{"marker": "W", "source": "secondary"}'),
('SYS-0106', '2025-12-19T06:02:50', '{"marker": "L", "source": "secondary"}'),
('SYS-0082', '2025-12-21T06:02:27', '{"marker": "O", "source": "secondary"}'),
('SYS-0180', '2025-12-19T06:17:31', '{"marker": "N", "source": "secondary"}'),
('SYS-0103', '2025-12-19T06:02:47', '{"marker": "F", "source": "secondary"}'),
('SYS-0096', '2025-12-21T06:02:41', '{"marker": "A", "source": "secondary"}'),
('SYS-0003', '2025-12-23T06:02:02', '{"marker": "Z", "source": "primary"}'),
('SYS-0157', '2025-12-19T06:17:11', '{"marker": "D", "source": "secondary"}'),
('SYS-0010', '2025-12-23T06:02:09', '{"marker": "C", "source": "primary"}'),
('SYS-0102', '2025-12-19T06:02:46', '{"marker": "O", "source": "secondary"}'),
('SYS-0072', '2025-12-23T06:02:18', '{"marker": "K", "source": "secondary"}'),
('SYS-0175', '2025-12-19T06:17:29', '{"marker": "D", "source": "secondary"}'),
('SYS-0214', '2025-12-19T06:23:14', '{"marker": "U", "source": "secondary"}'),
('SYS-0005', '2025-12-24T08:02:04', '{"marker": "P", "source": "primary"}'),
('SYS-0216', '2025-12-19T06:23:16', '{"marker": "A", "source": "secondary"}'),
('SYS-0047', '2025-12-23T15:55:46', '{"marker": "H", "source": "secondary"}'),
('SYS-0019', '2025-12-23T15:55:18', '{"marker": "H", "source": "secondary"}'),
('SYS-0043', '2025-12-23T15:55:42', '{"marker": "Q", "source": "secondary"}'),
('SYS-0039', '2025-12-23T15:55:38', '{"marker": "L", "source": "secondary"}'),
('SYS-0212', '2025-12-19T06:23:12', '{"marker": "K", "source": "secondary"}'),
('SYS-0133', '2025-12-19T06:03:17', '{"marker": "C", "source": "secondary"}'),
('SYS-0005', '2025-12-24T08:10:04', '{"marker": "N", "source": "primary"}'),
('SYS-0011', '2025-12-24T08:02:10', '{"marker": "J", "source": "primary"}'),
('SYS-0080', '2025-12-21T06:02:25', '{"marker": "N", "source": "secondary"}'),
('SYS-0230', '2025-12-19T06:23:30', '{"marker": "I", "source": "secondary"}'),
('SYS-0100', '2025-12-21T06:02:45', '{"marker": "N", "source": "secondary"}'),
('SYS-0061', '2025-12-23T06:02:07', '{"marker": "T", "source": "secondary"}'),
('SYS-0022', '2025-12-23T15:55:21', '{"marker": "K", "source": "secondary"}'),
('SYS-0149', '2025-12-19T06:17:03', '{"marker": "H", "source": "secondary"}'),
('SYS-0009', '2025-12-23T15:55:08', '{"marker": "D", "source": "secondary"}'),
('SYS-0002', '2025-12-24T08:02:01', '{"marker": "Y", "source": "primary"}'),
('SYS-0132', '2025-12-19T06:03:16', '{"marker": "K", "source": "secondary"}'),
('SYS-0033', '2025-12-23T15:55:32', '{"marker": "W", "source": "secondary"}'),
('SYS-0193', '2025-12-19T06:17:44', '{"marker": "C", "source": "secondary"}'),
('SYS-0167', '2025-12-19T06:17:21', '{"marker": "A", "source": "secondary"}'),
('SYS-0008', '2025-12-24T08:10:07', '{"marker": "F", "source": "primary"}'),
('SYS-0186', '2025-12-19T06:17:37', '{"marker": "L", "source": "secondary"}'),
('SYS-0078', '2025-12-21T06:02:23', '{"marker": "V", "source": "secondary"}'),
('SYS-0085', '2025-12-21T06:02:30', '{"marker": "Q", "source": "secondary"}'),
('SYS-0046', '2025-12-23T15:55:45', '{"marker": "G", "source": "secondary"}'),
('SYS-0054', '2025-12-23T15:55:36', '{"marker": "W", "source": "secondary"}'),
('SYS-0235', '2025-12-19T06:23:35', '{"marker": "D", "source": "secondary"}'),
('SYS-0089', '2025-12-21T06:02:34', '{"marker": "H", "source": "secondary"}'),
('SYS-0011', '2025-12-23T06:02:10', '{"marker": "J", "source": "primary"}'),
('SYS-0034', '2025-12-23T15:55:33', '{"marker": "K", "source": "secondary"}'),
('SYS-0129', '2025-12-19T06:03:13', '{"marker": "H", "source": "secondary"}'),
('SYS-0211', '2025-12-19T06:23:11', '{"marker": "W", "source": "secondary"}'),
('SYS-0151', '2025-12-19T06:17:05', '{"marker": "W", "source": "secondary"}'),
('SYS-0204', '2025-12-19T06:23:04', '{"marker": "S", "source": "secondary"}'),
('SYS-0213', '2025-12-19T06:23:13', '{"marker": "C", "source": "secondary"}'),
('SYS-0147', '2025-12-19T06:03:31', '{"marker": "A", "source": "secondary"}'),
('SYS-0140', '2025-12-19T06:03:24', '{"marker": "N", "source": "secondary"}'),
('SYS-0123', '2025-12-19T06:03:07', '{"marker": "F", "source": "secondary"}'),
('SYS-0052', '2025-12-23T15:55:34', '{"marker": "H", "source": "secondary"}'),
('SYS-0148', '2025-12-19T06:17:02', '{"marker": "G", "source": "secondary"}'),
('SYS-0059', '2025-12-23T06:02:05', '{"marker": "E", "source": "secondary"}'),
('SYS-0027', '2025-12-23T15:55:26', '{"marker": "K", "source": "secondary"}'),
('SYS-0055', '2025-12-23T15:55:37', '{"marker": "K", "source": "secondary"}'),
('SYS-0155', '2025-12-19T06:17:09', '{"marker": "D", "source": "secondary"}'),
('SYS-0107', '2025-12-19T06:02:51', '{"marker": "A", "source": "secondary"}'),
('SYS-0173', '2025-12-19T06:17:27', '{"marker": "C", "source": "secondary"}'),
('SYS-0007', '2025-12-23T06:02:06', '{"marker": "K", "source": "primary"}'),
('SYS-0016', '2025-12-23T15:55:15', '{"marker": "L", "source": "secondary"}'),
('SYS-0168', '2025-12-19T06:17:22', '{"marker": "G", "source": "secondary"}'),
('SYS-0196', '2025-12-19T06:17:47', '{"marker": "A", "source": "secondary"}'),
('SYS-0234', '2025-12-19T06:23:34', '{"marker": "U", "source": "secondary"}'),
('SYS-0210', '2025-12-19T06:23:10', '{"marker": "I", "source": "secondary"}'),
('SYS-0135', '2025-12-19T06:03:19', '{"marker": "D", "source": "secondary"}'),
('SYS-0150', '2025-12-19T06:17:04', '{"marker": "I", "source": "secondary"}'),
('SYS-0091', '2025-12-21T06:02:36', '{"marker": "W", "source": "secondary"}'),
('SYS-0154', '2025-12-19T06:17:08', '{"marker": "U", "source": "secondary"}'),
('SYS-0223', '2025-12-19T06:23:23', '{"marker": "F", "source": "secondary"}'),
('SYS-0020', '2025-12-23T15:55:19', '{"marker": "I", "source": "secondary"}'),
('SYS-0160', '2025-12-19T06:17:14', '{"marker": "N", "source": "secondary"}'),
('SYS-0128', '2025-12-19T06:03:12', '{"marker": "G", "source": "secondary"}'),
('SYS-0194', '2025-12-19T06:17:45', '{"marker": "U", "source": "secondary"}'),
('SYS-0231', '2025-12-19T06:23:31', '{"marker": "W", "source": "secondary"}'),
('SYS-0045', '2025-12-23T15:55:44', '{"marker": "A", "source": "secondary"}'),
('SYS-0190', '2025-12-19T06:17:41', '{"marker": "I", "source": "secondary"}'),
('SYS-0197', '2025-12-19T06:17:48', '{"marker": "D", "source": "secondary"}'),
('SYS-0002', '2025-12-23T06:02:01', '{"marker": "Y", "source": "primary"}'),
('SYS-0008', '2025-12-24T18:10:07', '{"marker": "F", "source": "secondary"}'),
('SYS-0158', '2025-12-19T06:17:12', '{"marker": "V", "source": "secondary"}'),
('SYS-0003', '2025-12-24T08:10:02', '{"marker": "V", "source": "primary"}'),
('SYS-0171', '2025-12-19T06:17:25', '{"marker": "W", "source": "secondary"}'),
('SYS-0229', '2025-12-19T06:23:29', '{"marker": "H", "source": "secondary"}'),
('SYS-0079', '2025-12-21T06:02:24', '{"marker": "E", "source": "secondary"}'),
('SYS-0215', '2025-12-19T06:23:15', '{"marker": "D", "source": "secondary"}'),
('SYS-0233', '2025-12-19T06:23:33', '{"marker": "C", "source": "secondary"}'),
('SYS-0060', '2025-12-23T06:02:06', '{"marker": "N", "source": "secondary"}'),
('SYS-0026', '2025-12-23T15:55:25', '{"marker": "W", "source": "secondary"}'),
('SYS-0125', '2025-12-19T06:03:09', '{"marker": "Q", "source": "secondary"}'),
('SYS-0001', '2025-12-24T18:10:00', '{"marker": "A", "source": "secondary"}'),
('SYS-0136', '2025-12-19T06:03:20', '{"marker": "A", "source": "secondary"}'),
('SYS-0010', '2025-12-24T08:02:09', '{"marker": "C", "source": "primary"}'),
('SYS-0010', '2025-12-24T08:10:09', '{"marker": "Q", "source": "primary"}'),
('SYS-0044', '2025-12-23T15:55:43', '{"marker": "L", "source": "secondary"}'),
('SYS-0115', '2025-12-19T06:02:59', '{"marker": "D", "source": "secondary"}'),
('SYS-0005', '2025-12-24T18:10:15', '{"marker": "W", "source": "secondary"}'),
('SYS-0094', '2025-12-21T06:02:39', '{"marker": "U", "source": "secondary"}'),
('SYS-0191', '2025-12-19T06:17:42', '{"marker": "W", "source": "secondary"}'),
('SYS-0182', '2025-12-19T06:17:33', '{"marker": "O", "source": "secondary"}'),
('SYS-0195', '2025-12-19T06:17:46', '{"marker": "D", "source": "secondary"}'),
('SYS-0227', '2025-12-19T06:23:27', '{"marker": "A", "source": "secondary"}'),
('SYS-0032', '2025-12-23T15:55:31', '{"marker": "I", "source": "secondary"}'),
('SYS-0177', '2025-12-19T06:17:31', '{"marker": "D", "source": "secondary"}'),
('SYS-0068', '2025-12-23T06:02:14', '{"marker": "G", "source": "secondary"}'),
('SYS-0069', '2025-12-23T06:02:15', '{"marker": "H", "source": "secondary"}'),
('SYS-0009', '2025-12-23T06:02:08', '{"marker": "U", "source": "primary"}'),
('SYS-0176', '2025-12-19T06:17:30', '{"marker": "A", "source": "secondary"}'),
('SYS-0008', '2025-12-23T06:02:07', '{"marker": "B", "source": "primary"}'),
('SYS-0099', '2025-12-21T06:02:44', '{"marker": "E", "source": "secondary"}'),
('SYS-0014', '2025-12-23T15:55:13', '{"marker": "S", "source": "secondary"}'),
('SYS-0042', '2025-12-23T15:55:41', '{"marker": "S", "source": "secondary"}'),
('SYS-0006', '2025-12-24T08:02:05', '{"marker": "R", "source": "primary"}'),
('SYS-0187', '2025-12-19T06:17:38', '{"marker": "A", "source": "secondary"}'),
('SYS-0174', '2025-12-19T06:17:28', '{"marker": "U", "source": "secondary"}'),
('SYS-0056', '2025-12-23T15:55:55', '{"marker": "A", "source": "secondary"}'),
('SYS-0199', '2025-12-19T06:17:50', '{"marker": "E", "source": "secondary"}'),
('SYS-0120', '2025-12-19T06:03:04', '{"marker": "N", "source": "secondary"}'),
('SYS-0003', '2025-12-24T08:02:02', '{"marker": "Z", "source": "primary"}'),
('SYS-0116', '2025-12-19T06:02:59', '{"marker": "A", "source": "secondary"}'),
('SYS-0008', '2025-12-24T18:10:18', '{"marker": "U", "source": "secondary"}'),
('SYS-0205', '2025-12-19T06:23:05', '{"marker": "Q", "source": "secondary"}'),
('SYS-0006', '2025-12-24T08:10:05', '{"marker": "T", "source": "primary"}'),
('SYS-0145', '2025-12-19T06:03:29', '{"marker": "Q", "source": "secondary"}'),
('SYS-0004', '2025-12-24T08:10:03', '{"marker": "E", "source": "primary"}'),
('SYS-0065', '2025-12-23T06:02:11', '{"marker": "Q", "source": "secondary"}'),
('SYS-0086', '2025-12-21T06:02:31', '{"marker": "L", "source": "secondary"}'),
('SYS-0188', '2025-12-19T06:17:39', '{"marker": "G", "source": "secondary"}'),
('SYS-0183', '2025-12-19T06:17:34', '{"marker": "F", "source": "secondary"}'),
('SYS-0062', '2025-12-23T06:02:08', '{"marker": "O", "source": "secondary"}'),
('SYS-0058', '2025-12-23T15:55:57', '{"marker": "V", "source": "secondary"}'),
('SYS-0169', '2025-12-19T06:17:23', '{"marker": "H", "source": "secondary"}'),
('SYS-0105', '2025-12-19T06:02:49', '{"marker": "Q", "source": "secondary"}'),
('SYS-0001', '2025-12-24T18:10:11', '{"marker": "A", "source": "secondary"}'),
('SYS-0005', '2025-12-23T06:02:04', '{"marker": "P", "source": "primary"}'),
('SYS-0131', '2025-12-19T06:03:15', '{"marker": "W", "source": "secondary"}'),
('SYS-0232', '2025-12-19T06:23:32', '{"marker": "K", "source": "secondary"}'),
('SYS-0028', '2025-12-23T15:55:27', '{"marker": "L", "source": "secondary"}'),
('SYS-0070', '2025-12-23T06:02:16', '{"marker": "I", "source": "secondary"}'),
('SYS-0029', '2025-12-23T15:55:28', '{"marker": "A", "source": "secondary"}'),
('SYS-0127', '2025-12-19T06:03:11', '{"marker": "A", "source": "secondary"}'),
('SYS-0010', '2025-12-24T18:10:09', '{"marker": "Q", "source": "secondary"}'),
('SYS-0025', '2025-12-23T15:55:24', '{"marker": "I", "source": "secondary"}'),
('SYS-0164', '2025-12-19T06:17:18', '{"marker": "S", "source": "secondary"}'),
('SYS-0114', '2025-12-19T06:02:58', '{"marker": "U", "source": "secondary"}'),
('SYS-0009', '2025-12-24T18:10:08', '{"marker": "S", "source": "secondary"}'),
('SYS-0184', '2025-12-19T06:17:35', '{"marker": "S", "source": "secondary"}'),
('SYS-0003', '2025-12-24T18:10:02', '{"marker": "V", "source": "secondary"}'),
('SYS-0053', '2025-12-23T15:55:35', '{"marker": "I", "source": "secondary"}'),
('SYS-0117', '2025-12-19T06:03:01', '{"marker": "D", "source": "secondary"}'),
('SYS-0074', '2025-12-23T06:02:20', '{"marker": "U", "source": "secondary"}'),
('SYS-0198', '2025-12-19T06:17:49', '{"marker": "V", "source": "secondary"}'),
('SYS-0036', '2025-12-23T15:55:35', '{"marker": "U", "source": "secondary"}'),
('SYS-0112', '2025-12-19T06:02:56', '{"marker": "K", "source": "secondary"}'),
('SYS-0023', '2025-12-23T15:55:22', '{"marker": "C", "source": "secondary"}'),
('SYS-0121', '2025-12-19T06:03:05', '{"marker": "T", "source": "secondary"}'),
('SYS-0181', '2025-12-19T06:17:32', '{"marker": "T", "source": "secondary"}'),
('SYS-0003', '2025-12-24T18:10:13', '{"marker": "H", "source": "secondary"}'),
('SYS-0130', '2025-12-19T06:03:14', '{"marker": "I", "source": "secondary"}'),
('SYS-0066', '2025-12-23T06:02:12', '{"marker": "L", "source": "secondary"}'),
('SYS-0203', '2025-12-19T06:17:54', '{"marker": "F", "source": "secondary"}'),
('SYS-0002', '2025-12-24T18:10:01', '{"marker": "D", "source": "secondary"}'),
('SYS-0104', '2025-12-19T06:02:48', '{"marker": "S", "source": "secondary"}'),
('SYS-0017', '2025-12-23T15:55:16', '{"marker": "A", "source": "secondary"}'),
('SYS-0004', '2025-12-24T18:10:03', '{"marker": "E", "source": "secondary"}'),
('SYS-0238', '2025-12-19T06:23:38', '{"marker": "D", "source": "secondary"}'),
('SYS-0110', '2025-12-19T06:02:54', '{"marker": "I", "source": "secondary"}'),
('SYS-0071', '2025-12-23T06:02:17', '{"marker": "W", "source": "secondary"}'),
('SYS-0004', '2025-12-24T18:10:14', '{"marker": "I", "source": "secondary"}'),
('SYS-0031', '2025-12-23T15:55:30', '{"marker": "H", "source": "secondary"}'),
('SYS-0004', '2025-12-23T06:02:03', '{"marker": "M", "source": "primary"}'),
('SYS-0141', '2025-12-19T06:03:25', '{"marker": "T", "source": "secondary"}'),
('SYS-0207', '2025-12-19T06:23:07', '{"marker": "A", "source": "secondary"}'),
('SYS-0122', '2025-12-19T06:03:06', '{"marker": "O", "source": "secondary"}'),
('SYS-0006', '2025-12-24T18:10:05', '{"marker": "T", "source": "secondary"}'),
('SYS-0030', '2025-12-23T15:55:29', '{"marker": "G", "source": "secondary"}'),
('SYS-0192', '2025-12-19T06:17:43', '{"marker": "K", "source": "secondary"}'),
('SYS-0144', '2025-12-19T06:03:28', '{"marker": "S", "source": "secondary"}'),
('SYS-0098', '2025-12-21T06:02:43', '{"marker": "V", "source": "secondary"}'),
('SYS-0200', '2025-12-19T06:17:51', '{"marker": "N", "source": "secondary"}'),
('SYS-0007', '2025-12-24T08:02:06', '{"marker": "K", "source": "primary"}'),
('SYS-0220', '2025-12-19T06:23:20', '{"marker": "N", "source": "secondary"}'),
('SYS-0063', '2025-12-23T06:02:09', '{"marker": "F", "source": "secondary"}'),
('SYS-0109', '2025-12-19T06:02:53', '{"marker": "H", "source": "secondary"}'),
('SYS-0179', '2025-12-19T06:17:30', '{"marker": "E", "source": "secondary"}'),
('SYS-0163', '2025-12-19T06:17:17', '{"marker": "F", "source": "secondary"}'),
('SYS-0162', '2025-12-19T06:17:16', '{"marker": "O", "source": "secondary"}'),
('SYS-0222', '2025-12-19T06:23:22', '{"marker": "O", "source": "secondary"}'),
('SYS-0134', '2025-12-19T06:03:18', '{"marker": "U", "source": "secondary"}'),
('SYS-0004', '2025-12-24T08:02:03', '{"marker": "M", "source": "primary"}'),
('SYS-0139', '2025-12-19T06:03:23', '{"marker": "E", "source": "secondary"}'),
('SYS-0097', '2025-12-21T06:02:42', '{"marker": "D", "source": "secondary"}'),
('SYS-0146', '2025-12-19T06:03:30', '{"marker": "L", "source": "secondary"}'),
('SYS-0101', '2025-12-19T06:02:45', '{"marker": "T", "source": "secondary"}'),
('SYS-0201', '2025-12-19T06:17:52', '{"marker": "T", "source": "secondary"}'),
('SYS-0051', '2025-12-23T15:55:33', '{"marker": "C", "source": "secondary"}'),
('SYS-0172', '2025-12-19T06:17:26', '{"marker": "K", "source": "secondary"}'),
('SYS-0083', '2025-12-21T06:02:28', '{"marker": "F", "source": "secondary"}'),
('SYS-0009', '2025-12-24T08:10:08', '{"marker": "S", "source": "primary"}'),
('SYS-0124', '2025-12-19T06:03:08', '{"marker": "S", "source": "secondary"}'),
('SYS-0041', '2025-12-23T15:55:40', '{"marker": "F", "source": "secondary"}'),
('SYS-0049', '2025-12-23T15:55:31', '{"marker": "W", "source": "secondary"}'),
('SYS-0202', '2025-12-19T06:17:53', '{"marker": "O", "source": "secondary"}'),
('SYS-0178', '2025-12-19T06:17:29', '{"marker": "V", "source": "secondary"}'),
('SYS-0090', '2025-12-21T06:02:35', '{"marker": "I", "source": "secondary"}'),
('SYS-0006', '2025-12-23T06:02:05', '{"marker": "R", "source": "primary"}'),
('SYS-0073', '2025-12-23T06:02:19', '{"marker": "C", "source": "secondary"}'),
('SYS-0076', '2025-12-23T06:02:22', '{"marker": "A", "source": "secondary"}'),
('SYS-0217', '2025-12-19T06:23:17', '{"marker": "D", "source": "secondary"}'),
('SYS-0159', '2025-12-19T06:17:13', '{"marker": "E", "source": "secondary"}'),
('SYS-0161', '2025-12-19T06:17:15', '{"marker": "T", "source": "secondary"}'),
('SYS-0221', '2025-12-19T06:23:21', '{"marker": "T", "source": "secondary"}'),
('SYS-0236', '2025-12-19T06:23:36', '{"marker": "U", "source": "secondary"}'),
('SYS-0064', '2025-12-23T06:02:10', '{"marker": "S", "source": "secondary"}'),
('SYS-0126', '2025-12-19T06:03:10', '{"marker": "L", "source": "secondary"}'),
('SYS-0024', '2025-12-23T15:55:23', '{"marker": "H", "source": "secondary"}'),
('SYS-0113', '2025-12-19T06:02:57', '{"marker": "C", "source": "secondary"}'),
('SYS-0038', '2025-12-23T15:55:37', '{"marker": "L", "source": "secondary"}'),
('SYS-0075', '2025-12-23T06:02:21', '{"marker": "D", "source": "secondary"}'),
('SYS-0088', '2025-12-21T06:02:33', '{"marker": "G", "source": "secondary"}'),
('SYS-0007', '2025-12-24T08:10:06', '{"marker": "O", "source": "primary"}'),
('SYS-0067', '2025-12-23T06:02:13', '{"marker": "A", "source": "secondary"}'),
('SYS-0156', '2025-12-19T06:17:10', '{"marker": "A", "source": "secondary"}'),
('SYS-0040', '2025-12-23T15:55:39', '{"marker": "O", "source": "secondary"}'),
('SYS-0219', '2025-12-19T06:23:19', '{"marker": "E", "source": "secondary"}'),
('SYS-0166', '2025-12-19T06:17:20', '{"marker": "L", "source": "secondary"}'),
('SYS-0007', '2025-12-24T18:10:17', '{"marker": "C", "source": "secondary"}'),
('SYS-0153', '2025-12-19T06:17:07', '{"marker": "C", "source": "secondary"}'),
('SYS-0077', '2025-12-23T06:02:23', '{"marker": "D", "source": "secondary"}'),
('SYS-0084', '2025-12-21T06:02:29', '{"marker": "S", "source": "secondary"}');

INSERT INTO incoming_dispatches (system_id, dispatched_at, payload) VALUES
('SYS-0013', '2025-12-23T15:55:12', '{"marker": "F", "source": "secondary"}'),
('SYS-0002', '2025-12-24T08:10:01', '{"marker": "D", "source": "primary"}'),
('SYS-0012', '2025-12-23T15:55:11', '{"marker": "O", "source": "secondary"}'),
('SYS-0108', '2025-12-19T06:02:52', '{"marker": "G", "source": "secondary"}'),
('SYS-0165', '2025-12-19T06:17:19', '{"marker": "Q", "source": "secondary"}'),
('SYS-0189', '2025-12-19T06:17:40', '{"marker": "H", "source": "secondary"}'),
('SYS-0050', '2025-12-23T15:55:32', '{"marker": "K", "source": "secondary"}'),
('SYS-0138', '2025-12-19T06:03:22', '{"marker": "V", "source": "secondary"}'),
('SYS-0087', '2025-12-21T06:02:32', '{"marker": "A", "source": "secondary"}'),
('SYS-0208', '2025-12-19T06:23:08', '{"marker": "G", "source": "secondary"}'),
('SYS-0057', '2025-12-23T15:55:56', '{"marker": "D", "source": "secondary"}'),
('SYS-0002', '2025-12-24T18:10:12', '{"marker": "G", "source": "secondary"}'),
('SYS-0011', '2025-12-24T08:10:10', '{"marker": "L", "source": "primary"}'),
('SYS-0226', '2025-12-19T06:23:26', '{"marker": "L", "source": "secondary"}'),
('SYS-0206', '2025-12-19T06:23:06', '{"marker": "L", "source": "secondary"}'),
('SYS-0118', '2025-12-19T06:03:02', '{"marker": "V", "source": "secondary"}'),
('SYS-0185', '2025-12-19T06:17:36', '{"marker": "Q", "source": "secondary"}'),
('SYS-0218', '2025-12-19T06:23:18', '{"marker": "V", "source": "secondary"}'),
('SYS-0018', '2025-12-23T15:55:17', '{"marker": "G", "source": "secondary"}'),
('SYS-0005', '2025-12-24T18:10:04', '{"marker": "N", "source": "secondary"}'),
('SYS-0093', '2025-12-21T06:02:38', '{"marker": "C", "source": "secondary"}'),
('SYS-0011', '2025-12-24T18:10:10', '{"marker": "L", "source": "secondary"}'),
('SYS-0011', '2025-12-23T15:55:10', '{"marker": "L", "source": "secondary"}'),
('SYS-0143', '2025-12-19T06:03:27', '{"marker": "F", "source": "secondary"}'),
('SYS-0152', '2025-12-19T06:17:06', '{"marker": "K", "source": "secondary"}'),
('SYS-0015', '2025-12-23T15:55:14', '{"marker": "Q", "source": "secondary"}'),
('SYS-0001', '2025-12-23T06:02:00', '{"marker": "X", "source": "primary"}'),
('SYS-0209', '2025-12-19T06:23:09', '{"marker": "H", "source": "secondary"}'),
('SYS-0133', '2025-12-19T06:03:17', '{"marker": "C", "source": "secondary"}'),
('SYS-0158', '2025-12-19T06:17:12', '{"marker": "V", "source": "secondary"}'),
('SYS-0195', '2025-12-19T06:17:46', '{"marker": "D", "source": "secondary"}'),
('SYS-0182', '2025-12-19T06:17:33', '{"marker": "O", "source": "secondary"}'),
('SYS-0033', '2025-12-23T15:55:32', '{"marker": "W", "source": "secondary"}'),
('SYS-0083', '2025-12-21T06:02:28', '{"marker": "F", "source": "secondary"}'),
('SYS-0002', '2025-12-23T06:02:01', '{"marker": "Y", "source": "primary"}'),
('SYS-0076', '2025-12-23T06:02:22', '{"marker": "A", "source": "secondary"}'),
('SYS-0009', '2025-12-24T08:02:08', '{"marker": "U", "source": "primary"}'),
('SYS-0194', '2025-12-19T06:17:45', '{"marker": "U", "source": "secondary"}');

```


```
$ sqlite3
SQLite version 3.50.4 2025-07-30 19:33:53
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite> .read day15-inserts.sql
Parse error near line 19: unrecognized token: ":"
  6:02:26', '{"marker": "T", "source": "secondary"}'::jsonb), ('SYS-0137', '2025
                                      error here ---^
Parse error near line 275: unrecognized token: ":"
  5:55:12', '{"marker": "F", "source": "secondary"}'::jsonb), ('SYS-0002', '2025
                                      error here ---^
sqlite> .read day15-inserts-sqlite.sql
sqlite> .schema
CREATE TABLE system_dispatches (
    id SERIAL PRIMARY KEY,
    system_id TEXT NOT NULL,
    dispatched_at TIMESTAMP NOT NULL,
    payload JSONB NOT NULL,
    marker_letter TEXT GENERATED ALWAYS AS (json_extract(payload, '$.marker')) STORED,
    UNIQUE (system_id, dispatched_at, payload)
);
CREATE TABLE incoming_dispatches (
    system_id TEXT,
    dispatched_at TIMESTAMP,
    payload JSONB
);
sqlite> .mode table
sqlite> SELECT * FROM system_dispatches LIMIT 10;
+----+-----------+---------------------+----------------------------------------+---------------+
| id | system_id |    dispatched_at    |                payload                 | marker_letter |
+----+-----------+---------------------+----------------------------------------+---------------+
|    | SYS-0081  | 2025-12-21T06:02:26 | {"marker": "T", "source": "secondary"} | T             |
|    | SYS-0137  | 2025-12-19T06:03:21 | {"marker": "D", "source": "secondary"} | D             |
|    | SYS-0237  | 2025-12-19T06:23:37 | {"marker": "D", "source": "secondary"} | D             |
|    | SYS-0006  | 2025-12-24T18:10:16 | {"marker": "K", "source": "secondary"} | K             |
|    | SYS-0170  | 2025-12-19T06:17:24 | {"marker": "I", "source": "secondary"} | I             |
|    | SYS-0224  | 2025-12-19T06:23:24 | {"marker": "S", "source": "secondary"} | S             |
|    | SYS-0007  | 2025-12-24T18:10:06 | {"marker": "O", "source": "secondary"} | O             |
|    | SYS-0035  | 2025-12-23T15:55:34 | {"marker": "C", "source": "secondary"} | C             |
|    | SYS-0010  | 2025-12-23T15:55:09 | {"marker": "L", "source": "secondary"} | L             |
|    | SYS-0037  | 2025-12-23T15:55:36 | {"marker": "D", "source": "secondary"} | D             |
+----+-----------+---------------------+----------------------------------------+---------------+
sqlite> SELECT * FROM incoming_dispatches LIMIT 10;
+-----------+---------------------+----------------------------------------+
| system_id |    dispatched_at    |                payload                 |
+-----------+---------------------+----------------------------------------+
| SYS-0013  | 2025-12-23T15:55:12 | {"marker": "F", "source": "secondary"} |
| SYS-0002  | 2025-12-24T08:10:01 | {"marker": "D", "source": "primary"}   |
| SYS-0012  | 2025-12-23T15:55:11 | {"marker": "O", "source": "secondary"} |
| SYS-0108  | 2025-12-19T06:02:52 | {"marker": "G", "source": "secondary"} |
| SYS-0165  | 2025-12-19T06:17:19 | {"marker": "Q", "source": "secondary"} |
| SYS-0189  | 2025-12-19T06:17:40 | {"marker": "H", "source": "secondary"} |
| SYS-0050  | 2025-12-23T15:55:32 | {"marker": "K", "source": "secondary"} |
| SYS-0138  | 2025-12-19T06:03:22 | {"marker": "V", "source": "secondary"} |
| SYS-0087  | 2025-12-21T06:02:32 | {"marker": "A", "source": "secondary"} |
| SYS-0208  | 2025-12-19T06:23:08 | {"marker": "G", "source": "secondary"} |
+-----------+---------------------+----------------------------------------+
sqlite>
```

Good to go!

## Problem

> Reconstruct the final confirmation phrase to text Santa based on the elves’ hazy recollection of how they solved this problem before.
> 
> Your final result should include the marker_letter for each system, using only the most recent dispatch from a primary source. Once the correct dispatch has been identified for every system, combine the results and order them by dispatched_at in ascending order to reveal the confirmation phrase.
> 
> The sleigh won’t launch without it.


So, we need to first make sure we have the right `dispatches`.

We have two tables.

1. `system_dispatches`
2. `incoming_dispatches`

The problem statement also states:

> The system kept throwing errors until we figured out how to handle duplicates. Whatever you do the records already in system_dispatches must take precedence.”


So, we need to take the `system_dispatches` as the ground truth. However we need to include something or the other from the `incoming_dispatches` as there are new entries being created.

So, we will first try to insert into the system dispatches everything in the `incoming_dispatches`

```sql
INSERT INTO system_dispatches (system_id, dispatched_at, payload)
SELECT system_id, dispatched_at, payload
FROM incoming_dispatches;
```

```
sqlite> INSERT INTO system_dispatches (system_id, dispatched_at, payload)
SELECT system_id, dispatched_at, payload
FROM incoming_dispatches;
Runtime error: UNIQUE constraint failed: system_dispatches.system_id, system_dispatches.dispatched_at, system_dispatches.payload (19)
sqlite>
```

Ops!

```sql

```
```
sqlite> SELECT COUNT(*) FROM system_dispatches ;
+----------+
| COUNT(*) |
+----------+
| 254      |
+----------+
sqlite> SELECT COUNT(*) FROM incoming_dispatches;
+----------+
| COUNT(*) |
+----------+
| 38       |
+----------+

sqlite> SELECT COUNT(*) FROM (SELECT system_id, dispatched_at, payload FROM system_dispatches UNION SELECT * FROM incoming_dispatches);
+----------+
| COUNT(*) |
+----------+
| 282      |
+----------+

sqlite> SELECT DISTINCT COUNT(*) FROM (SELECT system_id, dispatched_at, payload FROM system_dispatches UNION SELECT * FROM incoming_dispatches);
+----------+
| COUNT(*) |
+----------+
| 282      |
+----------+

sqlite> SELECT COUNT(*) FROM (SELECT system_id, dispatched_at, payload FROM system_dispatches INTERSECT SELECT * FROM incoming_dispatches);
+----------+
| COUNT(*) |
+----------+
| 10       |
+----------+
sqlite>


We need to be careful about what we insert into the `system_dispatches`.

I think the duplicates causes this conflict in constraints, we can ignore on the constraint failures and don't insert that row.

```sql
INSERT OR IGNORE INTO system_dispatches (system_id, dispatched_at, payload)
SELECT system_id, dispatched_at, payload
FROM incoming_dispatches;
```

Ok, now if we see the count, we get `282` which is making sense.

- The original `system_dispatches` had a count of `254` before merge.
- The original `incoming_dispatches` had a count of `38`.
- After merging `incoming_dispatches` into `system_dispatches` table, we have a `282` records in `system_dispatches` but `254+38=292`, what about the `10`? Those were already there in the `system_dispatches`. As we saw in the earlier query, there were `10` rows in common from the `incoming_dispatches` table before merge.

```
sqlite> SELECT COUNT(*) FROM system_dispatches ;
+----------+
| COUNT(*) |
+----------+
| 254      |
+----------+
sqlite> SELECT COUNT(*) FROM incoming_dispatches;
+----------+
| COUNT(*) |
+----------+
| 38       |
+----------+
sqlite> INSERT INTO system_dispatches (system_id, dispatched_at, payload)
SELECT system_id, dispatched_at, payload
FROM incoming_dispatches;
Runtime error: UNIQUE constraint failed: system_dispatches.system_id, system_dispatches.dispatched_at, system_dispatches.payload (19)
sqlite> INSERT OR IGNORE INTO system_dispatches (system_id, dispatched_at, payload)
SELECT system_id, dispatched_at, payload
FROM incoming_dispatches;
sqlite> SELECT COUNT(*) FROM system_dispatches ;
+----------+
| COUNT(*) |
+----------+
| 282      |
+----------+
sqlite>
```

Now, to the actual problem.

What we need to do now?

- Using only the most recent dispatch from a primary source. 
- Once the correct dispatch has been identified for every system, combine the results and order them by dispatched_at in ascending order 
- Reveal the confirmation phrase

We have the single source now!

```sql
SELECT * FROM system_dispatches LIMIT 10;
```

```
SELECT * FROM system_dispatches LIMIT 10;
+----+-----------+---------------------+----------------------------------------+---------------+
| id | system_id |    dispatched_at    |                payload                 | marker_letter |
+----+-----------+---------------------+----------------------------------------+---------------+
|    | SYS-0081  | 2025-12-21T06:02:26 | {"marker": "T", "source": "secondary"} | T             |
|    | SYS-0137  | 2025-12-19T06:03:21 | {"marker": "D", "source": "secondary"} | D             |
|    | SYS-0237  | 2025-12-19T06:23:37 | {"marker": "D", "source": "secondary"} | D             |
|    | SYS-0006  | 2025-12-24T18:10:16 | {"marker": "K", "source": "secondary"} | K             |
|    | SYS-0170  | 2025-12-19T06:17:24 | {"marker": "I", "source": "secondary"} | I             |
|    | SYS-0224  | 2025-12-19T06:23:24 | {"marker": "S", "source": "secondary"} | S             |
|    | SYS-0007  | 2025-12-24T18:10:06 | {"marker": "O", "source": "secondary"} | O             |
|    | SYS-0035  | 2025-12-23T15:55:34 | {"marker": "C", "source": "secondary"} | C             |
|    | SYS-0010  | 2025-12-23T15:55:09 | {"marker": "L", "source": "secondary"} | L             |
|    | SYS-0037  | 2025-12-23T15:55:36 | {"marker": "D", "source": "secondary"} | D             |
+----+-----------+---------------------+----------------------------------------+---------------+
sqlite>
```

We have 4 columns.

- system_id
- dispatched_at
- payload (json string)
- marker (its a generated table from the payload["marker"] string)

So, we need to group by system, that's where the `system_id` is for.

Then we need to order by `dispatched_at` with the latest ones.

Finally we also need to filter the records that has `payload["source"]` as `primary`.

But wait, if we group by `system_id` then how can we get the latest dispatched ordered correctly? We can't use it in the `ORDER BY` as the group would have already been created right?

Oh! We might need window functions, but just a moment, can we do it without them?

Let's try.

```sql
SELECT 
    system_id,
    MAX(dispatched_at),
    payload, 
    marker_letter
FROM system_dispatches
WHERE json_extract(payload, '$.source') = 'primary'
GROUP BY system_id
ORDER BY dispatched_at ASC;
```

Oh boy, that was simple!

We just grouped by `system_id`, filtered the source as `primary` and selected the row with `MAX(dispatched_at)` which will give us the latest dispatch record for a system. Boom!

The confirmation phrase is `ADVENTOFSQL`

We can just select the `marker_letter`

```sql
SELECT
    marker_letter
FROM system_dispatches
WHERE json_extract(payload, '$.source') = 'primary'
GROUP BY system_id
ORDER BY dispatched_at ASC;
```


```
sqlite> SELECT
    marker_letter
FROM system_dispatches
WHERE json_extract(payload, '$.source') = 'primary'
GROUP BY system_id
ORDER BY dispatched_at ASC;
+---------------+
| marker_letter |
+---------------+
| X             |
| Y             |
| Z             |
| M             |
| P             |
| R             |
| K             |
| B             |
| U             |
| C             |
| J             |
+---------------+
sqlite>
```

Whops! What happened, the marker letter changed?

Yep, because we have ordered by `dispatched_at DESC` but never told how to return the rows, by using `MAX(dispatched_at)` we were selecting the latest date record. Without it, we are selecting the first record in the group that could be the first or the oldest dispatched_at time.

So, we need `MAX(dispatched_at)` included in the selected result set.


```sql
SELECT
    marker_letter, MAX(dispatched_at)
FROM system_dispatches
WHERE json_extract(payload, '$.source') = 'primary'
GROUP BY system_id
ORDER BY dispatched_at ASC;
```

```
 SELECT
    marker_letter, MAX(dispatched_at)
FROM system_dispatches
WHERE json_extract(payload, '$.source') = 'primary'
GROUP BY system_id
ORDER BY dispatched_at ASC;
+---------------+---------------------+
| marker_letter | MAX(dispatched_at)  |
+---------------+---------------------+
| A             | 2025-12-24T08:10:00 |
| D             | 2025-12-24T08:10:01 |
| V             | 2025-12-24T08:10:02 |
| E             | 2025-12-24T08:10:03 |
| N             | 2025-12-24T08:10:04 |
| T             | 2025-12-24T08:10:05 |
| O             | 2025-12-24T08:10:06 |
| F             | 2025-12-24T08:10:07 |
| S             | 2025-12-24T08:10:08 |
| Q             | 2025-12-24T08:10:09 |
| L             | 2025-12-24T08:10:10 |
+---------------+---------------------+
sqlite>
```

There, we go! It scared me for a moment.

We order by, only keep the `source` as `primary` and then select the latest `dispatched_at` for a given system_id.

If you don't want to do this way! I found a few more hacks

### With Subquery

We can SELECT the `MAX(dispatched_at)` in a subquery for that system_id and filter based on the `primary` source as usual.

```sql
SELECT 
    *
FROM system_dispatches
WHERE 
    json_extract(system_dispatches.payload, '$.source') = 'primary'
    AND system_dispatches.dispatched_at = (
        SELECT 
            MAX(latest_dispatches.dispatched_at)
        FROM system_dispatches latest_dispatches
        WHERE 
            latest_dispatches.system_id = system_dispatches.system_id
            AND json_extract(latest_dispatches.payload, '$.source') = 'primary'
   )
ORDER BY system_dispatches.dispatched_at ASC;
```

```
sqlite> SELECT 
    *
FROM system_dispatches
WHERE 
    json_extract(system_dispatches.payload, '$.source') = 'primary'
    AND system_dispatches.dispatched_at = (
        SELECT 
            MAX(latest_dispatches.dispatched_at)
        FROM system_dispatches latest_dispatches
        WHERE 
            latest_dispatches.system_id = system_dispatches.system_id
            AND json_extract(latest_dispatches.payload, '$.source') = 'primary'
   )
ORDER BY system_dispatches.dispatched_at ASC;

+----+-----------+---------------------+--------------------------------------+---------------+
| id | system_id |    dispatched_at    |               payload                | marker_letter |
+----+-----------+---------------------+--------------------------------------+---------------+
|    | SYS-0001  | 2025-12-24T08:10:00 | {"marker": "A", "source": "primary"} | A             |
|    | SYS-0002  | 2025-12-24T08:10:01 | {"marker": "D", "source": "primary"} | D             |
|    | SYS-0003  | 2025-12-24T08:10:02 | {"marker": "V", "source": "primary"} | V             |
|    | SYS-0004  | 2025-12-24T08:10:03 | {"marker": "E", "source": "primary"} | E             |
|    | SYS-0005  | 2025-12-24T08:10:04 | {"marker": "N", "source": "primary"} | N             |
|    | SYS-0006  | 2025-12-24T08:10:05 | {"marker": "T", "source": "primary"} | T             |
|    | SYS-0007  | 2025-12-24T08:10:06 | {"marker": "O", "source": "primary"} | O             |
|    | SYS-0008  | 2025-12-24T08:10:07 | {"marker": "F", "source": "primary"} | F             |
|    | SYS-0009  | 2025-12-24T08:10:08 | {"marker": "S", "source": "primary"} | S             |
|    | SYS-0010  | 2025-12-24T08:10:09 | {"marker": "Q", "source": "primary"} | Q             |
|    | SYS-0011  | 2025-12-24T08:10:10 | {"marker": "L", "source": "primary"} | L             |
+----+-----------+---------------------+--------------------------------------+---------------+
sqlite>
```



### Window Funciton

We can even use the window function to partition by the `system_id` and there itself order by the `dispatched_at` latest time. And then select it as a CTE table.

We can use a `ROW_NUMBER` function to assign a rank to each row per system_id and ordered by the latest `dispatched_at`

```sql
WITH latest_dispatches AS (
    SELECT
        *,
        ROW_NUMBER() OVER (
            PARTITION BY system_id
            ORDER BY dispatched_at DESC
        ) as rank
    FROM system_dispatches
    WHERE json_extract(payload, '$.source') = 'primary'
) SELECT * FROM latest_dispatches WHERE rank = 1 ORDER BY dispatched_at ASC;
```

The rest is the same as the above query. We order by dispatched latest but then order by ascending (oldest first) in the final result set.

```
sqlite> WITH latest_dispatches AS (
    SELECT
        *,
        ROW_NUMBER() OVER (
            PARTITION BY system_id
            ORDER BY dispatched_at DESC
        ) as rank
    FROM system_dispatches
    WHERE json_extract(payload, '$.source') = 'primary'
) SELECT * FROM latest_dispatches WHERE rank = 1 ORDER BY dispatched_at ASC;
+----+-----------+---------------------+--------------------------------------+---------------+------+
| id | system_id |    dispatched_at    |               payload                | marker_letter | rank |
+----+-----------+---------------------+--------------------------------------+---------------+------+
|    | SYS-0001  | 2025-12-24T08:10:00 | {"marker": "A", "source": "primary"} | A             | 1    |
|    | SYS-0002  | 2025-12-24T08:10:01 | {"marker": "D", "source": "primary"} | D             | 1    |
|    | SYS-0003  | 2025-12-24T08:10:02 | {"marker": "V", "source": "primary"} | V             | 1    |
|    | SYS-0004  | 2025-12-24T08:10:03 | {"marker": "E", "source": "primary"} | E             | 1    |
|    | SYS-0005  | 2025-12-24T08:10:04 | {"marker": "N", "source": "primary"} | N             | 1    |
|    | SYS-0006  | 2025-12-24T08:10:05 | {"marker": "T", "source": "primary"} | T             | 1    |
|    | SYS-0007  | 2025-12-24T08:10:06 | {"marker": "O", "source": "primary"} | O             | 1    |
|    | SYS-0008  | 2025-12-24T08:10:07 | {"marker": "F", "source": "primary"} | F             | 1    |
|    | SYS-0009  | 2025-12-24T08:10:08 | {"marker": "S", "source": "primary"} | S             | 1    |
|    | SYS-0010  | 2025-12-24T08:10:09 | {"marker": "Q", "source": "primary"} | Q             | 1    |
|    | SYS-0011  | 2025-12-24T08:10:10 | {"marker": "L", "source": "primary"} | L             | 1    |
+----+-----------+---------------------+--------------------------------------+---------------+------+
sqlite>
```

Phew!

That was a good one to end the advent of SQL!

I enjoyed it!

I learnt a ton
- CTEs
- JOINs (some wired stuff can be done)
- Window Functions (LAG, LEAD, ROW_NUMBER)
- FTS (in SQLite)
- Json parsing 
- date shenanigans
- CTEs don't support insert and delete in SQLite (it ruined my day 10 solution)
- Recursive CTEs
- String manipulation (thanks to xml parsing)
- I need to write a post about explaining what I learned.

Thanks [Aaron francis](https://aaronfrancis.com/) from [databaseschool](https://databaseschool.com/) for this challenge and explaining in depth on each day(i didn't watch all) and [Kelsey Petrich](https://x.com/krpetrich) for the lore of each problem, those were really lovely to read!

Happy Coding :)
Merry Christmas
Happy New year
Whatever you celebrate!
