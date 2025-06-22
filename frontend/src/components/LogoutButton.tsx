"use client";

import { logoutUser } from "@/api/auth";
import { useAuthStore } from "@/store/useAuthStore";
import { useRouter } from "next/navigation";
import { Button } from "./ui/button";
import { toast } from "sonner";

export function LogoutButton() {
    const router = useRouter();

    const handleLogout = async () => {
        try {
            await logoutUser();
            useAuthStore.getState().clearUser();
            toast.success("Logged out successfully");
            router.push("/")
        } catch (error) {
            toast.error("Logout failed")
        }
    };

    return (
        <Button variant="outline" onClick={handleLogout}>
            Log Out
        </Button>
    )
}