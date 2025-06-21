import Navbar from "@/components/Navbar";
import { Button } from "@/components/ui/button";

export default function Home() {
  return (
    <main className="min-h-screen bg-background text-foreground">
      <Navbar />
      <section className="flex flex-col items-center justify-center h-[80vh] px-6 text-center">
        <h1 className="text-4xl font-bold mb-4">
          Run A/B Tests with Confidence
        </h1>
        <p className="text-lg text-muted-foreground max-w-xl mb-6">
          Create experiments, assign variants, and measure what matters — all in
          one place.
        </p>
        <Button asChild>
          <a href="/signup">Get Started</a>
        </Button>
      </section>
    </main>
  );
}
