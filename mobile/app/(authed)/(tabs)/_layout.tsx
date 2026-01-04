import { Tabs } from "expo-router";
import { Text } from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { ComponentProps } from "react";
import TabBarIcon from "@/components/navigation/TabBarIcon";


export default function TabLayout() {
	const tabs = [
		{
			showFor: [],
			name: "(events)",
			displayName: "Events",
			icon: "calendar",
			options: {
				headerShown: false,
			},
		},
		{
			showFor: [],
			name: "(tickets)",
			displayName: "My Tickets",
			icon: "ticket",
			options: {
				headerShown: false,
			},
		},
		{
			showFor: [],
			name: "scan-ticket",
			displayName: "Scan Ticket",
			icon: "scan",
			options: {
				headerShown: true,
			},
		},
		{
			showFor: [],
			name: "settings",
			displayName: "Settings",
			icon: "cog",
			options: {
				headerShown: true,
			},
		},
	];
	return (
		<Tabs>
			{tabs.map((tab) => (
				<Tabs.Screen
					key={tab.name}
					name={tab.name}
					options={{
						...tab.options,
						headerTitle: tab.displayName,
						// href: tab.showFor.includes()
						tabBarLabel: ({ focused }) =>  (
							<Text
								style={{
									color: focused ? "black" : "gray",
									fontSize: 12,
									marginBottom: 5,
								}}
							>
								{tab.displayName}
							</Text>
						),
						tabBarIcon: ({ focused }) => (
							<TabBarIcon
								name={tab.icon as ComponentProps<typeof Ionicons>["name"]}
								color={focused ? "black" : "gray"}
							/>
						),
					}}
				/>
			))}
		</Tabs>
	);
}
