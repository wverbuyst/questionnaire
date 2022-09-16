import axios from "axios";
import { Questionnaire } from "./interfaces";

const BACKEND_URL = "http://localhost:9090";

export const api = {
  getQuestionnaire: async (): Promise<Questionnaire> => {
    const response = await axios.get(`${BACKEND_URL}/questionnaire`);
    console.log(response);
    return response.data;
  },
};
