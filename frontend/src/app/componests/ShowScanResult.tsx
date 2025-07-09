import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { ScanResponse } from "@/lib/dto/scan.dto";
import { Success } from "./Success";
import { Pending } from "./Pending";

interface ShowScanResultProps {
  data?: ScanResponse;
  isSuccess?: boolean;
  isPending?: boolean;
}

export const ShowScanResult = ({
  data,
  isSuccess,
  isPending,
}: ShowScanResultProps) => {
  if (isPending) {
    return <Pending />;
  }

  if (isSuccess && data && data.findings.length > 0) {
    return (
      <>
        <h2>Scan Findings</h2>
        <Accordion type="single" collapsible>
          {data?.findings.map((finding, i) => (
            <AccordionItem value={`item-${i}`} key={i}>
              <AccordionTrigger>{finding.file}</AccordionTrigger>
              <AccordionContent>
                {finding.message} line: {finding.line}
              </AccordionContent>
            </AccordionItem>
          ))}
        </Accordion>
      </>
    );
  }

  if (isSuccess && data && data.findings.length === 0) {
    return <Success />;
  }

  return <div className="text-gray-500">No findings to display.</div>;
};
