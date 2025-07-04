import { create } from "zustand";
import { persist } from "zustand/middleware";

interface User{
    id: string;
    name: string;
    email: string;
    createdAt: Date;
}

interface AuthState {
    user: User | null;
    setUser: (user: User) => void;
    clearUser: () => void;
}

export const useAuthStore = create<AuthState>() (
    persist(
        (set) => ({
            user: null,
            setUser: (user) => set({ user }),
            clearUser: () => set({ user: null }),
        }),
        {
            name: "auth-store",
            partialize: (state) => ({ user: state.user })
        }
    )
);