import { useMutation } from "@tanstack/react-query";
import { CreateScanRequest, ScanResponse } from "../dto/scan.dto";
import { createScan } from "../services/scanService";

export const useCreateScan = () => {
  return useMutation<ScanResponse, Error, CreateScanRequest>({
    mutationFn: createScan,
  });
};
