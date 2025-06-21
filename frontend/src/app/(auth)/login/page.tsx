"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log({ email, password });
    // TODO: call backend API
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-background text-foreground">
      <form
        onSubmit={handleSubmit}
        className="bg-card text-card-foreground p-6 rounded-lg shadow-md w-full max-w-sm"
      >
        <h2 className="text-2xl font-semibold mb-4 text-center">Log In</h2>
        <div className="space-y-3">
          <Input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Email"
            required
          />
          <Input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Password"
            required
          />
        </div>
        <Button type="submit" className="w-full mt-4">
          Log In
        </Button>
      </form>
    </div>
  );
}
