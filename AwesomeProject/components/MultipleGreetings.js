import React, { Component } from 'react';
import {Text, View} from 'react-native';

class Greeting extends Component {
    render() {
        let hello = `Hello ${this.props.name} !!! How are you ?!?`
        return (
            <Text> {hello} </Text>
        )
    }
}

export default class MultipleGreetings extends Component {
    render() {
        return (
            <View
                style = {{alignItems: 'center'}} >
                <Greeting name="Boss"></Greeting>
                <Greeting name="Bitch"></Greeting>
                <Greeting name="Boss"></Greeting>
                <Greeting name="Bitch"></Greeting>
                <Greeting name="Boss"></Greeting>
                <Greeting name="Bitch"></Greeting>
                <Greeting name="Boss"></Greeting>
                <Greeting name="Bitch"></Greeting>
            </View>
        )
    }
}
