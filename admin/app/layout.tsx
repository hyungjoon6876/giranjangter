import type { Metadata } from "next";
import { Providers } from "@/lib/providers";
import { AdminShell } from "@/components/layout/admin-shell";
import "./globals.css";

export const metadata: Metadata = {
  title: "기란장터 Admin",
  description: "기란장터 관리자 대시보드",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ko">
      <head>
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link rel="preconnect" href="https://fonts.gstatic.com" crossOrigin="anonymous" />
      </head>
      <body>
        <Providers>
          <AdminShell>{children}</AdminShell>
        </Providers>
      </body>
    </html>
  );
}
