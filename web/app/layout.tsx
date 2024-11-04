import { LayoutHeader } from "@/src/components/LayoutHeader";
import type { Metadata } from "next";
import localFont from "next/font/local";
import "./globals.css";
import { LayoutFooter } from "@/src/components/LayoutFooter";

const geistSans = localFont({
  src: "../public/fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "../public/fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export const metadata: Metadata = {
  title: "F1 Go & Next.js app",
  description: "Check out f1 race results of the recent years",
};

type RootLayoutProps = Readonly<{
  children: React.ReactNode;
}>

export default function RootLayout(props: RootLayoutProps) {

  const { children } = props;

  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <LayoutHeader />
        {children}
        <LayoutFooter />
      </body>
    </html>
  );
}
