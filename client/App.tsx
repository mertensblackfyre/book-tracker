import { useFonts } from "expo-font";

import { StatusBar } from "expo-status-bar";

import { useColorScheme } from "react-native";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { Paragraph, Spacer, TamaguiProvider, Theme, YStack } from "tamagui";
import config from "./tamagui.config";
import React from "react";
import LoginScreen from "./screens/LoginScreen";
import SignScreen from "./screens/SignUp";
export default function App() {
   const Stack = createNativeStackNavigator();
   const colorScheme = useColorScheme();
   const [loaded] = useFonts({
      Inter: require("@tamagui/font-inter/otf/Inter-Medium.otf"),

      InterBold: require("@tamagui/font-inter/otf/Inter-Bold.otf"),
   });
   if (!loaded) {
      return null;
   }
   return (
      <NavigationContainer>
         <TamaguiProvider config={config}>
            <Theme name={"light"}>
               <Stack.Navigator>
                  <Stack.Screen
                     options={{ headerShown: false }}
                     name="Login"
                     component={LoginScreen}
                  />

                  <Stack.Screen
                     options={{ headerShown: false }}
                     name="Sign"
                     component={SignScreen}
                  />
               </Stack.Navigator>

               <StatusBar style="auto" />
            </Theme>
         </TamaguiProvider>
      </NavigationContainer>
   );
}
