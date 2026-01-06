import { HStack } from "@/components/HStack";
import { Text } from "@/components/Text";
import { VStack } from "@/components/VStack";
import { useAuth } from "@/context/AuthContext";
import { Alert } from "react-native";

export default function SettingsScreen() {
	const { logout } = useAuth();

	async function onLogout() {
		try {
			await logout();
		} catch (error) {
			console.error(error);
			Alert.alert("Error", "could not logout");
		}
	}
	return (
		<VStack my='auto'>
			<HStack alignItems='center' justifyContent='center'>
				<Text fontSize={30} bold underline onPress={onLogout}>
					LOGOUT
				</Text>
			</HStack>
		</VStack>
	);
}
