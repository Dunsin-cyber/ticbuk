import { HStack } from "@/components/HStack";
import { VStack } from "@/components/VStack";
import { Text } from "@/components/Text";
import { eventService } from "@/services/event";
import { Event } from "@/types/event";
import { useCallback, useEffect, useState } from "react";
import { Alert, FlatList, TouchableOpacity } from "react-native";
import { useAuth } from "@/context/AuthContext";
import { UserRole } from "@/types/user";
import { router, useNavigation } from "expo-router";
import TabBarIcon from "@/components/navigation/TabBarIcon";
import { Divider } from "@/components/Divider";
import { Button } from "@/components/Button";
import { useFocusEffect } from "@react-navigation/native";

export default function EventsScreen() {
	const [isLoading, setIsLoading] = useState(false);
	const [events, setEvents] = useState<Event[]>([]);
	const navigation = useNavigation();

	const { user } = useAuth();

	const fetchEvents = async () => {
		try {
			setIsLoading(true);
			const response = await eventService.getAll();
			if (response) {
				setEvents(response.data);
			}
		} catch (err) {
			console.error(err);
			Alert.alert("Error", err.message);
		} finally {
			setIsLoading(false);
		}
	};

	// useOnScreenListener("focus", fetchEvents);

	useFocusEffect(
		useCallback(() => {
			fetchEvents();
		}, [])
	);

	useEffect(() => {
		fetchEvents();

		navigation.setOptions({
			headerTitle: "Events",
			headerRight: user?.role === UserRole.Manager ? headerRight : null,
		});
	}, [navigation, user?.role]);

	function onGoToEventPage(id: number) {
		if (user?.role === UserRole.Manager) {
			router.push(`/(events)/event/${id}`);
		}
	}

	function buyTicket(id: number) {
		try {
			// await ticketService.createOne(id)
			Alert.alert("Success", "Ticket purchased successfully");
		} catch (error) {
			console.error(error);
			Alert.alert("Error", "Failed to buy ticket");
		}
	}

	return (
		<VStack flex={1} p={20} pb={0} gap={20}>
			<HStack alignItems='center' justifyContent='center'>
				<Text fontSize={18} bold>
					{events.length} Events
				</Text>
			</HStack>

			<FlatList
				data={events}
				keyExtractor={({ id }) => id.toString()}
				onRefresh={fetchEvents}
				refreshing={isLoading}
				ItemSeparatorComponent={() => <VStack h={20} />}
				renderItem={({ item: event }) => (
					<VStack
						gap={20}
						p={20}
						style={{
							backgroundColor: "white",
							borderRadius: 20,
						}}
						key={event.id}
					>
						<TouchableOpacity onPress={() => onGoToEventPage(event.id)}>
							<HStack alignItems='center' justifyContent='space-between'>
								<HStack alignItems='center'>
									<Text fontSize={26} bold>
										{event.name}
									</Text>
									<Text fontSize={26} bold>
										{" "}
										|{" "}
									</Text>
									<Text fontSize={16} bold>
										{event.location}
									</Text>
								</HStack>
								{user?.role === UserRole.Manager && (
									<TabBarIcon size={24} name='chevron-forward' />
								)}
							</HStack>
						</TouchableOpacity>

						<Divider />

						<HStack justifyContent='space-between'>
							<Text bold fontSize={16} color='gray'>
								Sold : {event.totalTicketsPurchased}
							</Text>
							<Text bold fontSize={16} color='green'>
								Entered : {event.totalTicketsEntered}
							</Text>
						</HStack>

						{user?.role === UserRole.Attendee && (
							<VStack>
								<Button
									variant='outlined'
									disabled={isLoading}
									onPress={() => buyTicket(event.id)} // TODO: finish once there is ticket services
								>
									Buy Ticket
								</Button>
							</VStack>
						)}

						<Text fontSize={13} color='gray'>
							{event?.date}
						</Text>
					</VStack>
				)}
			/>
		</VStack>
	);
}

const headerRight = () => {
	return (
		<TabBarIcon
			size={32}
			name='add-circle-outline'
			onPress={() => router.push("/(events)/new")}
		/>
	);
};
