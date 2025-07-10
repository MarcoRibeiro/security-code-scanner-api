import { CreateScanRequest, ScanResponse } from "../dto/scan.dto";
import { client } from "./client";

export const createScan = async (
  createScanRequest: CreateScanRequest
): Promise<ScanResponse> => {
  if (typeof createScanRequest.configuration?.exclude == "string") {
    createScanRequest.configuration.exclude =
      createScanRequest.configuration.exclude
        .split(",")
        .map((s) => s.trim())
        .filter((x) => x != "");
  }

  const response = await client.post<ScanResponse>(
    "/v1/scans",
    createScanRequest
  );
  return response.data;
};
