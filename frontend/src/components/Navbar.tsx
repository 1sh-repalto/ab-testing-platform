"use client";

import Link from "next/link";
import { useState } from "react";
import { Button } from "./ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import { Menu, ArrowUpRight } from "lucide-react";
import { useAuthStore } from "@/store/useAuthStore";
import { ModeToggle } from "./mode-toggle";
import { LogoutButton } from "./LogoutButton";

export default function Navbar() {
  const user = useAuthStore((state) => state.user);

  return (
    <nav className="fixed top-2 left-1/2 transform -translate-x-1/2 z-50 w-[97%] backdrop-blur-lg bg-white/10 shadow-lg rounded-xl px-6 py-4 flex items-center justify-between">
      {/* Logo */}
      <div className="flex justify-center items-center gap-14">
        <Link href="/" className="text-xl font-bold">
          ABTest
        </Link>
        <Link
          href="/docs"
          className="text-sm font-medium hover:underline flex items-center"
        >
          Docs
          <ArrowUpRight size={18} />
        </Link>
      </div>

      {/* Right-side links */}
      <div className="flex items-center gap-4">
        {/* Always show Docs link */}

        {!user ? (
          // If not logged in
          <>
            <Link href="/login">
              <Button variant="outline" size="sm">
                Login
              </Button>
            </Link>
            <Link href="/signup">
              <Button size="sm">Signup</Button>
            </Link>
          </>
        ) : (
          // If logged in: show menu icon
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button
                variant="ghost"
                className="p-0 w-12 h-12 flex items-center justify-center"
              >
                <Menu style={{ width: "24px", height: "24px" }} />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="w-40">
              <DropdownMenuItem asChild>
                <LogoutButton />
              </DropdownMenuItem>
              <DropdownMenuItem asChild>
                <ModeToggle />
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        )}
      </div>
    </nav>
  );
}
