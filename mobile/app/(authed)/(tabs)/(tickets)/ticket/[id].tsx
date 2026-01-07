import { Text } from "@/components/Text";
import { VStack } from "@/components/VStack";
import { ticketService } from "@/services/ticket";
import { Ticket } from "@/types/ticket";
import { useNavigation } from "@react-navigation/native";
import { router, useFocusEffect, useLocalSearchParams } from "expo-router";
import { useCallback, useEffect, useState } from "react";
import { Alert, Image } from "react-native";

export default function TicketDetailScreen() {
	const navigation = useNavigation();
	const { id } = useLocalSearchParams();

	const [ticket, setTicket] = useState<Ticket | null>(null);
	const [qrcode, setQrcode] = useState<string | null>(null);

	async function fetchTicket() {
		try {
			const response = await ticketService.getOne(Number(id));
			setTicket(response.data.ticket);
			setQrcode(response.data.qrCode);
		} catch (err) {
			Alert.alert("Error", err.message);
			router.back();
		}
	}

	useFocusEffect(
		useCallback(() => {
			fetchTicket();
		}, [])
	);

	useEffect(() => {
		navigation.setOptions({
			headerTitle: "",
		});
	}, [navigation]);

	if (!ticket) return null;

	return (
		<VStack
			alignItems='center'
			m={20}
			p={20}
			gap={20}
			flex={1}
			style={{
				backgroundColor: "white",
				borderRadius: 20,
			}}
		>
			<Text fontSize={50} bold>
				{ticket.event.name}{" "}
			</Text>
			<Text fontSize={20} bold>
				{ticket.event.location}{" "}
			</Text>

			<Text fontSize={16} color='gray'>
				{new Date(ticket.event.date).toLocaleString()}
			</Text>
			<Image
				style={{
					borderRadius: 20,
				}}
				width={300}
				height={300}
				source={{ uri: `data:image/png;base64,${qrcode}` }}
			/>
		</VStack>
	);
}
