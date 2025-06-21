"use client";

import Link from "next/link";
import { ModeToggle } from "./mode-toggle";

export default function Navbar() {
  return (
    <nav className="flex items-center justify-between px-6 py-4 shadow-md">
      <Link href="/" className="text-xl font-bold">
        ABTest
      </Link>
      <div className="flex items-center gap-4">
        <Link href="/login" className="hover:underline">
          Login
        </Link>
        <Link href="/signup" className="hover:underline">
          Signup
        </Link>
        <ModeToggle />
      </div>
    </nav>
  );
}
