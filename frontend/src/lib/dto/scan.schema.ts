import { z } from "zod/v4";

export const ScanFindingSchema = z.object({
  rule: z.string(),
  file: z.string(),
  message: z.string(),
  line: z.number().optional(),
});

export const ScanResponseSchema = z.object({
  findings: z.array(ScanFindingSchema),
});

export const CreateScanConfigurationSchema = z.object({
  exclude: z.array(z.string()).optional(),
});

export const CreateScanRequestSchema = z.object({
  path: z.string(),
  configuration: CreateScanConfigurationSchema.optional(),
});
