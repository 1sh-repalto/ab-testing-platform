import axios from "axios";
import { VariantResponse } from "./types";

const API_BASE_URL = "http://localhost:8080/api";

export async function initAB(experimentKey: string, user: { userId: string }) {
  const res = await axios.post<VariantResponse>(
    `${API_BASE_URL}/experiments/assign`,
    {
      experimentKey,
      userId: user.userId,
    },
    { withCredentials: true }
  );

  const variant = res.data;

  function track(event: "conversion" | "view") {
    return axios.post(
      `${API_BASE_URL}/events`,
      {
        experimentKey,
        variantId: variant.id,
        eventType: event,
        userId: user.userId,
      },
      { withCredentials: true }
    );
  }

  return { variant: variant.payload, track };
}
