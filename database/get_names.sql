-- THIS FILE IS UNUSED AND ONLY FOR DEMONSTRATION

-- strong one params
-- param1, param2
SELECT users."uid", users."first_name", users."last_name"
FROM (SELECT sdn_items.uid as "uid", sdn_akas.first_name as "first_name", sdn_akas.last_name as "last_name"
      FROM sdn_items
               INNER JOIN items_akas ON sdn_items.uid = items_akas.sdn_item_uid
               INNER JOIN sdn_akas ON sdn_akas.uid = items_akas.sdn_aka_uid
      UNION
      SELECT sdn_items.uid as "uid", sdn_items.first_name as "first_name", sdn_items.last_name as "last_name"
      FROM sdn_items) as users
WHERE LOWER(first_name) = 'muzonzini' OR LOWER(last_name) = 'muzonzini';

-- strong two params
-- param1, param2, param1, param2
SELECT users."uid", users."first_name", users."last_name"
FROM (SELECT sdn_items.uid as "uid", sdn_akas.first_name as "first_name", sdn_akas.last_name as "last_name"
      FROM sdn_items
               INNER JOIN items_akas ON sdn_items.uid = items_akas.sdn_item_uid
               INNER JOIN sdn_akas ON sdn_akas.uid = items_akas.sdn_aka_uid
      UNION
      SELECT sdn_items.uid as "uid", sdn_items.first_name as "first_name", sdn_items.last_name as "last_name"
      FROM sdn_items) as users
WHERE LOWER(first_name) = 'elisha' AND LOWER(last_name) = 'muzonzini'
   OR LOWER(last_name) = 'elisha' AND LOWER(first_name) = 'muzonzini';

-- weak two params
-- param1, param2, param1, param2, param1, param2, param1, param2
-- weak one param
-- param1, param1, param1, param1, param1, param1, param1, param1
SELECT *
FROM (SELECT sdn_items.uid as "uid", sdn_akas.first_name as "first_name", sdn_akas.last_name as "last_name"
      FROM sdn_items
               INNER JOIN items_akas ON sdn_items.uid = items_akas.sdn_item_uid
               INNER JOIN sdn_akas ON sdn_akas.uid = items_akas.sdn_aka_uid
      UNION
      SELECT sdn_items.uid as "uid", sdn_items.first_name as "first_name", sdn_items.last_name as "last_name"
      FROM sdn_items) as users
WHERE users."uid" = (SELECT best_match.item_id
                     FROM (SELECT users."uid" as item_id, COUNT(*) as cnt
                           FROM (SELECT sdn_items.uid as "uid", null as "first_name", sdn_akas.last_name as "last_name"
                                 FROM sdn_items
                                          INNER JOIN items_akas ON sdn_items.uid = items_akas.sdn_item_uid
                                          INNER JOIN sdn_akas ON sdn_akas.uid = items_akas.sdn_aka_uid
                                 WHERE LOWER(sdn_akas.last_name) LIKE '%mohammed%'
                                    OR LOWER(sdn_akas.last_name) LIKE '%musa%'
                                    OR LOWER(sdn_akas.first_name) LIKE '%mohammed%'
                                    OR LOWER(sdn_akas.first_name) LIKE '%musa%'
                                 UNION
                                 SELECT sdn_items.uid        as "uid",
                                        sdn_items.first_name as "first_name",
                                        sdn_items.last_name  as "last_name"
                                 FROM sdn_items
                                 WHERE LOWER(sdn_items.last_name) LIKE '%mohammed%'
                                    OR LOWER(sdn_items.first_name) LIKE '%musa%'
                                    OR LOWER(sdn_items.first_name) LIKE '%mohammed%'
                                    OR LOWER(sdn_items.last_name) LIKE '%musa%') as users
                           GROUP BY "uid"
                           ORDER BY cnt DESC
                           LIMIT 1) AS best_match);
