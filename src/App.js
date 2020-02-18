import React from "react";
import { 
  View,
  Text,
  StyleSheet
} from "react-native";
import RootStack from "./RootStack";
import { createStore, applyMiddleware } from 'redux';
import allReducers from './reducers/index';
import { Provider } from 'react-redux';

const store = createStore(allReducers);

export default class App extends React.Component{
  render(){
    return(
      <Provider store={store}>
        <RootStack />
      </Provider>
    )
  }
}
