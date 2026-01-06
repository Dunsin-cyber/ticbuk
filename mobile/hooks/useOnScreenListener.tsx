import { useNavigation } from "expo-router";
import { useEffect } from "react";

export function useOnScreenListener(
	eventType: "focus" | "blur" | "state" | "beforeRemove",
	callback: VoidFunction
) {
	const navigation = useNavigation();

	useEffect(() => {
		const subsribe = navigation.addListener(eventType, callback);
		return subsribe;
	}, [navigation, callback, eventType]);
}
