import axios from "axios";

export const client = axios.create({
  baseURL: process.env.NEXT_PUBLIC_SCAN_API_URL || "http://localhost:8080",
  headers: {
    "Content-Type": "application/json",
  },
});
