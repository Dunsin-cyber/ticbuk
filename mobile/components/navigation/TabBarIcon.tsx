import { ComponentProps } from "react";
import { Ionicons } from "@expo/vector-icons";

export default function TabBarIcon({
	name,
	size = 28,
	color = "black",
	onPress,
}: {
	name: ComponentProps<typeof Ionicons>["name"];
	color?: string;
	size: number;
	onPress?: VoidFunction;
}) {
	return (
		<Ionicons
			onPress={onPress}
			size={size}
			style={{ marginBottom: -3 }}
			name={name}
			color={color}
		/>
	);
}
