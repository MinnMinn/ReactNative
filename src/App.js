import React, { Component } from 'react'
import { View, TextInput} from 'react-native';
import RootStack from './RootStack';
import { createStore, applyMiddleware } from 'redux';
import allReducers from './reducers/index';
import { Provider } from 'react-redux';

const store = createStore(allReducers);

export default class App extends Component {
    render() {
        return (
            <Provider store={store}>
                <RootStack />
            </Provider>
        )
    }
}
