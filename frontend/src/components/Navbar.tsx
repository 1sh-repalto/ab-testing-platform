"use client"

import { LogoutButton } from "./LogoutButton";
import { useAuthStore } from "@/store/useAuthStore";
import { ModeToggle } from "./mode-toggle";
import Link from "next/link";

export default function Navbar() {
  const user = useAuthStore((state) => state.user);

  return (
    <nav className="flex items-center justify-between px-6 py-4 shadow-md">
      <Link href="/" className="text-xl font-bold">ABTest</Link>
      <div className="flex items-center gap-4">
        {user ? (
          <>
            <span className="text-sm">Hi, {user.name}</span>
            <LogoutButton />
          </>
        ) : (
          <>
            <Link href="/login">Login</Link>
            <Link href="/signup">Signup</Link>
          </>
        )}
        <ModeToggle />
      </div>
    </nav>
  );
}
