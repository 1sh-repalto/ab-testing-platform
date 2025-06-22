"use client";

import { useAuthStore } from "@/store/useAuthStore";

export default function DashboardPage() {
  const user = useAuthStore.getState().user;
  
  return (
    <div className="min-h-screen bg-background text-foreground flex items-center justify-center">
      <h1 className="text-3xl font-bold">Welcome to your dashboard 🎉</h1>
      <h3>{user?.name}</h3>
    </div>
  );
}
