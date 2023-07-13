import React, { useEffect, useState } from 'react';
import { createStackNavigator } from '@react-navigation/stack';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';

import RepoList from '../src/RepoList';
import Info from '../src/Info';
import User from '../src/User';


const Stack = createStackNavigator();
const Tab = createBottomTabNavigator();

function AttendeeTabs() {
    const [username, setUsername] = useState('');
    return (
        <Tab.Navigator>
            <Tab.Screen name="User">
                {() => <User setUsername={setUsername} />}
            </Tab.Screen>
            <Tab.Screen name="List">
                {() => <RepoList username={username} />}
            </Tab.Screen>
            <Tab.Screen name="Info">
                {() => <Info username={username} />}
            </Tab.Screen>
        </Tab.Navigator>
    );
}



export default function AppNavigator() {

    return (
        <NavigationContainer>
            <Stack.Navigator>
                <Stack.Screen name="App" component={AttendeeTabs} options={{ headerShown: false }} />
            </Stack.Navigator>
        </NavigationContainer>
    );
}
