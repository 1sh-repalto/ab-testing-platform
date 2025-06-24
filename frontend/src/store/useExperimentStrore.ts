import { create } from "zustand";
import { Experiment } from "@/types/experiment";

interface ExperimentState {
    experiments: Experiment[];
    setExperiments: (exps: Experiment[]) => void;
    addExperiment: (exp: Experiment) => void;
    clearExperiments: () => void; 
}

export const useExperimentStore = create<ExperimentState>()((set) => ({
    experiments: [],
    setExperiments: (exps) => set({ experiments: exps }),
    addExperiment: (exp) => set((state) => ({ experiments: [exp, ...state.experiments] })),
    clearExperiments: () => set({ experiments: [] }),
}));