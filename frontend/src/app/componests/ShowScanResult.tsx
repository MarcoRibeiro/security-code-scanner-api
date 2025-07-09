import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { ScanFinding, ScanResponse } from "@/lib/dto/scan.dto";
import { Success } from "./Success";
import { Pending } from "./Pending";
import { FileFindingsTable } from "./FileFindingsTable";
import { GeneralView } from "./GeneralView";
import { Badge } from "@/components/ui/badge";

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
    const dict = data.findings.reduce<Record<string, ScanFinding[]>>(
      (acc, scanFinding) => {
        if (!acc[scanFinding.file]) {
          acc[scanFinding.file] = [];
        }

        acc[scanFinding.file].push(scanFinding);
        return acc;
      },
      {}
    );

    return (
      <div className="space-y-8 bg-white/10 rounded-xl shadow-xl p-4 border border-white/20">
        <GeneralView findings={data.findings} />
        <Accordion type="single" collapsible>
          {Object.entries(dict).map(([file, findings]) => (
            <AccordionItem value={file} key={file}>
              <AccordionTrigger>{file}</AccordionTrigger>
              <AccordionContent>
                <FileFindingsTable findings={findings} />
              </AccordionContent>
            </AccordionItem>
          ))}
        </Accordion>
      </div>
    );
  }

  if (isSuccess && data && data.findings.length === 0) {
    return <Success />;
  }

  return <div className="text-gray-500">No findings to display.</div>;
};
