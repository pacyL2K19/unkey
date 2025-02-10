// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: verify_key.sql

package gen

import (
	"context"
	"database/sql"
)

const verifyKey = `-- name: VerifyKey :one
WITH direct_permissions AS (
    SELECT kp.key_id, p.name as permission_name
    FROM keys_permissions kp
    JOIN permissions p ON kp.permission_id = p.id
),
role_permissions AS (
    SELECT kr.key_id, p.name as permission_name
    FROM keys_roles kr
    JOIN roles_permissions rp ON kr.role_id = rp.role_id
    JOIN permissions p ON rp.permission_id = p.id
),
all_permissions AS (
    SELECT key_id, permission_name FROM direct_permissions
    UNION
    SELECT key_id, permission_name FROM role_permissions
),
all_ratelimits AS (
    SELECT
        key_id as target_id,
        'key' as target_type,
        name,
        ` + "`" + `limit` + "`" + `,
        duration
    FROM ratelimits
    WHERE key_id IS NOT NULL
    UNION
    SELECT
        identity_id as target_id,
        'identity' as target_type,
        name,
        ` + "`" + `limit` + "`" + `,
        duration
    FROM ratelimits
    WHERE identity_id IS NOT NULL
)
SELECT
    k.id, k.key_auth_id, k.hash, k.start, k.workspace_id, k.for_workspace_id, k.name, k.owner_id, k.identity_id, k.meta, k.created_at, k.expires, k.created_at_m, k.updated_at_m, k.deleted_at_m, k.deleted_at, k.refill_day, k.refill_amount, k.last_refill_at, k.enabled, k.remaining_requests, k.ratelimit_async, k.ratelimit_limit, k.ratelimit_duration, k.environment,
    i.id, i.external_id, i.workspace_id, i.environment, i.created_at, i.updated_at, i.meta,
    GROUP_CONCAT(DISTINCT rl.target_type) as ratelimit_types,
    GROUP_CONCAT(DISTINCT rl.name) as ratelimit_names,
    GROUP_CONCAT(DISTINCT rl.limit) as ratelimit_limits,
    GROUP_CONCAT(DISTINCT rl.duration) as ratelimit_durations,
    GROUP_CONCAT(DISTINCT perms.permission_name) as permissions
FROM ` + "`" + `keys` + "`" + ` k
LEFT JOIN identities i ON k.identity_id = i.id
LEFT JOIN all_permissions perms ON k.id = perms.key_id
LEFT JOIN all_ratelimits rl ON (
    (rl.target_type = 'key' AND rl.target_id = k.id) OR
    (rl.target_type = 'identity' AND rl.target_id = k.identity_id)
)
WHERE k.hash = ?
GROUP BY k.id
`

type VerifyKeyRow struct {
	Key                Key            `db:"key"`
	Identity           Identity       `db:"identity"`
	RatelimitTypes     sql.NullString `db:"ratelimit_types"`
	RatelimitNames     sql.NullString `db:"ratelimit_names"`
	RatelimitLimits    sql.NullString `db:"ratelimit_limits"`
	RatelimitDurations sql.NullString `db:"ratelimit_durations"`
	Permissions        sql.NullString `db:"permissions"`
}

// VerifyKey
//
//	WITH direct_permissions AS (
//	    SELECT kp.key_id, p.name as permission_name
//	    FROM keys_permissions kp
//	    JOIN permissions p ON kp.permission_id = p.id
//	),
//	role_permissions AS (
//	    SELECT kr.key_id, p.name as permission_name
//	    FROM keys_roles kr
//	    JOIN roles_permissions rp ON kr.role_id = rp.role_id
//	    JOIN permissions p ON rp.permission_id = p.id
//	),
//	all_permissions AS (
//	    SELECT key_id, permission_name FROM direct_permissions
//	    UNION
//	    SELECT key_id, permission_name FROM role_permissions
//	),
//	all_ratelimits AS (
//	    SELECT
//	        key_id as target_id,
//	        'key' as target_type,
//	        name,
//	        `limit`,
//	        duration
//	    FROM ratelimits
//	    WHERE key_id IS NOT NULL
//	    UNION
//	    SELECT
//	        identity_id as target_id,
//	        'identity' as target_type,
//	        name,
//	        `limit`,
//	        duration
//	    FROM ratelimits
//	    WHERE identity_id IS NOT NULL
//	)
//	SELECT
//	    k.id, k.key_auth_id, k.hash, k.start, k.workspace_id, k.for_workspace_id, k.name, k.owner_id, k.identity_id, k.meta, k.created_at, k.expires, k.created_at_m, k.updated_at_m, k.deleted_at_m, k.deleted_at, k.refill_day, k.refill_amount, k.last_refill_at, k.enabled, k.remaining_requests, k.ratelimit_async, k.ratelimit_limit, k.ratelimit_duration, k.environment,
//	    i.id, i.external_id, i.workspace_id, i.environment, i.created_at, i.updated_at, i.meta,
//	    GROUP_CONCAT(DISTINCT rl.target_type) as ratelimit_types,
//	    GROUP_CONCAT(DISTINCT rl.name) as ratelimit_names,
//	    GROUP_CONCAT(DISTINCT rl.limit) as ratelimit_limits,
//	    GROUP_CONCAT(DISTINCT rl.duration) as ratelimit_durations,
//	    GROUP_CONCAT(DISTINCT perms.permission_name) as permissions
//	FROM `keys` k
//	LEFT JOIN identities i ON k.identity_id = i.id
//	LEFT JOIN all_permissions perms ON k.id = perms.key_id
//	LEFT JOIN all_ratelimits rl ON (
//	    (rl.target_type = 'key' AND rl.target_id = k.id) OR
//	    (rl.target_type = 'identity' AND rl.target_id = k.identity_id)
//	)
//	WHERE k.hash = ?
//	GROUP BY k.id
func (q *Queries) VerifyKey(ctx context.Context, hash string) (VerifyKeyRow, error) {
	row := q.db.QueryRowContext(ctx, verifyKey, hash)
	var i VerifyKeyRow
	err := row.Scan(
		&i.Key.ID,
		&i.Key.KeyAuthID,
		&i.Key.Hash,
		&i.Key.Start,
		&i.Key.WorkspaceID,
		&i.Key.ForWorkspaceID,
		&i.Key.Name,
		&i.Key.OwnerID,
		&i.Key.IdentityID,
		&i.Key.Meta,
		&i.Key.CreatedAt,
		&i.Key.Expires,
		&i.Key.CreatedAtM,
		&i.Key.UpdatedAtM,
		&i.Key.DeletedAtM,
		&i.Key.DeletedAt,
		&i.Key.RefillDay,
		&i.Key.RefillAmount,
		&i.Key.LastRefillAt,
		&i.Key.Enabled,
		&i.Key.RemainingRequests,
		&i.Key.RatelimitAsync,
		&i.Key.RatelimitLimit,
		&i.Key.RatelimitDuration,
		&i.Key.Environment,
		&i.Identity.ID,
		&i.Identity.ExternalID,
		&i.Identity.WorkspaceID,
		&i.Identity.Environment,
		&i.Identity.CreatedAt,
		&i.Identity.UpdatedAt,
		&i.Identity.Meta,
		&i.RatelimitTypes,
		&i.RatelimitNames,
		&i.RatelimitLimits,
		&i.RatelimitDurations,
		&i.Permissions,
	)
	return i, err
}
