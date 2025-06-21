"use-client";

import { useTheme } from "next-themes";
import { Sun, Moon } from "lucide-react";
import { Switch } from "./ui/switch";
import { useEffect, useState } from "react";

export function ModeToggle() {
  const { setTheme, theme } = useTheme();
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) return null;

  const isDark = theme === "dark"

  const handleToggle = (checked: boolean) => {
    setTheme(checked ? "dark" : "light")
  }

  return (
    <div className="flex items-center gap-2">
      <Sun className="w-4 h-4" />
      <Switch checked={isDark} onCheckedChange={handleToggle} />
      <Moon className="w-4 h-4" />
    </div>
  );
}
