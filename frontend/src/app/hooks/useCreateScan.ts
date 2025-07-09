import { CreateScanRequest, ScanResponse } from "@/lib/dto/scan.dto";
import { createScan } from "@/lib/services/scanService";
import { useMutation } from "@tanstack/react-query";

export const useCreateScan = () => {
  return useMutation<ScanResponse, Error, CreateScanRequest>({
    mutationFn: createScan,
  });
};
