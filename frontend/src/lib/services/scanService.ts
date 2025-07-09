import { CreateScanRequest, ScanResponse } from "../dto/scan.dto";
import { client } from "./client";

export const createScan = async (
  createScanRequest: CreateScanRequest
): Promise<ScanResponse> => {
  const response = await client.post<ScanResponse>(
    "/v1/scans",
    createScanRequest
  );
  return response.data;
};
