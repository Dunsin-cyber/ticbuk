import { Stack, StackProps } from "./Stack";

// eslint-disable-next-line @typescript-eslint/no-empty-object-type
interface HStackProps extends StackProps {}


export function HStack(props: HStackProps) {
  return (
    <Stack {...props} direction="row" />
  )

}