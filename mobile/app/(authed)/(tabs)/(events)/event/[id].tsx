import { Button } from "@/components/Button";
import { Input } from "@/components/Input";
import { Text } from "@/components/Text";
import TabBarIcon from "@/components/navigation/TabBarIcon";
import { VStack } from "@/components/VStack";
import { eventService } from "@/services/event";
import { Event } from "@/types/event";
import { router, useLocalSearchParams, useNavigation } from "expo-router";
import { useState, useCallback, useEffect } from "react";
import { Alert } from "react-native";
import DateTimePicker from "@/components/DateTimePicker";
import { useFocusEffect } from "@react-navigation/native";

export default function EventDetailScreen() {
	const navigation = useNavigation();
	const { id } = useLocalSearchParams();

	const [isSubmitting, setIsSubitting] = useState(false);
	const [eventData, setEventData] = useState<Event | null>(null);

	function updatefield(field: keyof Event, value: string | Date) {
		setEventData((prev) => ({
			...prev!,
			[field]: value,
		}));
	}

	const onDelete = useCallback(async () => {
		if (!eventData) return;
		try {
			Alert.alert(
				"Delete Event",
				"Are you sure you want to delete this event?",
				[
					{ text: "Cancel" },
					{
						text: "Delete",
						onPress: async () => {
							await eventService.deleteOne(Number(id));
							router.back();
						},
					},
				]
			);
		} catch {
			Alert.alert("Error", "failed to delete event");
		}
	}, [eventData, id]);

	async function onSubmitChanges() {
		if (!eventData) return;
		try {
			setIsSubitting(true);
			await eventService.updateOne(
				Number(id),
				eventData.name,
				eventData.location,
				eventData.date
			);

			router.back();
		} catch (error) {
			console.error(error);
			Alert.alert("Error", "Failed to save changes");
		} finally {
			setIsSubitting(false);
		}
	}

	const fetchEvent = async () => {
		try {
			const response = await eventService.getOne(Number(id));
			setEventData(response.data);
		} catch (err) {
			console.error(err);
			router.back();
		}
	};

	// useOnScreenListener("focus", fetchEvent);
	
		useFocusEffect(
			useCallback(() => {
				fetchEvent();
			}, [])
		);

	useEffect(() => {
		navigation.setOptions({
			headerTitle: "Edit",
			headerRight: () => headerRight(onDelete),
		});
	}, [navigation, onDelete]);

	return (
		<VStack m={20} flex={1} gap={30}>
			<VStack gap={5}>
				<Text ml={10} fontSize={14} color='gray'>
					Name
				</Text>

				<Input
					value={eventData?.name}
					onChangeText={(value) => updatefield("name", value)}
					placeholder='Name'
					placeholderTextColor='darkgray'
					h={48}
					p={14}
				/>
			</VStack>

			<VStack gap={5}>
				<Text ml={10} fontSize={14} color='gray'>
					Location
				</Text>

				<Input
					value={eventData?.location}
					onChangeText={(value) => updatefield("location", value)}
					placeholder='Location'
					placeholderTextColor='darkgray'
					h={48}
					p={14}
				/>
			</VStack>

			{/* DateTimePicker */}
			<VStack gap={5}>
				<Text ml={10} fontSize={14} color='gray'>
					Date
				</Text>
				<DateTimePicker
					onChange={(date) => updatefield("date", date || new Date())}
					currentDate={new Date(eventData?.date || new Date())}
				/>
			</VStack>

			<Button
				mt={"auto"}
				isLoading={isSubmitting}
				disabled={isSubmitting}
				onPress={onSubmitChanges}
			>
				Save Changes
			</Button>
		</VStack>
	);
}

const headerRight = (onPress: VoidFunction) => {
	return <TabBarIcon size={30} name='trash' onPress={onPress} />;
};
