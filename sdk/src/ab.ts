import axios from "axios";
import { VariantResponse } from "./types";
import dotenv from "dotenv";

dotenv.config();
const API_BASE_URL = process.env.SDK_API_URL || "http://localhost:8080/api";

export async function initAB(experimentId: number, user: { userIdentifier: string }) {
  const res = await axios.post<VariantResponse>(
    `${API_BASE_URL}/experiments/assign`,
    {
      experiment_id: experimentId,
      user_identifier: user.userIdentifier,
    },
    { withCredentials: true }
  );

  const { variant } = res.data;

  function track(event: "conversion" | "view") {
    return axios.post(
      `${API_BASE_URL}/events`,
      {
        experiment_id: experimentId,
        variant_id: variant.id,
        event_type: event,
        user_identifier: user.userIdentifier,
      },
      { withCredentials: true }
    );
  }

  return { variant: variant.payload, track };
}
