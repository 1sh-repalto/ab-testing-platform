"use client"

import { useEffect, useState } from "react"
import Link from "next/link"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { fetchExperiments } from "@/api/experiments"
import { Experiment } from "@/types/experiment"

export default function ExperimentsPage() {
  const [experiments, setExperiments] = useState<Experiment[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const load = async () => {
      try {
        const data = await fetchExperiments()     
        setExperiments(data)
      } catch (err) {
        console.error("Failed to fetch experiments", err)
      } finally {
        setLoading(false)
      }
    }

    load()
  }, [])

  return (
    <main className="min-h-screen px-6 py-10 bg-background text-foreground">
      <div className="max-w-5xl mx-auto space-y-6">
        <div className="flex justify-between items-center">
          <h1 className="text-2xl font-bold">Your Experiments</h1>
          <Button asChild>
            <Link href="/dashboard/experiments/new">Create New</Link>
          </Button>
        </div>

        {loading ? (
          <div>Loading experiments...</div>
        ) : experiments.length === 0 ? (
          <div className="text-muted-foreground">No experiments found.</div>
        ) : (
          <div className="grid gap-4">
            {experiments.map((exp) => (
              <Card key={exp.id}>
                <CardContent className="p-4">
                  <h2 className="text-lg font-semibold">{exp.name}</h2>
                  <p className="text-sm text-muted-foreground">
                    {exp.description || "No description."}
                  </p>
                  <p className="text-xs mt-2 text-muted-foreground">
                    Created on {new Date(exp.created_at).toLocaleDateString()}
                  </p>
                </CardContent>
              </Card>
            ))}
          </div>
        )}
      </div>
    </main>
  )
}
