import React, { useEffect, useState } from 'react';
import { Text, View, StyleSheet, TextInput, Button } from 'react-native';

const Info = ({ username }) => {
    const [userData, setUserData] = useState({});
    const [email, setEmail] = useState(''); 

    const handleEmailSubmit = () => {
        fetch('http://localhost:8080/email', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                recipient: email,
            }),
        })
        .then(response => {
            if (!response.ok) {
                console.log(`Email to be sent to: ${email}`);
                throw new Error('HTTP error ' + response.status);
            }
        })
        .catch(error => console.error(error));
    }

    useEffect(() => {
        if (username !== '') {
            fetch(`http://localhost:8080/repos?username=${username}`)
                .then(response => response.json())
                .then(data => {
                    let totalStars = data.reduce((prev, cur) => prev + cur.stargazers_count, 0);

                    setUserData({
                        totalRepos: data.length,
                        totalStars: totalStars,
                    });
                })
                .catch(error => console.error(error));
        }
    }, [username]);

    return (
        <View style={styles.container}>
            <Text style={styles.title}>Github User Info</Text>
            {username !== '' ? (
                <>
                    <Text style={styles.info}>Username: {username}</Text>
                    <Text style={styles.info}>Total Repositories: {userData.totalRepos}</Text>
                    <Text style={styles.info}>Total Stars: {userData.totalStars}</Text>
                    <View style={styles.email}>
                        <Text style={styles.Info}>Send Info </Text>
                        <TextInput
                            style={styles.input}
                            onChangeText={setEmail}
                            value={email}
                            autoCapitalize='none'
                            placeholder="  Recipient's Email"
                        />
                        <Button title="Send Email" onPress={handleEmailSubmit} />
                    </View>
                </>
            ) : (
                <Text style={styles.info}>Please enter a username in the RepoList tab</Text>
            )}
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        padding: 20,
    },
    title: {
        fontSize: 24,
        marginBottom: 20,
    },
    info: {
        fontSize: 18,
        marginBottom: 10,
    },
    Info: {
        fontSize: 18,
        marginBottom: 10,
    },
    email: {
        padding: 10,
        flexDirection: 'column',
        alignItems: 'center',
        marginVertical: 4,
    },
    input: {
        height: 40,
        borderColor: 'gray',
        borderWidth: 1,
        borderRadius: 8,
        marginBottom: 16,
        width: '80%', // take full width of the container
    },
});

export default Info;
