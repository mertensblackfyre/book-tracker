import React from "react";
import { YStack, Paragraph, XStack, Button, Input, Stack } from "tamagui";

const LoginScreen = ({ navigation }: any) => {
   return (
      <YStack
         marginTop="$15"
         borderRadius="$10"
         space
         paddingHorizontal="$7"
         paddingVertical="$6"
         width={"99%"}
         shadowColor={"#00000020"}
         shadowRadius={26}
         shadowOffset={{ width: 0, height: 4 }}
         backgroundColor="$background"
      >
         <Paragraph
            size="$5"
            fontWeight={"700"}
            opacity={0.8}
            marginBottom="$1"
         >
            Sign in to your account
         </Paragraph>
         <XStack space justifyContent={"space-evenly"} theme="light">
            <Button
               size="$5"
               hoverStyle={{ opacity: 0.8 }}
               focusStyle={{ scale: 0.95 }}
               borderColor="$gray8Light"
            ></Button>
            <Button
               size="$5"
               hoverStyle={{ opacity: 0.8 }}
               focusStyle={{ scale: 0.95 }}
               borderColor="$gray8Light"
            ></Button>
            <Button
               size="$5"
               hoverStyle={{ opacity: 0.8 }}
               focusStyle={{ scale: 0.95 }}
               borderColor="$gray8Light"
            ></Button>
         </XStack>
         <XStack
            alignItems="center"
            width="100%"
            justifyContent="space-between"
         >
            <Stack
               height="$0.25"
               backgroundColor="black"
               width="$10"
               opacity={0.1}
            />
            <Paragraph size="$3" opacity={0.5}>
               or
            </Paragraph>
            <Stack
               height="$0.25"
               backgroundColor="black"
               width="$10"
               opacity={0.1}
            />
         </XStack>

         {/* email sign up option */}
         <Input autoCapitalize="none" placeholder="Email" />
         <Input
            autoCapitalize="none"
            placeholder="Password"
            textContentType="password"
            secureTextEntry
         />

         {/* sign up button */}
         <Button
            themeInverse
            hoverStyle={{ opacity: 0.8 }}
            onHoverIn={() => {}}
            onHoverOut={() => {}}
            focusStyle={{ scale: 0.975 }}
         >
            Sign in
         </Button>

         {/* or sign in, in small and less opaque font */}
         <XStack>
            <Paragraph size="$2" marginRight="$2" opacity={0.4}>
               Don’t have an account?
            </Paragraph>
            <Paragraph
               onPress={() => navigation.navigate("Sign")}
               cursor={"pointer"}
               size="$2"
               fontWeight={"700"}
               opacity={0.5}
               hoverStyle={{ opacity: 0.4 }}
            >
               Sign up
            </Paragraph>
         </XStack>
      </YStack>
   );
};
export default LoginScreen;
