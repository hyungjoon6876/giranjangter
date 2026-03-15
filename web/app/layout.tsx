import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "기란장터",
  description: "리니지 클래식 아이템 거래 중개 플랫폼",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ko">
      <body className="antialiased">{children}</body>
    </html>
  );
}
