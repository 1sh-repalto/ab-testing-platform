"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import axios from "@/api/axios";
import { toast } from "sonner";

export default function ProtectedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const checkAuth = async () => {
      try {
        await axios.get("/auth/me");
        setLoading(false);
      } catch (error) {
        router.replace("/login");
        toast("Please log in to continue");
      }
    };

    checkAuth();
  }, [router]);

  if (loading)
    return <div className="p-6 text-center">Checking authentication...</div>;

  return <>{children}</>;
}
