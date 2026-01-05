import { defaultShortcuts, ShortcutProps } from "@/styles/shortcuts";
import { PropsWithChildren } from "react";
import { TextProps, Text as ReactNativeText } from "react-native";

interface CustomTextProps extends PropsWithChildren, ShortcutProps, TextProps {
	fontSize?: number;
	bold?: boolean;
	underline?: boolean;
	color?: string;
}

export function Text({
	fontSize = 18,
	bold,
	underline,
	color,
	children,
	...restProps
}: CustomTextProps) {
	return (
		<ReactNativeText
			style={[
				defaultShortcuts(restProps),
				{
					fontSize,
					fontWeight: bold ? "bold" : "normal",
					textDecorationLine: underline ? "underline" : "none",
					color,
				},
			]}
			{...restProps}
		>
			{children}
		</ReactNativeText>
	);
}
