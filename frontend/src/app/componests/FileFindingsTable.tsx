import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { ScanFinding } from "@/lib/dto/scan.dto";

interface FileFindingsTableProps {
  findings: ScanFinding[];
}

export const FileFindingsTable = ({ findings }: FileFindingsTableProps) => {
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Type</TableHead>
          <TableHead>Finding</TableHead>
          <TableHead>Line</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {findings.map((finding) => (
          <TableRow key={`${finding.rule}-${finding.line}`}>
            <TableCell>{finding.rule}</TableCell>
            <TableCell>{finding.message}</TableCell>
            <TableCell>{finding.line}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};
