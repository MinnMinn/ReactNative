import React from "react";
import { 
  View,
  Text,
  StyleSheet,
  ImageBackground,
  Dimensions,
  Image
} from "react-native";
import { connect } from 'react-redux';
import getFood from './sagas/Api';
import { getFoodById } from './actions/index'

export default class Detail extends React.Component{

  constructor(props) {
    super(props);
    this.state = {
      item: this.props.navigation.getParam('item')
    }
  }

  render(){
    const { item } = this.state;
    console.log(item);
    return(
      <View style={styles.container}>
          <View style={styles.header}>
              <ImageBackground
              source={require("./asset/header.png")}
              style={styles.imageBackground}
              resizeMode="contain"
              >
                  <Text style={styles.title}>DETAIL</Text>
              </ImageBackground>
          </View>
          <View style={styles.tabbar}>
            <View>
              <Text style={styles.textPrice}>{item.description}</Text>
            </View>
            <View style={styles.image_container}>
              <Image 
                source={{uri: item.image}}
                style={styles.image}
              />
          </View>
          </View>
      </View>
    )
  }
}

const width = Dimensions.get("screen").width;

var styles = StyleSheet.create({
  container: {
    flex:1,
    backgroundColor:'white'
  },
  header: {
    marginTop:20,
    position:'absolute'
  },
  textPrice: {
    color:'green',
    fontWeight:'bold',
    paddingLeft: 15,
    paddingRight: 15
  },
  tabbar: {
    flex:1,
    marginTop: width*0.3,
    paddingHorizontal:30,
  },
  imageBackground: {
    width: width*0.4,
    height: width*0.4,
    alignItems:'center'
  },
  title: {
    color:'white',
    marginTop:25,
    fontWeight:'bold',
    fontSize:25
  },
  image_container: {
    width: 250,
    height: 200,
    flexDirection: 'row',
    alignItems: 'flex-end'
  },
  image: {
    width: '100%',
    height: '100%',
    borderWidth: 5,
    borderColor: 'white',
    borderRadius: 10
  }
});

// function mapStateToProps(state) {
//   return { 
    
//   };
// }

// function mapDispatchToProps(dispatch, props) {
//   return {
//     getFoodById : (food) => {
//       dispatch(getFoodById(food));
//     }
//   }
// }

// export default connect(mapStateToProps, mapDispatchToProps)(Detail);