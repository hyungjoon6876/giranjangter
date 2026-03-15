import type { Metadata } from "next";
import { Providers } from "@/lib/providers";
import { ResponsiveShell } from "@/components/layout/responsive-shell";
import "./globals.css";

export const metadata: Metadata = {
  title: "기란장터 — 리니지 클래식 거래",
  description: "리니지 클래식 아이템 거래 중개 플랫폼",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ko">
      <head>
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link rel="preconnect" href="https://fonts.gstatic.com" crossOrigin="anonymous" />
      </head>
      <body>
        <Providers>
          <ResponsiveShell>{children}</ResponsiveShell>
        </Providers>
      </body>
    </html>
  );
}
