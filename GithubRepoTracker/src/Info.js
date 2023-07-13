import React, { useEffect, useState } from 'react';
import { Text, View, StyleSheet, TextInput, Button } from 'react-native';
import { Modal, TouchableOpacity } from 'react-native';


const Info = ({ username }) => {
    const [userData, setUserData] = useState({});
    const [email, setEmail] = useState('');
    const [modalVisible, setModalVisible] = useState(false);


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
        // Empty the email field
        setEmail('');
        setModalVisible(true);
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
                        <Modal
                            animationType="slide"
                            transparent={true}
                            visible={modalVisible}
                            onRequestClose={() => {
                                setModalVisible(!modalVisible);
                            }}
                        >
                            <View style={styles.centeredView}>
                                <View style={styles.modalView}>
                                    <View style={styles.modalHeader}>
                                        <TouchableOpacity
                                            style={{ ...styles.openButton, backgroundColor: '#2196F3' }}
                                            onPress={() => {
                                                setModalVisible(!modalVisible);
                                            }}
                                        >
                                            <Text style={styles.textStyle}>X</Text>
                                        </TouchableOpacity>
                                    </View>
                                    <View style={styles.modalBody}>
                                        <Text style={styles.modalText}>Email Sent!👍</Text>
                                    </View>
                                </View>
                            </View>
                        </Modal>

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
    centeredView: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        marginTop: 12,
    },
    modalView: {
        margin: 20,
        height: 150,
        width: 200,
        backgroundColor: 'lightgreen',
        borderRadius: 20,
        padding: 15,
        flexDirection: 'column',
        justifyContent: 'flex-start',
        alignItems: 'stretch',
        shadowColor: '#000',
        shadowOffset: {
            width: 0,
            height: 2
        },
        shadowOpacity: 0.25,
        shadowRadius: 3.84,
        elevation: 5
    },
    openButton: {
        backgroundColor: '#F194FF',
        borderRadius: 20,
        padding: 10,
        elevation: 2
    },
    modalHeader: {
        alignSelf: 'flex-end',
    },
    textStyle: {
        color: 'white',
        fontWeight: 'bold',
        textAlign: 'right'
    },
    modalText: {
        marginBottom: 15,
        textAlign: 'center',
        fontSize: 18,
    },
    modalBody: {
        marginTop: 10,
    },
});

export default Info;
