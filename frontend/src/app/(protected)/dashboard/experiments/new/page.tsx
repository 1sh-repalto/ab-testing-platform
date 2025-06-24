"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { createExperiment } from "@/api/experiments";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
// import { Textarea } from "@/components/ui/textarea";
import { toast } from "sonner";

export default function NewExperimentPage() {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [loading, setLoading] = useState(false);
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    try {
      await createExperiment({ name, description });
      toast.success("Experiment created successfully!");
      router.push(`/dashboard/experiments`); // go to variant setup
    } catch (err) {
      toast.error("Failed to create experiment");
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <main className="min-h-screen px-6 py-10 bg-background text-foreground">
      <div className="max-w-xl mx-auto">
        <h1 className="text-2xl font-bold mb-6">Create New Experiment</h1>

        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block mb-1 font-medium">Name</label>
            <Input
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="Experiment name"
              required
            />
          </div>

          <div>
            <label className="block mb-1 font-medium">Description</label>
            <textarea
              value={description}
              onChange={(e: any) => setDescription(e.target.value)}
              placeholder="Optional description"
            />
          </div>

          <Button type="submit" disabled={loading} onClick={handleSubmit}>
            {loading ? "Creating..." : "Create Experiment"}
          </Button>
        </form>
      </div>
    </main>
  );
}
