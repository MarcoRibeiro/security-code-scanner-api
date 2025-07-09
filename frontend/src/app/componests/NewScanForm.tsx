import { CreateScanRequest } from "@/lib/dto/scan.dto";
import { useForm } from "react-hook-form";

import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { zodResolver } from "@hookform/resolvers/zod";
import { createScanRequestSchema } from "@/lib/dto/scan.schema";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Button } from "@/components/ui/button";

interface NewScanFormProps {
  isPending: boolean;
  isSuccess: boolean;
  createScan: (data: CreateScanRequest) => void;
}

export const NewScanForm = ({ isPending, createScan }: NewScanFormProps) => {
  const form = useForm<CreateScanRequest>({
    resolver: zodResolver(createScanRequestSchema),
    defaultValues: {
      path: "",
      configuration: {
        exclude: [],
      },
    },
  });

  const handleCreateScan = async (data: CreateScanRequest) => {
    createScan(data);
  };

  return (
    <>
      <Form {...form}>
        <form
          id="new-scan-form"
          onSubmit={form.handleSubmit(handleCreateScan)}
          className="space-y-8 bg-white/10 rounded-xl shadow-xl p-8 border border-white/20"
        >
          <FormField
            control={form.control}
            name="path"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Path</FormLabel>
                <FormControl>
                  <Input disabled={isPending} placeholder="Path" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="configuration.exclude"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Exclusions</FormLabel>
                <FormControl>
                  <Textarea
                    disabled={isPending}
                    placeholder="ex: node_modules, dist"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button form="new-scan-form" className="w-full">
            {isPending ? "Scanning..." : "Start Scan"}
          </Button>
        </form>
      </Form>
    </>
  );
};
