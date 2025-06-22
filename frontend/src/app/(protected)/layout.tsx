"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { useAuthStore } from "@/store/useAuthStore";
import { fetchUser } from "@/api/auth";
import { toast } from "sonner";

export default function ProtectedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();
  const { user, setUser } = useAuthStore();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const checkSession = async () => {
      try {
        if(!user) {
          const userData = await fetchUser();
          setUser(userData);
        }
      } catch (error) {
        router.replace("/login");
        toast("Please log in to continue");
      } finally {
        setLoading(false)
      }
    };

    checkSession();
  }, [user]);

  if (loading)
    return <div className="p-6 text-center">Checking authentication...</div>;

  return <>{children}</>;
}
