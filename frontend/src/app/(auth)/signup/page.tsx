"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export default function SignupPage() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log({ name, email, password });
    // TODO: call backend API
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-background text-foreground">
      <form
        onSubmit={handleSubmit}
        className="bg-card text-card-foreground p-6 rounded-lg shadow-md w-full max-w-sm"
      >
        <h2 className="text-2xl font-semibold mb-4 text-center">Sign Up</h2>
        <div className="space-y-3">
          <Input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Name"
            required
          />
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
          Sign Up
        </Button>
      </form>
    </div>
  );
}
