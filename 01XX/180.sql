SELECT DISTINCT num AS ConsecutiveNums
FROM (
    SELECT num,
           LEAD(num) OVER (ORDER BY id) AS next,
           LAG(num) OVER (ORDER BY id)  AS prev
    FROM Logs
) q
WHERE num = next AND num = prev