"use client";

import { CreateScanRequest } from "@/lib/dto/scan.dto";
import { useCreateScan } from "@/lib/hooks/useCreateScan";
import { useForm } from "react-hook-form";

export default function Home() {
  const { register, handleSubmit } = useForm<CreateScanRequest>();
  const { data, isSuccess, isPending, mutate: createScan } = useCreateScan();

  const handleCreateScan = async (data: CreateScanRequest) => {
    createScan(data);
  };

  if (isPending) {
    return <div>Creating scan...</div>;
  }

  if (isSuccess) {
    return (
      <div>
        {data.findings.map((finding, i) => (
          <div key={i}>{finding.message}</div>
        ))}
      </div>
    );
  }

  return (
    <main>
      <div>New Scan</div>
      <form onSubmit={handleSubmit(handleCreateScan)}>
        <input {...register("path")} />
        <input {...register("configuration.exclude")} />
        <button type="submit">Create Scan</button>
      </form>
    </main>
  );
}
