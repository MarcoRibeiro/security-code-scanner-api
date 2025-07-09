import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { ScanFinding } from "@/lib/dto/scan.dto";
import { FileFindingsTable } from "./FileFindingsTable";
import { GeneralView } from "./GeneralView";

interface ShowScanResultProps {
  findings: ScanFinding[];
}

export const ShowScanResult = ({ findings }: ShowScanResultProps) => {
  const dict = findings.reduce<Record<string, ScanFinding[]>>(
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
      <GeneralView findings={findings} />
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
};
