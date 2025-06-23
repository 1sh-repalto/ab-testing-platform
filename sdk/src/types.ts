export interface Variant {
    id: number;
    name: string;
    weight: number;
    payload: Record<string, any>;
    experiment_id: number;
}

export interface VariantResponse {
    variant: Variant;
}