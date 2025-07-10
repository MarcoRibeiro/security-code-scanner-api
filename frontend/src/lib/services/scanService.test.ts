import { CreateScanRequest, ScanResponse } from "../dto/scan.dto";
import { client } from "./client";
import { createScan } from "./scanService";

jest.mock("./client", () => ({
  client: {
    post: jest.fn(),
  },
}));

describe("createScan", () => {
  it("should convert the exclude string into a array", async () => {
    const request: CreateScanRequest = {
      path: "http://github.com/git.git",
      configuration: {
        exclude: "node_modules, dist ,  ",
      },
    };

    const mockResponse: ScanResponse = {
      path: "http://github.com/git.git",
      done: true,
      findings: [],
    };

    (client.post as jest.Mock).mockResolvedValueOnce({ data: mockResponse });

    // act
    const result = await createScan(request);

    // assert
    expect(client.post).toHaveBeenCalledWith("/v1/scans", {
      path: "http://github.com/git.git",
      configuration: {
        exclude: ["node_modules", "dist"],
      },
    });

    expect(result).toEqual(mockResponse);
  });
});
