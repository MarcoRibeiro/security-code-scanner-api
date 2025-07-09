import { z } from "zod/v4";
import {
  createScanRequestSchema,
  scanFindingSchema,
  scanResponseSchema,
} from "./scan.schema";

export type ScanFinding = z.infer<typeof scanFindingSchema>;
export type ScanResponse = z.infer<typeof scanResponseSchema>;
export type CreateScanRequest = z.infer<typeof createScanRequestSchema>;
