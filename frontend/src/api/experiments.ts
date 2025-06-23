import api from "./axios";
import { Experiment } from "@/types/experiment";

export async function fetchExperiments(): Promise<Experiment[]> {
    const res = await api.get("/experiments");
    const { experiments } = res.data.body;
    return Array.isArray(experiments) ? experiments : [];
}