import { render, screen } from "@testing-library/react";
import { FileFindingsTable } from "./FileFindingsTable";
import { ScanFinding } from "@/lib/dto/scan.dto";
import "@testing-library/jest-dom";

describe("FileFindingsTable", () => {
  const mockFindings: ScanFinding[] = [
    {
      rule: "XSS",
      message: "Possible cross-site scripting",
      line: 12,
      file: "fileOne",
    },
    {
      rule: "SQL_INJECTION",
      message: "Possible SQL injection",
      line: 45,
      file: "fileOne",
    },
  ];

  it("renders the table headers", () => {
    render(<FileFindingsTable findings={mockFindings} />);
    expect(screen.getByText("Type")).toBeInTheDocument();
    expect(screen.getByText("Finding")).toBeInTheDocument();
    expect(screen.getByText("Line")).toBeInTheDocument();
  });

  it("renders the correct number of rows", () => {
    render(<FileFindingsTable findings={mockFindings} />);
    const rows = screen.getAllByRole("row");
    expect(rows).toHaveLength(3);
  });

  it("renders each finding correctly", () => {
    render(<FileFindingsTable findings={mockFindings} />);
    expect(screen.getByText("XSS")).toBeInTheDocument();
    expect(
      screen.getByText("Possible cross-site scripting")
    ).toBeInTheDocument();
    expect(screen.getByText("12")).toBeInTheDocument();

    expect(screen.getByText("SQL_INJECTION")).toBeInTheDocument();
    expect(screen.getByText("Possible SQL injection")).toBeInTheDocument();
    expect(screen.getByText("45")).toBeInTheDocument();
  });
});
