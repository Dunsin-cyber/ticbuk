import { Button } from "@/components/Button";
import { Text } from "@/components/Text";
import { VStack } from "@/components/VStack";
import { ActivityIndicator, Vibration, Alert } from "react-native";
import {
	BarcodeScanningResult,
	CameraView,
	useCameraPermissions,
} from "expo-camera";
import { useState } from "react";
import { ticketService } from "@/services/ticket";

export default function ScanTicketScreen() {
	const [permission, requestPermission] = useCameraPermissions();
	const [scanningEnabled, setScanningEnabled] = useState(true);

	if (!permission) {
		return (
			<VStack flex={1} justifyContent='center' alignItems='center'>
				<ActivityIndicator size='large' />
			</VStack>
		);
	}

	if (!permission.granted) {
		return (
			<VStack gap={20} flex={1} justifyContent='center' alignItems='center'>
				<Text>Camera access is required to scan tickets</Text>
				<Button onPress={requestPermission}>Allow Camera Access</Button>
			</VStack>
		);
	}

	async function onBarcodeScanned({ data }: BarcodeScanningResult) {
		if (!scanningEnabled) return;

		try {
			Vibration.vibrate();
			setScanningEnabled(false);

			const [ticket, owner] = data.split(",");
			const ticketId = parseInt(ticket.split(":")[1]);
			const ownerId = parseInt(owner.split(":")[1]);

			await ticketService.validateOne(ticketId, ownerId);

			Alert.alert("Success", "Ticket Validated", [
				{text: "Ok", onPress: () => setScanningEnabled(true)}
			])
		} catch (err) {
			Alert.alert("Error", "Faiiled to validate ticket, please try again");
			setScanningEnabled(true);
		}
	}

	return (
		<CameraView
			style={{ flex: 1 }}
			facing='back'
			onBarcodeScanned={onBarcodeScanned}
			barcodeScannerSettings={{
				barcodeTypes: ["qr"],
			}}
		/>
	);
}
