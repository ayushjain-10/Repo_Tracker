import React, { useEffect, useState } from 'react';
import { Text, View, StyleSheet } from 'react-native';

const Info = ({ username }) => {
    const [userData, setUserData] = useState({});

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
});

export default Info;
