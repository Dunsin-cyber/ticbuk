import { Stack, StackProps } from "./Stack";

// eslint-disable-next-line @typescript-eslint/no-empty-object-type
interface VStackProps extends StackProps {}

export function VStack(props: VStackProps) {
	return <Stack {...props} direction='column' />;
}
