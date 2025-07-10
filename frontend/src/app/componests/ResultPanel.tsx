import { ScanResponse } from "@/lib/dto/scan.dto";
import { Pending } from "./Pending";
import { ShowScanResult } from "./ShowScanResult";
import { Success } from "./Success";
import { Error } from "./Error";

interface ResultPanelProps {
  data?: ScanResponse;
  isSuccess?: boolean;
  isPending?: boolean;
  isError?: boolean;
}

export const ResultPanel = ({
  data,
  isSuccess,
  isPending,
  isError,
}: ResultPanelProps) => {
  if (isError) {
    return <Error />;
  }

  if (isPending) {
    return <Pending />;
  }

  if (isSuccess && data && data.findings.length > 0) {
    return <ShowScanResult findings={data.findings} />;
  }

  if (isSuccess && data && data.findings.length === 0) {
    return <Success />;
  }

  return (
    <div className="text-gray-500">
      Shall we look for some dangerous things?
    </div>
  );
};
