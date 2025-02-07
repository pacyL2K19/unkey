import { TRPCError } from "@trpc/server";
import { z } from "zod";

import { insertAuditLogs } from "@/lib/audit";
import { db, eq, schema } from "@/lib/db";
import { auth, t } from "../../trpc";

export const setDefaultApiPrefix = t.procedure
  .use(auth)
  .input(
    z.object({
      defaultPrefix: z.string().max(8, "Prefix can be a maximum of 8 characters"),
      keyAuthId: z.string(),
    }),
  )
  .mutation(async ({ ctx, input }) => {
    const keyAuth = await db.query.keyAuth
      .findFirst({
        where: (table, { eq, and, isNull }) =>
          and(
            eq(table.workspaceId, ctx.workspace.id),
            eq(table.id, input.keyAuthId),
            isNull(table.deletedAt),
          ),
      })
      .catch((_err) => {
        throw new TRPCError({
          code: "INTERNAL_SERVER_ERROR",
          message:
            "We were unable to update the key auth. Please try again or contact support@unkey.dev",
        });
      });
    if (!keyAuth) {
      throw new TRPCError({
        code: "NOT_FOUND",
        message:
          "We are unable to find the correct key auth. Please try again or contact support@unkey.dev.",
      });
    }

    await db
      .transaction(async (tx) => {
        await tx
          .update(schema.keyAuth)
          .set({
            defaultPrefix: input.defaultPrefix,
          })
          .where(eq(schema.keyAuth.id, keyAuth.id))
          .catch((_err) => {
            throw new TRPCError({
              code: "INTERNAL_SERVER_ERROR",
              message:
                "We were unable to update the API default prefix. Please try again or contact support@unkey.dev.",
            });
          });
        await insertAuditLogs(tx, ctx.workspace.auditLogBucket.id, {
          workspaceId: ctx.workspace.id,
          actor: {
            type: "user",
            id: ctx.user.id,
          },
          event: "api.update",
          description: `Changed ${keyAuth.id} default prefix from ${keyAuth.defaultPrefix} to ${input.defaultPrefix}`,
          resources: [
            {
              type: "keyAuth",
              id: keyAuth.id,
            },
          ],
          context: {
            location: ctx.audit.location,
            userAgent: ctx.audit.userAgent,
          },
        });
      })
      .catch((_err) => {
        throw new TRPCError({
          code: "INTERNAL_SERVER_ERROR",
          message:
            "We were unable to update the default prefix. Please try again or contact support@unkey.dev.",
        });
      });
  });
