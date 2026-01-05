import { userService } from "@/services/user";
import { User } from "@/types/user";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { router } from "expo-router";
import {
	createContext,
	PropsWithChildren,
	useContext,
	useEffect,
	useState,
} from "react";

interface AuthContextProps {
	isLoggedIn: boolean;
	isLoadingAuth: boolean;
	authenticate: (
		authMode: "login" | "register",
		email: string,
		password: string
	) => Promise<void>;
	logout: VoidFunction;
	user: User | null;
}

const AuthContext = createContext({} as AuthContextProps);

// create custom useAuth hook
export function useAuth() {
	return useContext(AuthContext);
}

export function AuthenticationProvider({ children }: PropsWithChildren) {
	const [isLoggedIn, setIsLoggedIn] = useState(false);
	const [isLoadingAuth, setIsLoadingAuth] = useState(false);

	const [user, setUser] = useState<User | null>(null);

	useEffect(() => {
		async function checkedIfLoggedIn() {
			const token = await AsyncStorage.getItem("token");
			const user = await AsyncStorage.getItem("user");

			if (token && user) {
				setIsLoggedIn(true);
				setUser(JSON.parse(user));
				router.replace("(authed)");
			} else {
				setIsLoggedIn(false);
			}
		}

		checkedIfLoggedIn();
	}, []);

	async function authenticate(
		authMode: "login" | "register",
		email: string,
		password: string
	): Promise<void> {
		try {
			console.log("got here");
			setIsLoadingAuth(true);

			const response = await userService[authMode]({ email, password });

			if (response) {
				const { data } = response;
				const { user, token } = data;

				await AsyncStorage.setItem("token", token);
				await AsyncStorage.setItem("user", JSON.stringify(user));
				setUser(user);
				router.replace("(authed)");
				setIsLoggedIn(true);
			}
		} catch (error) {
			console.log(error);
			setIsLoggedIn(false);
		} finally {
			setIsLoadingAuth(false);
		}
	}

	async function logout() {
		setIsLoggedIn(false);
		await AsyncStorage.removeItem("token");
		await AsyncStorage.removeItem("user");
	}

	return (
		<AuthContext.Provider
			value={{ authenticate, logout, isLoggedIn, isLoadingAuth, user }}
		>
			{children}
		</AuthContext.Provider>
	);
}
