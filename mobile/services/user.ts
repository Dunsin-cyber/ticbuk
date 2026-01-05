import { AuthResponse } from "@/types/user";
import { Api } from "./api";

export type Credentials = {
	email: string;
	password: string;
};

export function login(credentials: Credentials): Promise<AuthResponse> {
	return Api.post("/auth/login", credentials);
}

export function register(credentials: Credentials): Promise<AuthResponse> {
	return Api.post("/auth/register", credentials);
}

const userService = {
	login,
	register,
};

export { userService };
