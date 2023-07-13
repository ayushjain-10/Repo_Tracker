import React, { useState, useEffect } from 'react';
import { TextInput, TouchableOpacity, View, StyleSheet, KeyboardAvoidingView, Text, ScrollView } from 'react-native';
import { initializeApp } from 'firebase/app';
import { getAuth } from 'firebase/auth';
import { getFirestore, collection, addDoc, onSnapshot, query } from 'firebase/firestore';

const firebaseConfig = {
    apiKey: "AIzaSyANRsw-hBkz7T__fwOvhO-rRcUdjXApXzA",
    authDomain: "repotracker-2cba5.firebaseapp.com",
    projectId: "repotracker-2cba5",
    storageBucket: "repotracker-2cba5.appspot.com",
    messagingSenderId: "478035371720",
    appId: "1:478035371720:web:ee10fb8f646942f48827fd"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);

// Get instances of the services
const auth = getAuth();
const db = getFirestore(app);

const User = ({ setUsername }) => {
    const [inputUsername, setInputUsername] = useState('');
    const [currentUsername, setCurrentUsername] = useState('');
    const [userList, setUserList] = useState([]);

    const handleSubmit = async () => {
        try {
            await addDoc(collection(db, 'users'), { name: inputUsername });
            setCurrentUsername(inputUsername); // set current username
            setUsername(inputUsername);
            setInputUsername(''); // clear the input field
        } catch (e) {
            console.error("Error adding document: ", e);
        }
    };

    const handleUserPress = (username) => {
        setUsername(username);
    };

    useEffect(() => {
        const q = query(collection(db, 'users'));
        const unsubscribe = onSnapshot(q, snapshot => {
            const users = snapshot.docs.map(doc => doc.data().name);
            setUserList(users);
        });

        return () => unsubscribe(); // unsubscribe from the event on cleanup
    }, []);


    return (
        <KeyboardAvoidingView style={styles.container} behavior={Platform.OS === 'ios' ? 'padding' : null}>
            <TextInput
                style={styles.input}
                onChangeText={setInputUsername}
                value={inputUsername}
                placeholder="Enter Github username"
                autoCapitalize='none'
            />
            <TouchableOpacity style={styles.button} onPress={handleSubmit}>
                <Text style={styles.buttonText}>Submit</Text>
            </TouchableOpacity>
            <ScrollView>
                {userList.map((user, index) => (
                    <TouchableOpacity key={index} onPress={() => handleUserPress(user)}>
                        <View style={styles.bulletContainer}>
                            <Text style={styles.bulletText}>•{user}</Text>
                        </View>
                    </TouchableOpacity>
                ))}
            </ScrollView>
        </KeyboardAvoidingView>
    );
}

const styles = StyleSheet.create({
    container: {
        alignItems: 'center',
        justifyContent: 'center', // centers items vertically in the available space
    },
    input: {
        height: 40,
        borderColor: 'gray',
        borderWidth: 1,
        marginBottom: 16,
        width: '100%', // take full width of the container
    },
    button: {
        backgroundColor: 'lightblue',
        padding: 10,
        marginBottom: 16,
        alignItems: 'center',
        borderRadius: 5,
        width: '30%', // take full width of the container
    },
    buttonText: {
        color: 'white',
    },
    bulletContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        marginVertical: 4,
    },
    bulletText: {
        fontSize: 16,
        lineHeight: 16,
        marginRight: 8,
        marginBottom: 10,
        textDecorationLine: 'underline',
    },
    userText: {
        fontSize: 16,
    },
});

export default User;
