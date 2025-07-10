"use client";

import { NewScanForm } from "./componests/NewScanForm";
import { ResultPanel } from "./componests/ResultPanel";
import { useCreateScan } from "./hooks/useCreateScan";

export default function Home() {
  const {
    data,
    isError,
    isSuccess,
    isPending,
    mutate: createScan,
  } = useCreateScan();

  return (
    <div className="grid grid-cols-2 w-screen h-[calc(100vh-100px)]">
      <div className="col-span-1 p-4 border-r">
        <NewScanForm
          createScan={createScan}
          isPending={isPending}
          isSuccess={isSuccess}
        />
      </div>
      <div className="col-span-1 p-4 overflow-y-auto">
        <ResultPanel
          data={data}
          isSuccess={isSuccess}
          isPending={isPending}
          isError={isError}
        />
      </div>
    </div>
  );
}
