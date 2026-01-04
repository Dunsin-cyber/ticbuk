import { ComponentProps } from "react";
import { Ionicons } from "@expo/vector-icons";


export default function TabBarIcon({
	name,
	color,
}: {
	name: ComponentProps<typeof Ionicons>["name"];
	color: string;
}) {
	return (
		<Ionicons
			size={28}
			style={{ marginBottom: -3 }}
			name={name}
			color={color}
		/>
	);
}
