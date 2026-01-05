import { HStack } from "@/components/HStack";
import { VStack } from "@/components/VStack";
import { KeyboardAvoidingView, ScrollView } from "react-native";
import { Text } from "@/components/Text";
import TabBarIcon from "@/components/navigation/TabBarIcon";
import { useState } from "react";
import { Input } from "@/components/Input";
import { Button } from "@/components/Button";
import { Divider } from "@/components/Divider";
import { useAuth } from "@/context/AuthContext";

export default function Login() {
	const { authenticate, isLoadingAuth } = useAuth();

	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const [authMode, setAuthMode] = useState<"login" | "register">("login");

	function onToggleAuthMode() {
		setAuthMode(authMode === "login" ? "register" : "login");
	}

	async function onAuthenticate() {
		await authenticate(authMode, email, password);
	}

	return (
		<KeyboardAvoidingView behavior='padding' style={{ flex: 1 }}>
			<ScrollView contentContainerStyle={{ flex: 1 }}>
				<VStack
					flex={1}
					justifyContent='center'
					alignItems='center'
					p={40}
					gap={40}
				>
					<HStack gap={10}>
						<Text fontSize={30} bold mb={20}>
							Ticket Booking
						</Text>
						<TabBarIcon name='ticket' size={40} />
					</HStack>

					<VStack w={"100%"} gap={30}>
						<VStack gap={5}>
							<Text ml={10} fontSize={14} color='gray'>
								Email
							</Text>

							<Input
								value={email}
								onChangeText={setEmail}
								placeholder='Email'
								placeholderTextColor='darkgray'
								autoCapitalize='none'
								autoCorrect={false}
								h={48}
								p={14}
							/>
						</VStack>

						<VStack gap={5}>
							<Text ml={10} fontSize={14} color='gray'>
								Password
							</Text>

							<Input
								secureTextEntry
								value={password}
								onChangeText={setPassword}
								placeholder='Password'
								placeholderTextColor='darkgray'
								autoCapitalize='none'
								autoCorrect={false}
								h={48}
								p={14}
							/>
						</VStack>

						<Button
							isLoading={isLoadingAuth} //TODO: finish once we hook up api
							onPress={onAuthenticate} //TODO: finish once we hook up api
						>
							{authMode === "login" ? "Login" : "Register"}
						</Button>
					</VStack>

					<Divider w={"90%"} />

					<Text onPress={onToggleAuthMode} fontSize={16} underline>
						{authMode === "login" ? "Register new account" : "Login to account"}
					</Text>
				</VStack>
			</ScrollView>
		</KeyboardAvoidingView>
	);
}
