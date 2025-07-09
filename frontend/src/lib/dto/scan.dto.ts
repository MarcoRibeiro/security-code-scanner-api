import { z } from "zod/v4";
import {
  CreateScanConfigurationSchema,
  CreateScanRequestSchema,
  ScanFindingSchema,
  ScanResponseSchema,
} from "./scan.schema";

export type ScanFinding = z.infer<typeof ScanFindingSchema>;
export type ScanResponse = z.infer<typeof ScanResponseSchema>;
export type CreateScanConfiguration = z.infer<
  typeof CreateScanConfigurationSchema
>;
export type CreateScanRequest = z.infer<typeof CreateScanRequestSchema>;
