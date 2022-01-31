--  1

SELECT email AS Email
FROM
  (SELECT email,
          COUNT(email) AS num
   FROM Person
   GROUP BY email) AS statistic
WHERE num > 1;

--  2

SELECT email AS Email
FROM Person
GROUP BY email
HAVING COUNT(email) > 1;
