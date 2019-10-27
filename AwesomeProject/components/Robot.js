import React, { Component } from 'react';
import {Image} from 'react-native';

export default class Robot extends Component {
    render() {
        const imageUri = {
            uri: "https://www.thegadgetstore.ie/user/products/large/tobbie-diy-robot-build-your-own-gadget.png"
        };
        return (
            <>
            <Image source={imageUri}
                    style={{width:300, height:300}}
            ></Image>
            </>
        )
    }
}
