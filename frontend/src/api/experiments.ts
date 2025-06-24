import { useExperimentStore } from "@/store/useExperimentStrore";
import api from "./axios";
import { Experiment } from "@/types/experiment";

export async function fetchAndStoreExperiments(): Promise<Experiment[]> {
    const res = await api.get("/experiments");
    const experiments = Array.isArray(res.data.body.experiments) ? res.data.body.experiments : [];
    useExperimentStore.getState().setExperiments(experiments);
    return experiments;
}

export async function createExperiment(payload: {
    name: string;
    description?: string;
}): Promise<Experiment> {
    const res = await api.post("/experiments", payload);
    const { experiment } = res.data.body;
    useExperimentStore.getState().addExperiment(experiment);
    return experiment as Experiment;
}