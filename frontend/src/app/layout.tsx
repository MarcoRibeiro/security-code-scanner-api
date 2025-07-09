import type { Metadata } from "next";
import "./globals.css";
import { Providers } from "./providers";

export const metadata: Metadata = {
  title: "Code Scanner",
  description: "Code Scanner Application",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="min-h-screen m-0 bg-gradient-to-br from-[#232526] to-[#414345] flex flex-col">
        <h1 className="text-center text-white tracking-wider mt-8 font-sans text-3xl font-bold drop-shadow-lg">
          Code Scanner
        </h1>
        <main className="flex flex-1 items-center justify-center min-h-[60vh]">
          <Providers>{children}</Providers>
        </main>
      </body>
    </html>
  );
}
