import path from "path";
import { z } from "zod/v4";

export const scanFindingSchema = z.object({
  rule: z.string(),
  file: z.string(),
  message: z.string(),
  line: z.number(),
});

export const scanResponseSchema = z.object({
  done: z.boolean(),
  path: z.string(),
  findings: z.array(scanFindingSchema),
});

export const createScanRequestSchema = z.object({
  path: z.string().nonempty(),
  configuration: z.object({
    exclude: z.union([z.array(z.string()), z.string()]).optional(),
  }),
});
