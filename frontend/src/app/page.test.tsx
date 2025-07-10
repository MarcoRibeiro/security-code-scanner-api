import { render, screen, within } from "@testing-library/react";
import Home from "./page";
import "@testing-library/jest-dom";

jest.mock("./componests/NewScanForm", () => {
  const NewScanForm = ({ createScan, isPending, isSuccess }: any) => {
    return (
      <div data-testid="new-scan-form">
        NewScanForm - isPending: {String(isPending)}, isSuccess:{" "}
        {String(isSuccess)}
      </div>
    );
  };
  return { NewScanForm };
});

jest.mock("./componests/ResultPanel", () => {
  const ResultPanel = ({ isSuccess, isPending, isError, data }: any) => {
    return (
      <div data-testid="result-panel">
        ResultPanel - isSuccess: {String(isSuccess)}, isPending:{" "}
        {String(isPending)}, isError: {String(isError)}
      </div>
    );
  };
  return { ResultPanel };
});

jest.mock("./hooks/useCreateScan", () => ({
  useCreateScan: () => ({
    data: { results: [] },
    isError: false,
    isSuccess: true,
    isPending: false,
    mutate: jest.fn(),
  }),
}));

describe("Home page", () => {
  it("renders the NewScanForm and ResultPanel with correct props", () => {
    render(<Home />);

    expect(screen.getByTestId("new-scan-form")).toBeInTheDocument();
    expect(screen.getByTestId("result-panel")).toBeInTheDocument();

    const newScanForm = screen.getByTestId("new-scan-form");
    expect(
      within(newScanForm).getByText(/isPending: false/i)
    ).toBeInTheDocument();
    expect(
      within(newScanForm).getByText(/isSuccess: true/i)
    ).toBeInTheDocument();

    const resultPanel = screen.getByTestId("result-panel");
    expect(
      within(resultPanel).getByText(/isPending: false/i)
    ).toBeInTheDocument();
    expect(
      within(resultPanel).getByText(/isSuccess: true/i)
    ).toBeInTheDocument();
    expect(
      within(resultPanel).getByText(/isError: false/i)
    ).toBeInTheDocument();
  });
});
