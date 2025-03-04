// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package gen

import (
	"context"
	"database/sql"
)

type Querier interface {
	//DeleteRatelimitNamespace
	//
	//  UPDATE `ratelimit_namespaces`
	//  SET deleted_at = ?
	//  WHERE id = ?
	DeleteRatelimitNamespace(ctx context.Context, arg DeleteRatelimitNamespaceParams) (sql.Result, error)
	//DeleteRatelimitOverride
	//
	//  UPDATE `ratelimit_overrides`
	//  SET
	//      deleted_at =  ?
	//  WHERE id = ?
	DeleteRatelimitOverride(ctx context.Context, arg DeleteRatelimitOverrideParams) (sql.Result, error)
	//FindKeyByHash
	//
	//  SELECT
	//      id, key_auth_id, hash, start, workspace_id, for_workspace_id, name, owner_id, identity_id, meta, created_at, expires, created_at_m, updated_at_m, deleted_at_m, deleted_at, refill_day, refill_amount, last_refill_at, enabled, remaining_requests, ratelimit_async, ratelimit_limit, ratelimit_duration, environment
	//  FROM `keys`
	//  WHERE hash = ?
	FindKeyByHash(ctx context.Context, hash string) (Key, error)
	//FindKeyByID
	//
	//  SELECT
	//      k.id, k.key_auth_id, k.hash, k.start, k.workspace_id, k.for_workspace_id, k.name, k.owner_id, k.identity_id, k.meta, k.created_at, k.expires, k.created_at_m, k.updated_at_m, k.deleted_at_m, k.deleted_at, k.refill_day, k.refill_amount, k.last_refill_at, k.enabled, k.remaining_requests, k.ratelimit_async, k.ratelimit_limit, k.ratelimit_duration, k.environment,
	//      i.id, i.external_id, i.workspace_id, i.environment, i.created_at, i.updated_at, i.meta
	//  FROM `keys` k
	//  LEFT JOIN identities i ON k.identity_id = i.id
	//  WHERE k.id = ?
	FindKeyByID(ctx context.Context, id string) (FindKeyByIDRow, error)
	//FindKeyForVerification
	//
	//  WITH direct_permissions AS (
	//      SELECT kp.key_id, p.name as permission_name
	//      FROM keys_permissions kp
	//      JOIN permissions p ON kp.permission_id = p.id
	//  ),
	//  role_permissions AS (
	//      SELECT kr.key_id, p.name as permission_name
	//      FROM keys_roles kr
	//      JOIN roles_permissions rp ON kr.role_id = rp.role_id
	//      JOIN permissions p ON rp.permission_id = p.id
	//  ),
	//  all_permissions AS (
	//      SELECT key_id, permission_name FROM direct_permissions
	//      UNION
	//      SELECT key_id, permission_name FROM role_permissions
	//  ),
	//  all_ratelimits AS (
	//      SELECT
	//          key_id as target_id,
	//          'key' as target_type,
	//          name,
	//          `limit`,
	//          duration
	//      FROM ratelimits
	//      WHERE key_id IS NOT NULL
	//      UNION
	//      SELECT
	//          identity_id as target_id,
	//          'identity' as target_type,
	//          name,
	//          `limit`,
	//          duration
	//      FROM ratelimits
	//      WHERE identity_id IS NOT NULL
	//  )
	//  SELECT
	//      k.id, k.key_auth_id, k.hash, k.start, k.workspace_id, k.for_workspace_id, k.name, k.owner_id, k.identity_id, k.meta, k.created_at, k.expires, k.created_at_m, k.updated_at_m, k.deleted_at_m, k.deleted_at, k.refill_day, k.refill_amount, k.last_refill_at, k.enabled, k.remaining_requests, k.ratelimit_async, k.ratelimit_limit, k.ratelimit_duration, k.environment,
	//      i.id, i.external_id, i.workspace_id, i.environment, i.created_at, i.updated_at, i.meta,
	//      JSON_ARRAYAGG(
	//          JSON_OBJECT(
	//              'target_type', rl.target_type,
	//              'name', rl.name,
	//              'limit', rl.limit,
	//              'duration', rl.duration
	//          )
	//      ) as ratelimits,
	//      GROUP_CONCAT(DISTINCT perms.permission_name) as permissions
	//  FROM `keys` k
	//  LEFT JOIN identities i ON k.identity_id = i.id
	//  LEFT JOIN all_permissions perms ON k.id = perms.key_id
	//  LEFT JOIN all_ratelimits rl ON (
	//      (rl.target_type = 'key' AND rl.target_id = k.id) OR
	//      (rl.target_type = 'identity' AND rl.target_id = k.identity_id)
	//  )
	//  WHERE k.hash = ?
	//  GROUP BY k.id
	FindKeyForVerification(ctx context.Context, hash string) (FindKeyForVerificationRow, error)
	//FindPermissionsForKey
	//
	//  WITH direct_permissions AS (
	//      SELECT p.name as permission_name
	//      FROM keys_permissions kp
	//      JOIN permissions p ON kp.permission_id = p.id
	//      WHERE kp.key_id = ?
	//  ),
	//  role_permissions AS (
	//      SELECT p.name as permission_name
	//      FROM keys_roles kr
	//      JOIN roles_permissions rp ON kr.role_id = rp.role_id
	//      JOIN permissions p ON rp.permission_id = p.id
	//      WHERE kr.key_id = ?
	//  )
	//  SELECT DISTINCT permission_name
	//  FROM (
	//      SELECT permission_name FROM direct_permissions
	//      UNION ALL
	//      SELECT permission_name FROM role_permissions
	//  ) all_permissions
	FindPermissionsForKey(ctx context.Context, arg FindPermissionsForKeyParams) ([]string, error)
	//FindRatelimitNamespaceByID
	//
	//  SELECT id, workspace_id, name, created_at, updated_at, deleted_at FROM `ratelimit_namespaces`
	//  WHERE id = ?
	FindRatelimitNamespaceByID(ctx context.Context, id string) (RatelimitNamespace, error)
	//FindRatelimitNamespaceByName
	//
	//  SELECT id, workspace_id, name, created_at, updated_at, deleted_at FROM `ratelimit_namespaces`
	//  WHERE name = ?
	//  AND workspace_id = ?
	FindRatelimitNamespaceByName(ctx context.Context, arg FindRatelimitNamespaceByNameParams) (RatelimitNamespace, error)
	//FindRatelimitOverridesById
	//
	//  SELECT id, workspace_id, namespace_id, identifier, `limit`, duration, async, sharding, created_at, updated_at, deleted_at FROM ratelimit_overrides
	//  WHERE
	//      workspace_id = ?
	//      AND id = ?
	FindRatelimitOverridesById(ctx context.Context, arg FindRatelimitOverridesByIdParams) (RatelimitOverride, error)
	//FindRatelimitOverridesByIdentifier
	//
	//  SELECT id, workspace_id, namespace_id, identifier, `limit`, duration, async, sharding, created_at, updated_at, deleted_at FROM ratelimit_overrides
	//  WHERE
	//      workspace_id = ?
	//      AND namespace_id = ?
	//      AND identifier LIKE ?
	FindRatelimitOverridesByIdentifier(ctx context.Context, arg FindRatelimitOverridesByIdentifierParams) ([]RatelimitOverride, error)
	//FindWorkspaceByID
	//
	//  SELECT id, tenant_id, name, created_at, deleted_at, plan, stripe_customer_id, stripe_subscription_id, trial_ends, beta_features, features, plan_locked_until, plan_downgrade_request, plan_changed, subscriptions, enabled, delete_protection FROM `workspaces`
	//  WHERE id = ?
	FindWorkspaceByID(ctx context.Context, id string) (Workspace, error)
	//HardDeleteWorkspace
	//
	//  DELETE FROM `workspaces`
	//  WHERE id = ?
	//  AND delete_protection = false
	HardDeleteWorkspace(ctx context.Context, id string) (sql.Result, error)
	//InsertKey
	//
	//  INSERT INTO `keys` (
	//      id,
	//      key_auth_id,
	//      hash,
	//      start,
	//      workspace_id,
	//      for_workspace_id,
	//      name,
	//      owner_id,
	//      identity_id,
	//      meta,
	//      created_at,
	//      expires,
	//      created_at_m,
	//      enabled,
	//      remaining_requests,
	//      ratelimit_async,
	//      ratelimit_limit,
	//      ratelimit_duration,
	//      environment
	//  ) VALUES (
	//      ?,
	//      ?,
	//      ?,
	//      ?,
	//      ?,
	//      ?,
	//      ?,
	//      null,
	//      ?,
	//      ?,
	//      ?,
	//      ?,
	//      UNIX_TIMESTAMP() * 1000,
	//      ?,
	//      ?,
	//      ?,
	//      ?,
	//      ?,
	//      ?
	//  )
	InsertKey(ctx context.Context, arg InsertKeyParams) error
	//InsertKeyring
	//
	//  INSERT INTO `key_auth` (
	//      id,
	//      workspace_id,
	//      created_at,
	//      created_at_m,
	//      store_encrypted_keys,
	//      default_prefix,
	//      default_bytes,
	//      size_approx,
	//      size_last_updated_at
	//  ) VALUES (
	//      ?,
	//      ?,
	//       ?,
	//        ?,
	//      ?,
	//      ?,
	//      ?,
	//      0,
	//      0
	//  )
	InsertKeyring(ctx context.Context, arg InsertKeyringParams) error
	//InsertRatelimitNamespace
	//
	//  INSERT INTO
	//      `ratelimit_namespaces` (
	//          id,
	//          workspace_id,
	//          name,
	//          created_at,
	//          updated_at,
	//          deleted_at
	//          )
	//  VALUES
	//      (
	//          ?,
	//          ?,
	//          ?,
	//           ?,
	//          NULL,
	//          NULL
	//      )
	InsertRatelimitNamespace(ctx context.Context, arg InsertRatelimitNamespaceParams) error
	//InsertRatelimitOverride
	//
	//  INSERT INTO
	//      `ratelimit_overrides` (
	//          id,
	//          workspace_id,
	//          namespace_id,
	//          identifier,
	//          `limit`,
	//          duration,
	//          async,
	//          created_at
	//      )
	//  VALUES
	//      (
	//          ?,
	//          ?,
	//          ?,
	//          ?,
	//          ?,
	//          ?,
	//          false,
	//           ?
	//      )
	InsertRatelimitOverride(ctx context.Context, arg InsertRatelimitOverrideParams) error
	//InsertWorkspace
	//
	//  INSERT INTO `workspaces` (
	//      id,
	//      tenant_id,
	//      name,
	//      created_at,
	//      plan,
	//      beta_features,
	//      features,
	//      enabled,
	//      delete_protection
	//  )
	//  VALUES (
	//      ?,
	//      ?,
	//      ?,
	//       ?,
	//      'free',
	//      '{}',
	//      '{}',
	//      true,
	//      true
	//  )
	InsertWorkspace(ctx context.Context, arg InsertWorkspaceParams) error
	//SoftDeleteWorkspace
	//
	//  UPDATE `workspaces`
	//  SET deleted_at = ?
	//  WHERE id = ?
	//  AND delete_protection = false
	SoftDeleteWorkspace(ctx context.Context, arg SoftDeleteWorkspaceParams) (sql.Result, error)
	//UpdateRatelimitOverride
	//
	//  UPDATE `ratelimit_overrides`
	//  SET
	//      `limit` = ?,
	//      duration = ?,
	//      async = ?,
	//      updated_at = ?
	//  WHERE id = ?
	UpdateRatelimitOverride(ctx context.Context, arg UpdateRatelimitOverrideParams) (sql.Result, error)
	//UpdateWorkspaceEnabled
	//
	//  UPDATE `workspaces`
	//  SET enabled = ?
	//  WHERE id = ?
	UpdateWorkspaceEnabled(ctx context.Context, arg UpdateWorkspaceEnabledParams) (sql.Result, error)
	//UpdateWorkspacePlan
	//
	//  UPDATE `workspaces`
	//  SET plan = ?
	//  WHERE id = ?
	UpdateWorkspacePlan(ctx context.Context, arg UpdateWorkspacePlanParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)
