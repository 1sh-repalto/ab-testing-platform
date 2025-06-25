import Navbar from "@/components/Navbar";
import { Button } from "@/components/ui/button";
import { ArrowRight } from "lucide-react";
import Link from "next/link";

export default function Home() {
  return (
    <main className="min-h-screen bg-background text-foreground">
      <Navbar />
      <section className="relative w-[95%] h-160 mx-auto flex items-center justify-center text-center rounded-b-xl">
        <div className="absolute h-full w-full backdrop-blur-xl rounded-b-xl shadow-md py-12 bg-gradient-to-t from-violet-900 to-background" />
        {/* Content with glassmorphism */}
        <div className="relative z-10 px-6 max-w-3xl w-full">
          <h1 className="text-4xl md:text-6xl font-extrabold leading-tight mb-4 text-foreground/70">
            Build smarter with A/B Testing
          </h1>
          <p className="text-white/80 text-lg mb-10">
            Experiment with confidence and let real user data shape your
            product.
          </p>
          <Link href="/auth/signup">
            <Button size="lg" className="rounded-full p-7 text-md bg-foreground/80 hover:bg-foreground hover:scale-102 transition-transform duration-200 ease-out cursor-pointer">
              Get Started
              <ArrowRight />
            </Button>
          </Link>
        </div>
      </section>
      <div className="h-140 bg-blue-400 mt-100" />
    </main>
  );
}
