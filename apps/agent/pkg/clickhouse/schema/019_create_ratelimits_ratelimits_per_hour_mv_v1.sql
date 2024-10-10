-- +goose up
CREATE MATERIALIZED VIEW ratelimits.ratelimits_per_hour_mv_v1
TO ratelimits.ratelimits_per_hour_v1
AS
SELECT
  workspace_id,
  namespace_id,
  identifier,
  pass,
  count(*) as count,
  toStartOfHour(fromUnixTimestamp64Milli(time)) AS time
FROM ratelimits.raw_ratelimits_v1
GROUP BY
  workspace_id,
  namespace_id,
  identifier,
  pass,
  time
;
