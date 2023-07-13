import React, { useEffect, useState } from 'react';
import { FlatList, View, StyleSheet, Text } from 'react-native';

const RepoList = ({ username }) => {
    const [repos, setRepos] = useState([]);

    useEffect(() => {
        if (username !== '') {
            fetch(`http://localhost:8080/repos?username=${username}`)
                .then(response => response.json())
                .then(data => setRepos(data))
                .catch(error => console.error(error));
        }
    }, [username]);

    return (
        // add text "Hi"
        
        <FlatList
            data={repos}
            renderItem={({ item }) => (
                <View style={styles.repoContainer}>
                    <Text style={styles.repoLink}>Name: {item.name}</Text>
                    <Text>{item.html_url}</Text>
                    <Text>Stars: {item.stargazers_count}</Text>
                    <Text>Forks: {item.forks_count}</Text>
                    <Text>Issues: {item.open_issues_count}</Text>
                </View>
            )}
            keyExtractor={(item, index) => index.toString()}
            style={styles.flatList} // apply style to limit the height
        />
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
