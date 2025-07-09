import { ScanFinding } from "@/lib/dto/scan.dto";

interface GeneralViewProps {
  findings: ScanFinding[];
}

export const GeneralView = ({ findings }: GeneralViewProps) => {
  const totalFindings = findings.length;
  const findingsByRule = findings.reduce<Record<string, number>>(
    (acc, finding) => {
      acc[finding.rule] = (acc[finding.rule] || 0) + 1;
      return acc;
    },
    {}
  );

  return (
    <div className="mb-4 p-4 rounded-lg bg-black/30 text-white flex flex-col gap-2">
      <span className="font-bold text-lg">Scan Summary</span>
      <span>
        Total of findings:{" "}
        <span className="font-semibold text-yellow-300">{totalFindings}</span>
      </span>
      <div className="flex flex-wrap gap-4 mt-2">
        {Object.entries(findingsByRule).map(([rule, count]) => (
          <span
            key={rule}
            className="bg-yellow-900/60 px-3 py-1 rounded text-yellow-200 text-sm"
          >
            {rule}: <span className="font-bold">{count}</span>
          </span>
        ))}
      </div>
    </div>
  );
};
