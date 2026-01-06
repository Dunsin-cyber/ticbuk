import { VStack } from "@/components/VStack";
import { eventService } from "@/services/event";
import { router, useNavigation } from "expo-router";
import { useEffect, useState } from "react";
import { Alert } from "react-native";
import { Text } from "@/components/Text";
import { Input } from "@/components/Input";
import { Button } from "@/components/Button";
import DateTimePicker from "@/components/DateTimePicker";

export default function NewEvent() {
	const [name, setName] = useState("");
	const [location, setLocation] = useState("");
	const [date, setDate] = useState(new Date());

	const [isSubmitting, setIsSubitting] = useState(false);
	const navigation = useNavigation();

	async function onSubmit() {
		try {
			setIsSubitting(true);

			await eventService.createOne(name, location, date.toISOString());
			router.back();
		} catch (err) {
			console.error(err);
			Alert.alert("Error", "could not create the event");
		} finally {
			setIsSubitting(false);
		}
	}

	function onchangeDate(date?: Date) {
		setDate(date || new Date());
	}

	useEffect(() => {
		navigation.setOptions({
			headerTitle: "New Event",
		});
	}, [navigation]);

	return (
		<VStack m={20} flex={1} gap={30}>
			<VStack gap={5}>
				<Text ml={10} fontSize={14} color='gray'>
					Name
				</Text>

				<Input
					value={name}
					onChangeText={setName}
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
					value={location}
					onChangeText={setLocation}
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
				<DateTimePicker onChange={onchangeDate} currentDate={date} />
			</VStack>

			<Button
				mt={"auto"}
				isLoading={isSubmitting}
				disabled={isSubmitting}
				onPress={onSubmit}
			>
				Save
			</Button>
		</VStack>
	);
}
