import type { Metadata } from "next";
import { Providers } from "@/lib/providers";
import { ResponsiveShell } from "@/components/layout/responsive-shell";
import "./globals.css";

export const metadata: Metadata = {
  title: "기란JT — 리니지 클래식 거래",
  description: "리니지 클래식 아이템 거래 중개 플랫폼",
  openGraph: {
    title: "기란JT — 리니지 클래식 거래 플랫폼",
    description: "무료 커뮤니티 기반 리니지 클래식 아이템 거래 중개",
    images: ["/images/og-image.png"],
  },
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ko">
      <head>
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link rel="preconnect" href="https://fonts.gstatic.com" crossOrigin="anonymous" />
        <script src="https://accounts.google.com/gsi/client" async defer />
      </head>
      <body>
        <Providers>
          <ResponsiveShell>{children}</ResponsiveShell>
        </Providers>
      </body>
    </html>
  );
}
