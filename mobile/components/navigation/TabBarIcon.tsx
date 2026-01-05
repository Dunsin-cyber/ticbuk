import { ComponentProps } from "react";
import { Ionicons } from "@expo/vector-icons";

export default function TabBarIcon({
	name,
	size = 28,
	color = "black",
}: {
	name: ComponentProps<typeof Ionicons>["name"];
	color?: string;
	size: number;
}) {
	return (
		<Ionicons
			size={size}
			style={{ marginBottom: -3 }}
			name={name}
			color={color}
		/>
	);
}
