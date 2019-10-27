/**
 * @format
 */

import React, { Component } from 'react';
import {Text, View, StyleSheet, AppRegistry} from 'react-native';
import App from './App';
import {name as appName} from './app.json';
import Robot from './components/Robot';
import MultipleGreetings from './components/MultipleGreetings';
import TextBlink from './components/TextBlink';

export default class index extends Component {
    render() {
        return (
            <View style={styles.container}>
                <Text style={styles.firstText}>AAAAAAAAA</Text>
            </View>
        );
    }
}

const styles = StyleSheet.create({
    container: {
        marginTop: 50
    },
    firstText: {
        margin: 10,
        color: 'red'
    }
})
AppRegistry.registerComponent(appName, () => index);
