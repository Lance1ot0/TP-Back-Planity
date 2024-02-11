export type LoginResponse = {
  success: boolean;
  role?: string;
  token?: string;
  error?: string;
};