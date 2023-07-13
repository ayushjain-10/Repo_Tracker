import React, { useState } from 'react';
import { ScrollView, Text, TextInput, TouchableOpacity, FlatList, View, StyleSheet, KeyboardAvoidingView } from 'react-native';

const RepoList = () => {
    const [repos, setRepos] = useState([]);
    const [username, setUsername] = useState('');

    const handlePress = () => {
        fetch(`http://localhost:8080/repos?username=${username}`)
            .then(response => response.json())
            .then(data => setRepos(data))
            .catch(error => console.error(error));
    };

    return (
        <KeyboardAvoidingView style={styles.container} behavior={Platform.OS === 'ios' ? 'padding' : null}>
            <TextInput
                style={styles.input}
                onChangeText={setUsername}
                value={username}
                placeholder="Enter Github username"
                autoCapitalize='none'
            />
            <TouchableOpacity style={styles.button} onPress={handlePress}>
                <Text style={styles.buttonText}>Get Repos</Text>
            </TouchableOpacity>
            <FlatList
                data={repos}
                renderItem={({ item }) => (
                    <View style={styles.repoContainer}>
                        <Text style={styles.repoLink}>{item.link}</Text>
                        <Text>Stars: {item.stars}</Text>
                        <Text>Forks: {item.forks}</Text>
                        <Text>Issues: {item.issues}</Text>
                    </View>
                )}
                keyExtractor={(item, index) => index.toString()}
                style={styles.flatList} // apply style to limit the height
            />
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
        text: 'hi',
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
    flatList: {
        height: '100%', // limit the height to 70% of the remaining screen space
    },
    repoContainer: {
        marginBottom: 16,
      },
      repoLink: {
        fontWeight: 'bold',
        marginBottom: 8,
      }
});

export default RepoList;
