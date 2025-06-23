"use client";

import { useAuthStore } from "@/store/useAuthStore";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import Navbar from "@/components/Navbar";

export default function DashboardPage() {
  const { user } = useAuthStore();

  const mockStats = [
    { label: "Total Experiments", value: 3 },
    { label: "Total Variants", value: 7 },
    { label: "Total Conversions", value: 421 },
    { label: "Views This Week", value: 1230 },
  ];

  return (
    <>
      <Navbar />
      <main className="min-h-screen px-6 py-10 bg-background text-foreground">
        <div className="max-w-5xl mx-auto space-y-8">
          <section>
            <h1 className="text-3xl font-bold mb-1">
              Welcome back{user ? `, ${user.name}` : ""} 👋
            </h1>
            <p className="text-muted-foreground">
              Here’s what’s going on with your experiments.
            </p>
          </section>

          <section className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-4">
            {mockStats.map((stat) => (
              <Card key={stat.label}>
                <CardContent className="p-4">
                  <p className="text-sm text-muted-foreground">{stat.label}</p>
                  <p className="text-2xl font-bold">{stat.value}</p>
                </CardContent>
              </Card>
            ))}
          </section>

          <section>
            <div className="flex justify-between items-center mb-4">
              <h2 className="text-xl font-semibold">Recent Experiments</h2>
              <Link href="/dashboard/experiments">
                <Button size="sm">View All</Button>
              </Link>
            </div>
            <div className="border rounded-lg p-4 text-muted-foreground text-sm">
              No experiments yet.{" "}
              <Link href="/dashboard/experiments" className="underline">
                Create your first one
              </Link>
              .
            </div>
          </section>

          <section className="mt-10 bg-muted rounded-xl p-6 text-sm">
            <p className="font-medium mb-1">
              ⚙️ Don’t forget to integrate the SDK!
            </p>
            <p className="text-muted-foreground">
              Use our SDK to assign variants and track events in your frontend.{" "}
              <Link href="/docs/sdk" className="underline">
                View SDK docs
              </Link>
            </p>
          </section>
        </div>
      </main>
    </>
  );
}
